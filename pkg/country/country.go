package country

import (
	"encoding/csv"
	"fmt"
	"github.com/go-gota/gota/dataframe"

	"log"
	"os"
)

func LoadTsvFileAsCsv(filePath string) *csv.Reader {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal("Error: LoadCountryDataFrame", err)
	}
	csvReader := csv.NewReader(file)

	csvReader.Comma = '\t'
	csvReader.Comment = '#'
	csvReader.FieldsPerRecord = -1

	return csvReader
}
func ReadCountryAsDataFrame() dataframe.DataFrame {
	csv := LoadTsvFileAsCsv("countryInfo.tsv")
	records := make([][]string, 0)
	for {
		row, err := csv.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Printf("Error reading row: %v\n", err)
			continue
		}
		length := len(row)

		if length != 18 {
			continue
		}

		records = append(records, row)
	}
	return dataframe.LoadRecords(records)
}

func LoadCountryDataFrame() dataframe.DataFrame {
	df := ReadCountryAsDataFrame()

	sub := df.Subset([]int{0, 2})

	fmt.Println(sub.String())
	fmt.Println((df.Names()))
	return dataframe.DataFrame{}
}
