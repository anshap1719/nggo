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
	"fmt"

	"github.com/spf13/cobra"
	"os"
	//"github.com/mgutz/ansi"
	"io/ioutil"
)

var jsonFileContent = "package utils\r\n\r\nimport (\r\n  \"encoding/json\"\r\n  \"bytes\"\r\n)\r\n\r\nfunc StructToJson (data interface{}) ([]byte, error) {\r\n  buf := new(bytes.Buffer)\r\n\r\n  if err := json.NewEncoder(buf).Encode(data); err != nil {\r\n    return nil, err\r\n  }\r\n\r\n  return buf.Bytes(), nil\r\n}"

var appComponentContent = "import {Component, OnInit} from '@angular/core';\r\nimport {HelloWorldService} from './hello-world.service';\r\n\r\n@Component({\r\n  selector: 'app-root',\r\n  templateUrl: './app.component.html',\r\n  styleUrls: ['./app.component.scss']\r\n})\r\nexport class AppComponent implements OnInit {\r\n\r\n  title;\r\n\r\n  constructor(private hw: HelloWorldService) {}\r\n\r\n  ngOnInit() {\r\n    this.hw.getTitle()\r\n      .subscribe(data => this.title = data.title);\r\n  }\r\n\r\n}\r\n"

var appModuleContent = "import { BrowserModule } from '@angular/platform-browser';\r\nimport { NgModule } from '@angular/core';\r\n\r\nimport { AppComponent } from './app.component';\r\nimport {HelloWorldService} from './hello-world.service';\r\nimport {HttpModule} from '@angular/http';\r\n\r\n@NgModule({\r\n  declarations: [\r\n    AppComponent\r\n  ],\r\n  imports: [\r\n    BrowserModule,\r\n    HttpModule\r\n  ],\r\n  providers: [HelloWorldService],\r\n  bootstrap: [AppComponent]\r\n})\r\nexport class AppModule { }"

var helloWorldServiceContent = "import { Injectable } from '@angular/core';\r\nimport {Http} from '@angular/http';\r\nimport {environment} from '../environments/environment';\r\nimport 'rxjs/add/operator/map';\r\n\r\n@Injectable()\r\nexport class HelloWorldService {\r\n\r\n  constructor(private http: Http) { }\r\n\r\n  getTitle() {\r\n    return this.http.get(`${environment.serverUrl}/hello-world`)\r\n      .map(response => response.json());\r\n  }\r\n\r\n}"

var environmentContent = "// The file contents for the current environment will overwrite these during build.\r\n// The build system defaults to the dev environment which uses `environment.ts`, but if you do\r\n// `ng build --env=prod` then `environment.prod.ts` will be used instead.\r\n// The list of which env maps to which file can be found in `.angular-cli.json`.\r\n\r\nexport const environment = {\r\n  production: false,\r\n  serverUrl: 'http://localhost:4201'\r\n};\r\n"

var styles string
var name string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			fmt.Errorf("%s", "no name provided")
			return
		}
		generateAngularProject(name)
		generateGoProject(name)
		if err := modifyAngularFiles(); err != nil {
			fmt.Errorf("error conecting go and angular: %s", err.Error())
			return
		}
	},
}

func init() {
	generateCmd.Flags().StringVarP(&styles, "style", "s", "", "Set CSS Preprocessor To SCSS/LESS")
	generateCmd.Flags().StringVarP(&name, "name", "n", "", "Set Project Name")
	generateCmd.MarkFlagRequired("name")
	RootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func generateAngularProject(name string) {
	runExternalCmd("ng", []string{
		"new",
		name,
	})
}

func generateGoProject(name string) {
	fmt.Println(BlueFunc()("Generating Go Files"))
	var mainFileContent = "package main\r\n\r\nimport (\r\n  \"github.com/gorilla/mux\"\r\n  \"net/http\"\r\n  \"os\"\r\n  \"log\"\r\n\"" + name + "/src/server/utils\"\r\n  \"fmt\"\r\n  \"github.com/rs/cors\"\r\n)\r\n\r\nfunc main() {\r\n  r := mux.NewRouter()\r\n\r\n  r.HandleFunc(\"/hello-world\", helloWorld)\r\n\r\n  // Solves Cross Origin Access Issue\r\n  c := cors.New(cors.Options{\r\n    AllowedOrigins: []string{\"http://localhost:4200\"},\r\n  })\r\n  handler := c.Handler(r)\r\n\r\n  srv := &http.Server{\r\n    Handler: handler,\r\n    Addr:    \":\" + os.Getenv(\"PORT\"),\r\n  }\r\n\r\n  log.Fatal(srv.ListenAndServe())\r\n}\r\n\r\nfunc helloWorld(w http.ResponseWriter, r *http.Request) {\r\n  var data = struct {\r\n    Title string `json:\"title\"`\r\n  }{\r\n    Title: \"Golang + Angular Starter Kit\",\r\n  }\r\n\r\n  jsonBytes, err := utils.StructToJson(data); if err != nil {\r\n    fmt.Print(err)\r\n  }\r\n\r\n  w.Header().Set(\"Content-Type\", \"application/json\")\r\n  w.Write(jsonBytes)\r\n  return\r\n}"
	src := "./" + name + "/src/"
	os.Mkdir(src+"server", 0700)
	os.Chdir(src + "server")

	f, err := os.Create("main.go")
	if err != nil {
		fmt.Errorf("error generating project: %s", err.Error())
	}
	f.WriteString(mainFileContent)
	f.Close()

	os.Mkdir("utils", 0700)
	os.Chdir("utils")

	f2, err := os.Create("json.go")
	if err != nil {
		fmt.Errorf("error generating project: %s", err.Error())
	}
	f2.WriteString(jsonFileContent)

	fmt.Println(BlueFunc()("Formatting Go Files"))
	os.Chdir("../")
	runExternalCmd("go", []string{"fmt", "./..."})
	fmt.Println(BlueFunc()("Done"))
}

func modifyAngularFiles() error {
	fmt.Println("Connecting Go and Angular")

	os.Chdir("../../")

	if err := ioutil.WriteFile("./src/app/app.component.ts", []byte(appComponentContent), 0644); err != nil {
		return err
	}

	if err := ioutil.WriteFile("./src/app/hello-world.service.ts", []byte(helloWorldServiceContent), 0644); err != nil {
		return err
	}

	if err := ioutil.WriteFile("./src/app/app.module.ts", []byte(appModuleContent), 0644); err != nil {
		return err
	}

	if err := ioutil.WriteFile("./src/environments/environment.ts", []byte(environmentContent), 0644); err != nil {
		return err
	}

	return nil
}
