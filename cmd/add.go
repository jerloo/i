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
	"os"
	"path"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/hokaccha/go-prettyjson"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
)

var reposAddName string
var reposAddPath string
var reposAddAddress string
var reposAddDesc string

func PrintObject(obj interface{}) {
	bts, err := prettyjson.Marshal(obj)
	CheckIfError(err)
	Info(string(bts))
}

func RealPathToStoragePath(dirpath string) string {
	homedir, err := os.UserHomeDir()
	CheckIfError(err)
	return strings.ReplaceAll(dirpath, homedir, "~")
}

func StoragePathToRealPath(storagePath string) string {
	homedir, err := os.UserHomeDir()
	CheckIfError(err)
	return strings.ReplaceAll(storagePath, "~", homedir)
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "添加仓库到配置文件",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 && args[0] == "." {
			wd, err := os.Getwd()
			CheckIfError(err)
			reposAddPath = RealPathToStoragePath(wd)
			reposAddName = path.Base(wd)
			r, err := git.PlainOpen(wd)
			CheckIfError(err)
			remotes, err := r.Remotes()
			CheckIfError(err)
			if len(remotes) == 0 {
				Warning("远程仓库不能为空")
				return
			}
			reposAddAddress = remotes[0].Config().URLs[0]
		}

		if reposAddName == "" || reposAddPath == "" || reposAddAddress == "" {
			Warning("参数不能为空")
			return
		}
		storage := GetRepoStorage()

		for _, item := range storage.Repos {
			if item.Name == reposAddName || item.Path == reposAddPath || item.Address == reposAddAddress {
				Warning("已经存在相同仓库")
				PrintObject(item)
				return
			}
		}
		newRepo := &Repo{
			ID:        uuid.NewV4().String(),
			Name:      reposAddName,
			Path:      reposAddPath,
			Address:   reposAddAddress,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		storage.Repos = append(storage.Repos, newRepo)
		err := storage.Save()
		CheckIfError(err)
	},
}

func init() {
	reposCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	addCmd.Flags().StringVar(&reposAddName, "name", "", "repo name")
	addCmd.Flags().StringVar(&reposAddPath, "path", "", "repo path")
	addCmd.Flags().StringVar(&reposAddAddress, "address", "", "repo address")
	addCmd.Flags().StringVar(&reposAddDesc, "description", "", "repo description")
}
