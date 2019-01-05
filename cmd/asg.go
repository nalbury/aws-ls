// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/spf13/cobra"
)

func listAutoScalingGroups(profile string, noHeaders bool) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Profile:           profile,
	}))

	asgSvc := autoscaling.New(sess)
	result, err := asgSvc.DescribeAutoScalingGroups(nil)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		const padding = 4
		w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)
		if !noHeaders {
			titleAttributes := []string{"INDEX", "NAME", "DESIRED", "CURRENT", "MIN", "MAX", "LAUNCH_CONFIG"}
			titleRow := strings.Join(titleAttributes, "\t")
			fmt.Fprintln(w, titleRow)
		}
		groups := result.AutoScalingGroups
		sort.Slice(groups, func(i, j int) bool {
			return *groups[i].AutoScalingGroupName < *groups[j].AutoScalingGroupName
		})
		for i, group := range groups {
			stringdex := strconv.Itoa(i)
			asg := *group
			asgName := cleanStringP(asg.AutoScalingGroupName)
			desiredInstances := strconv.FormatInt(*asg.DesiredCapacity, 10)
			currentInstances := strconv.Itoa(len(asg.Instances))
			minInstances := strconv.FormatInt(*asg.MinSize, 10)
			maxInstances := strconv.FormatInt(*asg.MaxSize, 10)
			launchConfigName := cleanStringP(asg.LaunchConfigurationName)

			asgAttributes := []string{stringdex, asgName, desiredInstances, currentInstances, minInstances, maxInstances, launchConfigName}
			asgRow := strings.Join(asgAttributes, "\t")
			fmt.Fprintln(w, asgRow)

		}
		w.Flush()
	}
}

// asgCmd represents the asg command
var asgCmd = &cobra.Command{
	Use:   "asg",
	Short: "ls your auto scaling groups",
	Long: `ls your auto scaling groups. 
Provides the name, desired/current/min/max instance counts, and the launch config name`,
	Run: func(cmd *cobra.Command, args []string) {
		listAutoScalingGroups(profile, noHeaders)
	},
}

func init() {
	rootCmd.AddCommand(asgCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// asgCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// asgCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
