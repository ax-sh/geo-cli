package main

import (
	_ "embed"
	"geo/pkg/country"
	"geo/pkg/tui"
)

//go:embed VERSION
var version string

func main() {
	tui.FilterPhone(func(countryCode string) string {
		fil := country.FilterCountryByCountryCodeDataFrame(countryCode)
		sel := fil.Drop([]string{"EquivalentFipsCode", "Postal Code Regex", "Postal Code Format"}).Drop("Area(in sq km)")
		result := tui.PrintDataframe(sel)
		return result.String()
	})
	//s := "gopher"
	//fmt.Printf("Hello and welcome, %s! version [%s]\n", s, version)
	//countryCode := "41"
	//

	////sel := fil.Select([]string{"ISO", "ISO3", "ISO-Numeric", "fips",
	////	"Country",
	////	"Capital",
	////	//,
	////	"Population",
	////	//"Continent	tld	CurrencyCode",
	////	"CurrencyName",
	////	"Phone",
	////	//"Postal Code Format	Postal Code Regex",
	////	"Languages",
	////	"geonameid",
	////	"neighbours",
	////	//"EquivalentFipsCode",
	////})

	//fmt.Println(result)

}
