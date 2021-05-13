/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"errors"
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "推送本地所有代码",
	Long:  "如果所在命令是一个仓库, 就推送当前仓库, 否则就推送目录下所有仓库",
	Run: func(cmd *cobra.Command, args []string) {
		if wd, ok := IfInGitRepoDir(); ok {
			r, err := git.PlainOpen(wd)
			CheckIfError(err)

			if IfRepoIsClean(wd) {
				err = r.Push(&git.PushOptions{})
				if errors.Is(err, git.NoErrAlreadyUpToDate) {
					fmt.Println("已是最新")
				}
			} else {
				Warning("当前仓库不干净")
			}
		}
	},
}

func init() {
	gitCmd.AddCommand(pushCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pushCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pushCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
