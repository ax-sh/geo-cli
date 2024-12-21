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
				return "Type to filter"
			}
			fil := country.FilterCountryByCountryCodeDataFrame(countryCode)
			sel := country.DropUselessCountryColumn(fil).
				Drop("Area(in sq km)")
			sel = country.MoveImportantColumnsToStart(sel)
			sel = country.MoveColumnsToStart(sel, "Phone")
			result := tui.PrintDataframe(sel)
			return result.String()
		}
		tui.FilterPhone(callback)
	},
}

func init() {
	rootCmd.AddCommand(geoSubCmd)
}
