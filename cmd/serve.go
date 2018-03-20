// Copyright © 2018 Anshul Sanghi <anshap1719@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var ngserve string
var goserve string


// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		quit := make(chan int)

		argss := []string{"serve"}

		if ngserve != "" {
			ngArgs := strings.Split(ngserve, " ")
			for _, arg := range ngArgs {
				argss = append(argss, arg)
			}
		}

		go runExternalCmd("ng", argss)

		argss = []string{}

		if goserve != "" {
			goArgs := strings.Split(ngserve, " ")
			for _, arg := range goArgs {
				argss = append(argss, arg)
			}
		}
		os.Chdir("./src/server/")
		go runExternalCmd("gin", []string{})

		<-quit
	},
}

func init() {
	serveCmd.Flags().StringVarP(&ngserve, "ng", "", "", "Set Arguments For ng serve")
	RootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
