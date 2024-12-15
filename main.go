package main

import (
	_ "embed"
	"fmt"
	"geo/pkg/table"
)

//go:embed VERSION
var version string

func main() {
	s := "gopher"
	fmt.Printf("Hello and welcome, %s! version [%s]\n", s, version)
	table.TBBB()
}
