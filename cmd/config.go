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
	"time"

	"github.com/spf13/cobra"
)

var reposConfigDesc string
var reposConfigAddress string
var reposConfigPath string

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "配置仓库",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			Warning("请指定仓库名称")
			return
		}
		storage := GetRepoStorage()
		var repo *Repo

		for _, item := range storage.Repos {
			if item.Name == args[0] {
				repo = item
			}
		}
		if repo == nil {
			Warning("仓库不存在")
			return
		}
		flag := false
		if reposConfigPath != "" {
			repo.Path = reposConfigPath
			flag = true
		}
		if reposConfigDesc != "" {
			repo.Description = reposConfigDesc
			flag = true
		}
		if reposConfigAddress != "" {
			repo.Address = reposConfigAddress
			flag = true
		}
		if flag {
			repo.UpdatedAt = time.Now()
			err := storage.Save()
			CheckIfError(err)
		}
	},
}

func init() {
	reposCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	configCmd.Flags().StringVar(&reposConfigPath, "path", "", "repo new path")
	configCmd.Flags().StringVar(&reposConfigDesc, "description", "", "repo new description")
	configCmd.Flags().StringVar(&reposConfigAddress, "address", "", "repo new address")
}
