package cmd

import (
	"geo/pkg/country"
	"geo/pkg/tui"
	"github.com/spf13/cobra"
)

func parseArgs(args []string) string {
	var tld string
	if len(args) == 0 {
		println("Please specify tld code")
		tld = "us"
	} else {
		tld = args[0]
	}
	return tld
}

// tldSubCmd represents the icon command
var tldSubCmd = &cobra.Command{
	Use:   "tld",
	Short: "Filter country by phone country code",
	Run: func(cmd *cobra.Command, args []string) {
		tld := parseArgs(args)
		println("Top level domain", tld)

		callback := func(tld string) string {
			fil := country.FilterCountryByTLDDataFrame(tld)
			sel := country.NormalizeCountryDataFrame(fil)
			sel = sel.Drop([]string{"ISO", "ISO3", "ISO-Numeric"}).
				Drop("neighbours").
				Drop("Languages")

			result := tui.PrintDataframe(sel)
			return result.String()
		}
		s := callback(tld)
		println(s)
		//tui.FilterTLD(callback)
	},
}

func init() {
	rootCmd.AddCommand(tldSubCmd)
}
