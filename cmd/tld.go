package cmd

import (
	"geo/pkg/country"
	"geo/pkg/tui"
	"github.com/spf13/cobra"
)

// tldSubCmd represents the icon command
var tldSubCmd = &cobra.Command{
	Use:   "tld",
	Short: "Filter country by phone country code",
	Run: func(cmd *cobra.Command, args []string) {
		tld := args[0]
		if tld == "" {
			println("Please specify tld code")
			return
		}
		callback := func(tld string) string {

			fil := country.FilterCountryByTLDDataFrame(tld)
			sel := country.NormalizeCountryDataFrame(fil)
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
