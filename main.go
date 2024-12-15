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
	countryCode := "41"
	fil := country.FilterCountryByCountryCodeDataFrame(countryCode)
	sel := fil.Drop([]string{"EquivalentFipsCode", "Postal Code Regex", "Postal Code Format"}).Drop("Area(in sq km)")

	//sel := fil.Select([]string{"ISO", "ISO3", "ISO-Numeric", "fips",
	//	"Country",
	//	"Capital",
	//	//,
	//	"Population",
	//	//"Continent	tld	CurrencyCode",
	//	"CurrencyName",
	//	"Phone",
	//	//"Postal Code Format	Postal Code Regex",
	//	"Languages",
	//	"geonameid",
	//	"neighbours",
	//	//"EquivalentFipsCode",
	//})
	result := tui.PrintDataframe(sel)
	fmt.Println(result)

}
