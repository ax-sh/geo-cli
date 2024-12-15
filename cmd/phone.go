package cmd

import (
	"geo/pkg/country"
	"geo/pkg/tui"
	"github.com/spf13/cobra"
)

// geoSubCmd represents the icon command
var geoSubCmd = &cobra.Command{
	Use:   "phone",
	Short: "Filter country by phone country code",
	Run: func(cmd *cobra.Command, args []string) {
		callback := func(countryCode string) string {
			if len(countryCode) == 0 {
				return ""
			}
			fil := country.FilterCountryByCountryCodeDataFrame(countryCode)
			sel := fil.Drop([]string{"EquivalentFipsCode", "Postal Code Regex", "Postal Code Format"}).Drop("Area(in sq km)")
			result := tui.PrintDataframe(sel)
			return result.String()
		}
		tui.FilterPhone(callback)
	},
}

func init() {
	rootCmd.AddCommand(geoSubCmd)
}
