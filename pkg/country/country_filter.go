package country

import (
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"strings"
)

func FilterCountryByTLDDataFrame(countryCode string) dataframe.DataFrame {
	df := ReadCountryAsDataFrame()
	stringContains := func(substring string) func(el series.Element) bool {
		return func(el series.Element) bool {
			if el.Type() == series.String {
				if val, ok := el.Val().(string); ok {
					return strings.Contains(val, substring)
				}
			}
			return false
		}
	}
	fil := df.Filter(dataframe.F{Colname: "tld", Comparator: series.CompFunc, Comparando: stringContains(countryCode)})

	//fil := df.Filter(dataframe.F{Colname: "Phone", Comparator: series.Eq, Comparando: countryCode})
	return fil
}
