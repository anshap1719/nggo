# NGGO (WIP)

A CLI tool for working with Angular + Go projects. Currently under development.

## Prerequisites
- You must have Go installed and GOPATH & GOBIN setup properly
- You must have angular-cli installed (any version)

**NOTE:** *Generated angular project version will be based on your angular-cli version*

## Installation
Run `go get -u github.com/anshap1719/nggo` to install the tool globally.

## Generate New Project
Run `nggo generate -n="{project_name}` to generate a new project. You must provide the name of the project using either -n or --name.

## Project Generator Configuration
This tool supports all angular flags and options that can be used with `ng new`. Simply provide an additional argument to `nggo generate` with flag `--ng`. Ex. `nggo generate -n="new-project" --ng="--skip-install --style=scss"`

## Install Dependencies
By default, the angular dependencies will be installed automatically on generation of project (not if you use --skip-install). You may run `nggo install` inside your project folder to install node + go dependencies. This is a crucial step as the generated go project has dependencies apart from standard library

**Note:** *Projects must be generatead in any subdirectory of your GOPATH*

## Start Development Server
Run `nggo serve` inside project folder to run the webpack dev server for angular and gin live server for go.

## Authors

* **Anshul Sanghi** - [Anshap1719](https://github.com/anshap1719)

## License

This project is licensed under the Apache License - see the [LICENSE.md](LICENSE) file for details.
