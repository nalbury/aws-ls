// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/spf13/cobra"
)

func nameTag(tags []*ec2.Tag) string {
	var name string
	for _, tag := range tags {
		if *tag.Key == "Name" {
			name = *tag.Value
		}
	}

	return name
}

func cleanStringP(s *string) string {
	var clean string
	if s != nil {
		clean = *s
	} else {
		clean = "None"
	}
	return clean
}

func listInstances(profile string, noHeaders bool) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Profile:           profile,
	}))
	ec2Svc := ec2.New(sess)

	var instances []*ec2.Instance
	//result, err := ec2Svc.DescribeInstances(nil)
	err := ec2Svc.DescribeInstancesPages(nil,
		func(page *ec2.DescribeInstancesOutput, lastPage bool) bool {
			for _, reservation := range page.Reservations {
				for _, instance := range reservation.Instances {
					instances = append(instances, instance)
				}
			}
			return lastPage
		})

	if err != nil {
		fmt.Println("Error", err)
	} else {
		const padding = 4
		w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)
		if !noHeaders {
			titleAttributes := []string{"INDEX", "NAME", "INSTANCE_ID", "PRIVATE_IP", "INSTANCE_TYPE", "STATUS"}
			titleRow := strings.Join(titleAttributes, "\t")
			fmt.Fprintln(w, titleRow)
		}

		sort.Slice(instances, func(i, j int) bool {
			return nameTag(instances[i].Tags) < nameTag(instances[j].Tags)
		})

		for i, instance := range instances {
			stringdex := strconv.Itoa(i)
			inst := *instance
			instanceName := nameTag(inst.Tags)
			instanceId := cleanStringP(inst.InstanceId)
			privateIp := cleanStringP(inst.PrivateIpAddress)
			instanceType := cleanStringP(inst.InstanceType)
			status := cleanStringP(inst.State.Name)
			instanceAttributes := []string{stringdex, instanceName, instanceId, privateIp, instanceType, status}
			instanceRow := strings.Join(instanceAttributes, "\t")
			fmt.Fprintln(w, instanceRow)
		}
		w.Flush()
	}
}

// ec2Cmd represents the ec2 command
var ec2Cmd = &cobra.Command{
	Use:   "ec2",
	Short: "ls your ec2 instances",
	Long: `ls your ec2 instances. 
Provides instance id, name, status and private IP.`,
	Run: func(cmd *cobra.Command, args []string) {
		listInstances(profile, noHeaders)
	},
}

func init() {
	rootCmd.AddCommand(ec2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ec2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ec2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
