package filter_phone

import (
	"geo/pkg/country"
	"geo/pkg/tui"
)

func FilterPhoneTui() {
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
}
