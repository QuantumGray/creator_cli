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
	"creator/util/contexts"
	"creator/util/createapp"
	"creator/util/flutter"
	"creator/util/websupport"

	"github.com/spf13/cobra"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create an app from a template with a SHA",
	Long:  `do it 4real`,
	Args: func(cmd *cobra.Command, args []string) error {
		/*
			if err := validators.CreateCommandArgsValidation(args); err != nil {
				return err
			}
			return nil
		*/
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		ctx := contexts.NewContext("createContext")
		ctx.GetValue["SHA"] = args[0]
		createapp.CreateApp(ctx)
		if contains(args, "web") {
			websupport.EnableWeb(ctx)
			ctx.GetValue["WEB"] = "enable"
		}
		if contains(args, "run") {
			flutter.Run(ctx)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
