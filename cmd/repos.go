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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
)

type Repo struct {
	ID          string
	Name        string
	Description string
	Remotes     []*RepoRemote
	Path        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type RepoRemote struct {
	Name          string
	Address       string
	CurrentBranch string
}

type RepoStorage struct {
	ID          string
	Version     int
	Description string
	Repos       []*Repo
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (rs *RepoStorage) Add(repo *Repo) error {
	for _, item := range rs.Repos {
		if item.Name == repo.Name || item.Path == repo.Path {
			return fmt.Errorf("已经存在相同仓库")
		}
	}
	repo.ID = uuid.NewV4().String()
	repo.CreatedAt = time.Now()
	repo.UpdatedAt = time.Now()
	rs.Repos = append(rs.Repos, repo)
	return rs.Save()
}

func (rs *RepoStorage) Save() error {
	rs.UpdatedAt = time.Now()

	bts, err := json.MarshalIndent(rs, "", "    ")
	if err != nil {
		return err
	}
	homedir, err := os.UserHomeDir()
	storagePath := path.Join(homedir, ".repos.json")
	CheckIfError(err)
	return ioutil.WriteFile(storagePath, bts, 0644)
}

func GetRepoStorage() *RepoStorage {
	homedir, err := os.UserHomeDir()
	CheckIfError(err)

	storage := &RepoStorage{
		ID:          uuid.NewV4().String(),
		Version:     1,
		Description: "我的仓库",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	storagePath := path.Join(homedir, ".repos.json")
	_, err = os.Stat(storagePath)
	if err == nil {
		bts, err := ioutil.ReadFile(storagePath)
		CheckIfError(err)
		err = json.Unmarshal(bts, &storage)
		CheckIfError(err)
	}

	return storage
}

// reposCmd represents the repos command
var reposCmd = &cobra.Command{
	Use:   "repos",
	Short: "仓库管理",
	Run: func(cmd *cobra.Command, args []string) {
		storage := GetRepoStorage()
		for _, item := range storage.Repos {
			Info(fmt.Sprintf("%s %s", strings.Split(item.ID, "-")[0], item.Name))
		}
		if len(storage.Repos) == 0 {
			Info("当前没有仓库")
		}
	},
}

func init() {
	rootCmd.AddCommand(reposCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reposCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reposCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
