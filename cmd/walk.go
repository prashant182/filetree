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
	"encoding/json"
	"fmt"
	"github.com/prashant182/filetree/pkg/node"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/spf13/cobra"
)

// walkCmd represents the walk command
var walkCmd = &cobra.Command{
	Use:   "walk",
	Short: "walk allows you to recursively find all the files and directories in a given location.",
	Long: `

walk command allows you to recursively find all the files and directories in a given location 
and export that information into either YAML or JSON format. It can be customized using either --json or --yaml. 
If you choose to include only a certain files in that walk you can provide the --contains flag. 
If you wish to remove extension from the output use --no-extn flag.By default the value is "true". 
if --dry-run flag is used to show the output on the console before checking out the file. 
--in and --out flags are used respectively to consume input and output path.`,
	Run: func(cmd *cobra.Command, args []string) {
		root := node.DFS(inp)
		m := root.ToMap(contains, camelCase, noExtn)
		if outType == "json" {
			jsn = true
		}

		if jsn {
			output, err := json.Marshal(m)
			if err != nil {
				log.Println("JSON marshal error ", err)
			}
			if dryrun {
				fmt.Println(string(output))
			} else {
				err := ioutil.WriteFile(path.Join(op, "out.json"), output, 0644)
				if err != nil {
					log.Println("Error writing JSON file ", err)
				}
			}
		} else {
			output, err := yaml.Marshal(m)
			if err != nil {
				log.Println("Y marshal error ", err)
			}
			if dryrun {
				fmt.Println(string(output))
			} else {
				err := ioutil.WriteFile(path.Join(op, "out.yaml"), output, 0644)
				if err != nil {
					log.Println("Error writing YAML file ", err)
				}
			}
		}

	},
}

var inp = ""
var op = ""
var contains = ""
var noExtn bool
var camelCase bool
var outType = ""
var dryrun bool
var jsn = false

func init() {
	rootCmd.AddCommand(walkCmd)
	cwd, err := os.Getwd()
	if err != nil {
		log.Println("Unable to read current working dir error: ", err)
	}
	// and all subcommands,
	walkCmd.Flags().StringVar(&inp, "in", cwd, "Path of the directory that you want to walk.")
	walkCmd.Flags().StringVar(&op, "out", cwd, "output where you want to store JSON/YAML file.")
	walkCmd.Flags().StringVar(&outType, "outType", "yaml", "Either json or yaml")
	walkCmd.Flags().StringVar(&contains, "contains", "", "filters the output against the match")
	walkCmd.Flags().BoolVar(&camelCase, "camel-case", false, "converts the file names against to camelcase from snake case")
	walkCmd.Flags().BoolVar(&dryrun, "dry-run", false, "prints output on console before not on file")
	walkCmd.Flags().BoolVar(&noExtn, "no-extn", false, "removes the extension from the filename at the time of output")
}
