/*
Copyright 2023 The Cloud-Barista Authors.

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
	"github.com/spf13/cobra"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import dummy data into the service",
	Long:  `Import data into the service with the credentials you have entered.`,
}

func init() {
	rootCmd.AddCommand(importCmd)

	importCmd.PersistentFlags().StringVarP(&datamoldParams.CredentialPath, "credential-path", "C", "", "Json file path containing the user's credentials")
	importCmd.PersistentFlags().StringVarP(&datamoldParams.DstPath, "dst-path", "d", "", "Destination path where dummy data exists")
	importCmd.PersistentFlags().BoolVarP(&datamoldParams.TaskTarget, "task", "T", false, "Select a destination(src, dst) to work with in the credential-path")
	importCmd.MarkFlagsRequiredTogether("credential-path", "dst-path")
}
