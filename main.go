package main

import (
	_ "embed"
	"geo/cmd"
)

//go:embed VERSION
var version string

func main() {

	cmd.Execute(version)

}
