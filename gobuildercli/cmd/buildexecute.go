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
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// buildexecuteCmd represents the buildexecute command
var buildexecuteCmd = &cobra.Command{
	Use:   "buildexecute",
	Short: "A cli tool",
	Long:  `buildexecute cli tool`,
	Run: func(cmd *cobra.Command, args []string) {
		copyDir, _ := cmd.Flags().GetString("copydir")
		// fmt.Println(cd)
		buildDir, _ := cmd.Flags().GetString("builddir")
		// fmt.Println(bd)
		executable, _ := cmd.Flags().GetString("exe")
		// fmt.Println(ex)
		excludeTest, _ := cmd.Flags().GetBool("exclude-tests")
		// fmt.Println(testIgnore)
		processAllCommands(copyDir, buildDir, executable, excludeTest)
	},
}

func processAllCommands(copyDir, buildDir, executable string, exludeTests bool) {
	if copyDir != "" {
		path, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		if path == copyDir && buildDir == "" {
			fmt.Println("same directory")
		} else {
			if buildDir != "" {
				if exludeTests {
					_, err := exec.Command("rsync", "-avz", "--exclude", "*test.go", copyDir, buildDir).Output()
					if err != nil {
						panic(err)
					}
				} else {
					_, err = exec.Command("cp", "-R", copyDir, buildDir).Output()
					if err != nil {
						panic(err)
					}
				}
				if executable != "" {
					os.Chdir(buildDir)
					newDir, err := os.Getwd()
					if err != nil {
					}
					fmt.Printf("Current Working Direcotry: %s\n", newDir)
					cmd := exec.Command("go", "build")
					err = cmd.Run()
					if err != nil {
						panic(err)
					}
					out, err := exec.Command("./" + executable).Output()
					if err != nil {
						panic(err)
					}
					fmt.Println(string(out[:]))
				}
			} else {
				_, err = exec.Command("cp", "-R", copyDir, ".").Output()
				if err != nil {
					panic(err)
				}
			}

		}
	}
}

func init() {
	rootCmd.AddCommand(buildexecuteCmd)
	// buildexecuteCmd.Flags().BoolP("copydir", "", false, "copy directory")
	// buildexecuteCmd.Flags().BoolP("builddir", "", false, "build directory")
	buildexecuteCmd.Flags().StringP("copydir", "c", "", "copydirectory")
	buildexecuteCmd.Flags().StringP("builddir", "b", "", "builddirectory")
	buildexecuteCmd.Flags().StringP("exe", "e", "", "execute")
	buildexecuteCmd.Flags().BoolP("exclude-tests", "", false, "exclude-tests")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildexecuteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildexecuteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
