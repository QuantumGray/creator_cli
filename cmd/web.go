package cmd

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

import (
	"creator/util/websupport"

	"github.com/spf13/cobra"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "toggle the flutter web config",
	Long:  `do it 4real`,
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "enable" {
			websupport.ToggleWebIntegration(true)
		} else if args[0] == "disable" {
			websupport.ToggleWebIntegration(false)
		} else {
			websupport.ToggleWebIntegration(false)
		}
	},
}

func init() {
	rootCmd.AddCommand(webCmd)
}
