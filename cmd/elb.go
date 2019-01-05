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
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/spf13/cobra"
)

func createListenerDescription(strs ...string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
	}
	return sb.String()
}

func listLoadBalancers(profile string, noHeaders bool) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Profile:           profile,
	}))

	elbSvc := elb.New(sess)
	result, err := elbSvc.DescribeLoadBalancers(nil)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		const padding = 4
		w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)
		if !noHeaders {
			titleAttributes := []string{"INDEX", "DNS_NAME", "INSTANCE_COUNT", "HEALTH_CHECK", "LISTENERS"}
			titleRow := strings.Join(titleAttributes, "\t")
			fmt.Fprintln(w, titleRow)
		}
		loadBalancers := result.LoadBalancerDescriptions

		sort.Slice(loadBalancers, func(i, j int) bool {
			return *loadBalancers[i].DNSName < *loadBalancers[j].DNSName
		})
		for i, loadBalancer := range loadBalancers {
			stringdex := strconv.Itoa(i)
			lb := *loadBalancer
			lbName := cleanStringP(lb.DNSName)
			lbInstanceCount := strconv.Itoa(len(lb.Instances))
			lbHealthCheck := cleanStringP(lb.HealthCheck.Target)

			var listeners []string

			for _, listenerDescription := range loadBalancer.ListenerDescriptions {
				listener := listenerDescription.Listener
				lbProtocol := cleanStringP(listener.Protocol)
				lbPort := strconv.FormatInt(*listener.LoadBalancerPort, 10)
				instanceProtocol := cleanStringP(listener.InstanceProtocol)
				instancePort := strconv.FormatInt(*listener.InstancePort, 10)
				listenerDescription := createListenerDescription(lbProtocol, ":", lbPort, " -> ", instanceProtocol, ":", instancePort)

				listeners = append(listeners, listenerDescription)
			}
			lbListeners := strings.Join(listeners, ", ")

			lbAttributes := []string{stringdex, lbName, lbInstanceCount, lbHealthCheck, lbListeners}
			lbRow := strings.Join(lbAttributes, "\t")
			fmt.Fprintln(w, lbRow)

		}
		w.Flush()
	}
}

var elbCmd = &cobra.Command{
	Use:   "elb",
	Short: "ls your ec2 load balancers",
	Long:  `ls your ec2 instances. Provides the DNS name, instance count, health check, and listener configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		listLoadBalancers(profile, noHeaders)
	},
}

func init() {
	rootCmd.AddCommand(elbCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// elbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// elbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
