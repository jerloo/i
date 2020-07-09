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
	"os"

	"github.com/sgoby/opencc"
	"github.com/spf13/cobra"
)

// translateCmd represents the translate command
var translateCmd = &cobra.Command{
	Use:   "translate",
	Short: "文本翻译",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("必须制定需要翻译的文件路径和输出文件路径")
		} else {
			cc, err := opencc.NewOpenCC("s2hk")
			if err != nil {
				fmt.Print(err)
			}
			inFile, err := os.Open(args[0])
			if err != nil {
				fmt.Print(err)
			}
			outFile, err := os.OpenFile(args[1], os.O_CREATE|os.O_APPEND, 0644)
			if err != nil {
				fmt.Print(err)
			}
			err = cc.ConvertFile(inFile, outFile)
			if err != nil {
				fmt.Print(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(translateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// translateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// translateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
