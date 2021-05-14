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
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/hokaccha/go-prettyjson"
	"github.com/spf13/cobra"
)

var reposAddName string
var reposAddPath string
var reposAddDesc string

func PrintObject(obj interface{}) {
	bts, err := prettyjson.Marshal(obj)
	CheckIfError(err)
	Info(string(bts))
}

func RealPathToStoragePath(dirpath string) string {
	if !path.IsAbs(dirpath) {
		wd, _ := os.Getwd()
		dirpath = path.Join(wd, dirpath)
	}
	homedir, err := os.UserHomeDir()
	CheckIfError(err)
	return strings.ReplaceAll(dirpath, homedir, "~")
}

func StoragePathToRealPath(storagePath string) string {
	homedir, err := os.UserHomeDir()
	CheckIfError(err)
	return strings.ReplaceAll(storagePath, "~", homedir)
}

func NewRepo(name, dirpath, desc string) (*Repo, error) {
	r, err := git.PlainOpen(dirpath)
	if err != nil {
		return nil, err
	}
	remotes, err := r.Remotes()
	CheckIfError(err)
	if len(remotes) == 0 {
		return nil, fmt.Errorf("远程仓库不能为空")
	}
	reposAddPath = RealPathToStoragePath(dirpath)
	reposAddName = path.Base(dirpath)

	repo := &Repo{
		Name:        reposAddName,
		Path:        reposAddPath,
		Description: desc,
	}
	whitelist := []string{"github.com", "coding.net"}
	for _, remote := range remotes {
		flag := false
		for _, white := range whitelist {
			containsAuthor := strings.Contains(remote.Config().URLs[0], "jeremaihloo") || strings.Contains(remote.Config().URLs[0], "jerloo")
			if strings.Contains(remote.Config().URLs[0], white) && containsAuthor {
				flag = true
			}
		}

		if flag {
			head, _ := r.Head()
			rr := &RepoRemote{
				Name:          remote.Config().Name,
				Address:       remote.Config().URLs[0],
				CurrentBranch: head.Name().Short(),
			}
			repo.Remotes = append(repo.Remotes, rr)
		} else {
			return nil, fmt.Errorf("不得添加除白名单意外的仓库地址")
		}
	}
	return repo, nil
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "添加仓库到配置文件",
	Run: func(cmd *cobra.Command, args []string) {
		storage := GetRepoStorage()

		if len(args) == 0 {
			Warning("需要指定添加目录")
			return
		}

		workdir := args[0]

		_, err := git.PlainOpen(workdir)
		if err == nil {
			repo, err := NewRepo(path.Base(workdir), workdir, reposAddDesc)
			CheckIfError(err)
			storage.Add(repo)
			return
		}

		dirs, err := os.ReadDir(workdir)
		CheckIfError(err)
		for _, item := range dirs {
			if item.IsDir() {
				dirPath := path.Join(workdir, item.Name())
				repo, err := NewRepo(item.Name(), dirPath, "")
				if err == nil {
					err = storage.Add(repo)
					if err != nil {
						Warning("忽略 %s %s", dirPath, err.Error())
					} else {
						Info("Added %s", dirPath)
					}
				} else {
					Warning("忽略 %s", dirPath)
				}
			}
		}
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
	addCmd.Flags().StringVar(&reposAddDesc, "description", "", "repo description")
}
