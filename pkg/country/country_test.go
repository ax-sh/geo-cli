package country

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"strings"
	"testing"
)

func TestSanity(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(123, 123, "they should be equal")
}

func TestCsvCountry(t *testing.T) {
	csvStr := `
Country,Date,Age,Amount,Id
"United States",2012-02-01,50,112.1,01234
"United States",2012-02-01,32,321.31,54320
"United Kingdom",2012-02-01,17,18.2,12345
"United States",2012-02-01,32,321.31,54320
"United Kingdom",2012-02-01,NA,18.2,12345
"United States",2012-02-01,32,321.31,54320
"United States",2012-02-01,32,321.31,54320
Spain,2012-02-01,66,555.42,00241
`
	df := dataframe.ReadCSV(strings.NewReader(csvStr))
	fmt.Println(df)
}
func TestCountryJson(t *testing.T) {
	df := LoadCountryDataFrame()
	fmt.Println("[[[TestCountryJson]]]", df)
}

func TestCountryRows(t *testing.T) {
	df := ReadCountryAsDataFrame()
	length := len(df.Records())
	assert.Equal(t, 253, length)
}

func TestCountryTsv(t *testing.T) {
	file, err := os.Open("countryInfo.tsv")

	if err != nil {
		log.Fatal("Error: LoadCountryDataFrame", err)
	}
	df := dataframe.ReadCSV(file,
		dataframe.HasHeader(true),
		dataframe.WithComments('#'),
		dataframe.NaNValues([]string{""}),
		dataframe.Names("D"),
	)

	//
	//// Basic DataFrame operations
	//fmt.Println("DataFrame Shape:", df.Dims())
	fmt.Println("Column Names:", df.Names())

	//// Print first few rows
	sub := df.Subset([]int{0, 2})
	fmt.Println("Subset:", sub)
	//fmt.Println(df.String())
}

func TestTsv(t *testing.T) {
	// Using with a string reader
	csvString := "ISO\tISO3\tISO-Numeric\tfips\tCountry\tCapital\tArea(in sq km)\tPopulation\tContinent\ttld\tCurrencyCode\tCurrencyName\tPhone\tPostal Code Format\tPostal Code Regex\tLanguages\tgeonameid\tneighbours\tEquivalentFipsCode\nAD\tAND\t020\tAN\tAndorra\tAndorra la Vella\t468\t77006\tEU\t.ad\tEUR\tEuro\t376\tAD###\t^(?:AD)*(\\d{3})$\tca\t3041565\tES,FR"
	df := dataframe.ReadCSV(strings.NewReader(csvString))
	fmt.Println("TestTsv", df)

}
func TestFilterPhone(t *testing.T) {
	result := FilterCountryByCountryCode("41")
	for _, name := range result.Array() {
		fmt.Println(name.Get("Phone"), "222")

	}

	//PrintJSONTable(rows)
	//
	//pp.Println(result.Value())
	//
	//fmt.Println(len(value.Array()), "$$$$<<<")
}
