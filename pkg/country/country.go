package country

import (
	"github.com/go-gota/gota/dataframe"
	"log"
	"os"
)

func LoadCountryDataFrame() dataframe.DataFrame {
	file, err := os.Open("countryInfo.tsv")

	if err != nil {
		log.Fatal("Error: LoadCountryDataFrame", err)
	}
	//defer file.Close()
	df := dataframe.ReadCSV(file,
		dataframe.WithDelimiter('\t'),
		dataframe.HasHeader(false),
		dataframe.WithComments('#'),
		dataframe.Names("a", "b"),
	)
	return df
}
