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
	"os/exec"
	"path"
	"path/filepath"

	"github.com/spf13/cobra"
)

var cpNight bool

// cpCmd represents the cp command
var cpCmd = &cobra.Command{
	Use:   "cp",
	Short: "mipmap 拷贝工作",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("请输入正确的参数")
		} else {
			dirNames := []string{
				"mipmap-mdpi",
				"mipmap-hdpi",
				"mipmap-xhdpi",
				"mipmap-xxhdpi",
				"mipmap-xxxhdpi",
			}
			if cpNight {
				dirNames = []string{
					"mipmap-mdpi-night",
					"mipmap-hdpi-night",
					"mipmap-xhdpi-night",
					"mipmap-xxhdpi-night",
					"mipmap-xxxhdpi-night",
				}
			}
			// currentDir, _ := os.Getwd()
			for _, item := range dirNames {
				srcAbs, _ := filepath.Abs(filepath.Join(args[0], item))
				files, _ := ioutil.ReadDir(srcAbs)
				dstAbs, _ := filepath.Abs(filepath.Join(args[1], item))
				for _, file := range files {
					srcfp := path.Join(srcAbs, file.Name())
					dstfp := path.Join(dstAbs, file.Name())
					cmd := exec.Command("cp", srcfp, dstfp)
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					err := cmd.Run()
					if err != nil {
						fmt.Printf("操作失败 %v", err)
					}
				}
			}
		}
	},
}

func init() {
	mipmapCmd.AddCommand(cpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	cpCmd.Flags().BoolVar(&cpNight, "night", false, "是否是夜间模式资源 -night 结尾")
}
