package main

import (
	_ "embed"
	"fmt"
	"geo/pkg/country"
	"geo/pkg/tui"
)

//go:embed VERSION
var version string

func main() {
	s := "gopher"
	fmt.Printf("Hello and welcome, %s! version [%s]\n", s, version)
	countryCode := "33"
	fil := country.FilterCountryByCountryCodeDataFrame(countryCode)
	result := tui.PrintDataframe(fil)
	println(result)

}
