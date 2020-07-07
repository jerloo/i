/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var renameDrawable bool

// androidMipmapRenameCmd represents the androidMipmapRename command
var androidMipmapRenameCmd = &cobra.Command{
	Use:   "rename",
	Short: "批量重命名文件",
	Run: func(cmd *cobra.Command, args []string) {
		root, _ := os.Getwd()
		dirs, _ := ioutil.ReadDir(root)
		for _, item := range dirs {
			if strings.Contains(item.Name(), "drawable") {
				old := filepath.Join(root, item.Name())
				newPath := filepath.Join(root, strings.ReplaceAll(item.Name(), "drawable", "mipmap"))
				err := os.Rename(old, newPath)
				if err != nil {
					fmt.Printf("%s \n%s \n%v", old, newPath, err)
				}
			}
		}
	},
}

func init() {
	mipmapCmd.AddCommand(androidMipmapRenameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// androidMipmapRenameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// androidMipmapRenameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	androidMipmapRenameCmd.Flags().BoolVar(&renameDrawable, "drawable", false, "将 drawable 文件夹重命名为 mipmap 文件夹")
}
