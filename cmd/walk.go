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
	"log"
	"os"

	"github.com/spf13/cobra"
)

// walkCmd represents the walk command
var walkCmd = &cobra.Command{
	Use:   "walk",
	Short: "walk allows you to recursively find all the files and directories in a given location.",
	Long: `walk command allows you to recursively find all the files and directories in a given location 
and export that information into either YAML or JSON format. It can be customized using either --json or --yaml. 
If you choose to include only a certain files in that walk you can provide the --contains flag. If you wish to remove extension 
from the output use --no-extn flag. 

Finally the output can be customized to show certain values. For example as the value for the filename key you can choose --show-size --last-modified,
by default the value is "true". if --dry-run flag is used to show the output on the console before checking out the file. --in and --out flags are used respectively
to consume input and output path.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("walk called")
	},
}

var cwd=""

func init() {
	rootCmd.AddCommand(walkCmd)
	cwd,err := os.Getwd()
	if err!=nil{
		log.Println("Unable to read current working dir error: ",err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	walkCmd.Flags().String("in", cwd, "Path of the directory that you want to walk. Default is working dir")
	walkCmd.Flags().String("out", cwd, "output where you want to store JSON/YAML file. Default is working dir")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// walkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
