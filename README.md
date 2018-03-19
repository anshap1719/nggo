# NGGO (WIP)

A CLI tool for working with Angular + Go projects. Currently under development.

## Prerequisites
- You must have Go installed and GOPATH setup properly
- You must have angular-cli installed (any version)

## Installation
Run `go get github.com/anshap1719/nggo` to install the tool globally.

## Generate New Project
Run `nggo generate -n="{project_name}` to generate a new project. You must provide the name of the project using either -n or --name.
**Note:** *Projects must be generatead in any subdirectory of your GOPATH*

## Start Development Server
Run `nggo serve` inside project folder to run the webpack dev server for angular and gin live reload utility for go.

## Authors

* **Anshul Sanghi** - [Anshap1719](https://github.com/anshap1719)

## License

This project is licensed under the Apache License - see the [LICENSE.md](LICENSE.md) file for details.
