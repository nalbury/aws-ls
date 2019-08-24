/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/spf13/cobra"
)

func listEcrRepos(profile string, noHeaders bool) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Profile:           profile,
	}))
	ecrSvc := ecr.New(sess)
	result, err := ecrSvc.DescribeRepositories(nil)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		const padding = 4
		w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)
		if !noHeaders {
			titleAttributes := []string{"INDEX", "NAME", "URI"}
			titleRow := strings.Join(titleAttributes, "\t")
			fmt.Fprintln(w, titleRow)
		}
		repos := result.Repositories
		sort.Slice(repos, func(i, j int) bool {
			return *repos[i].RepositoryName < *repos[j].RepositoryName
		})
		for i, repo := range repos {
			stringdex := strconv.Itoa(i)
			r := *repo
			repoName := cleanStringP(r.RepositoryName)
			repoUri := cleanStringP(r.RepositoryUri)

			repoAttributes := []string{stringdex, repoName, repoUri}
			repoRow := strings.Join(repoAttributes, "\t")
			fmt.Fprintln(w, repoRow)

		}
		w.Flush()
	}
}

// ecrCmd represents the ecr command
var ecrCmd = &cobra.Command{
	Use:   "ecr",
	Short: "ls your ecr repos",
	Long: `ls your ecr repos.
Provides the name, and ...`,
	Run: func(cmd *cobra.Command, args []string) {
		listEcrRepos(profile, noHeaders)
	},
}

func init() {
	rootCmd.AddCommand(ecrCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ecrCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ecrCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
