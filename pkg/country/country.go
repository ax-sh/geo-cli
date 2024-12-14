package country

import (
	"encoding/csv"
	"fmt"
	"github.com/go-gota/gota/dataframe"

	"log"
	"os"
)

func readCsv() {
	file, err := os.Open("countryInfo.tsv")

	if err != nil {
		log.Fatal("Error: LoadCountryDataFrame", err)
	}
	//defer file.Close()
	csvReader := csv.NewReader(file)
	csvReader.Comma = '\t'
	csvReader.Comment = '#'
	csvReader.FieldsPerRecord = -1

	// Read and print header
	header, err := csvReader.Read()
	if err != nil {
		log.Fatal(err)
	}
	expectedFieldCount := len(header)
	fmt.Printf("Expected field count: %d\n", expectedFieldCount)

	records, err := csvReader.ReadAll()
	//if err != nil {
	//	return dataframe.DataFrame{Err: err}
	//}
	fmt.Printf("Found %d records\n", len(records))
	//for {
	//	row, err := csvReader.Read()
	//	if err != nil {
	//		// Handle EOF and errors
	//		if err.Error() == "EOF" {
	//			break
	//		}
	//		fmt.Printf("Error reading row: %v\n", err)
	//		continue
	//	}
	//	// Process each row
	//	fmt.Printf("Row: %v\n", row)
	//}
	//// Reset reader
	//file.Seek(0, 0)
	//csvReader = csv.NewReader(file)
	//// Inspect problematic line
	//for lineNum := 1; ; lineNum++ {
	//	record, err := csvReader.Read()
	//	if err == io.EOF {
	//		break
	//	}
	//	if err != nil {
	//		if parseErr, ok := err.(*csv.ParseError); ok {
	//			if parseErr.Err == csv.ErrFieldCount {
	//				fmt.Printf("Problem at line %d\n", lineNum)
	//				fmt.Printf("Line contents: %v\n", record)
	//				fmt.Printf("Field count: %d\n", len(record))
	//			}
	//		}
	//	}
	//}
}

func LoadTsvFileAsCsv() *csv.Reader {
	file, err := os.Open("countryInfo.tsv")

	if err != nil {
		log.Fatal("Error: LoadCountryDataFrame", err)
	}
	//defer file.Close()
	csvReader := csv.NewReader(file)

	csvReader.Comma = '\t'
	csvReader.Comment = '#'
	csvReader.FieldsPerRecord = -1
	return csvReader
}
func ReadCountry() dataframe.DataFrame {
	csv := LoadTsvFileAsCsv()
	for {
		row, err := csv.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Printf("Error reading row: %v\n", err)
			continue
		}

		fmt.Printf("Row length: %d, Row: %v\n", len(row), row)
		if len(row) > 17 {
			fmt.Println("Column 18:", row[17])
		} else {
			fmt.Printf("Row too short: %v\n", row)
		}
	}
	return dataframe.DataFrame{}
	//return dataframe.LoadRecords(records)
}

func LoadCountryDataFrame() dataframe.DataFrame {
	df := ReadCountry()
	fmt.Println(df, 555)
	//readCsv()

	//defer file.Close()
	//n := []string{"EquivalentFipsCode"}
	//df := dataframe.ReadCSV(file,
	//	dataframe.WithComments('#'),
	//	dataframe.WithDelimiter('\t'),
	//	dataframe.HasHeader(true),
	//	dataframe.DetectTypes(false),
	//	dataframe.NaNValues(n),
	//	dataframe.WithLazyQuotes(true),
	//
	//	//dataframe.Names("a", "b"),
	//)
	//sub := df.Subset([]int{0, 2})
	//df.
	//	fmt.Println(333, sub.Names())
	////fmt.Println(len(df.Names()))
	return dataframe.DataFrame{}
}
