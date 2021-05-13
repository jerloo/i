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
	"path/filepath"
	"strings"

	"os"

	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

func IfInGitRepoDir() (string, bool) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	info, err := os.Stat(filepath.Join(wd, ".git"))
	if err != nil {
		return wd, false
	}
	return wd, info.IsDir()
}

// CheckArgs should be used to ensure the right command line arguments are
// passed before executing an example.
func CheckArgs(arg ...string) {
	if len(os.Args) < len(arg)+1 {
		Warning("Usage: %s %s", os.Args[0], strings.Join(arg, " "))
		os.Exit(1)
	}
}

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

// Info should be used to describe the example commands that are about to run.
func Info(format string, args ...interface{}) {
	green := color.New(color.FgGreen).SprintFunc()
	fmt.Println(green(fmt.Sprintf(format, args...)))
}

// Warning should be used to display a warning
func Warning(format string, args ...interface{}) {
	red := color.New(color.FgRed).SprintFunc()
	fmt.Println(red(fmt.Sprintf(format, args...)))
}

func IfRepoIsClean(wd string) bool {
	r, err := git.PlainOpen(wd)
	CheckIfError(err)

	w, err := r.Worktree()
	CheckIfError(err)

	status, err := w.Status()
	CheckIfError(err)

	return status.IsClean()
}

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "查看当前目录仓库状态",
	Long:  "查看当前目录仓库状态, 如果当前目录为仓库目录, 就查询当前目录状态, 否则就查询当前目录下所有仓库状态",
	Run: func(cmd *cobra.Command, args []string) {
		if wd, ok := IfInGitRepoDir(); ok {
			if IfRepoIsClean(wd) {
				fmt.Println("仓库很干净")
			} else {
				fmt.Println("仓库不干净")
			}
		}
	},
}

func init() {
	gitCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
