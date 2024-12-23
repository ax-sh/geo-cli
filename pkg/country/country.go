package country

import (
	"bufio"
	"bytes"
	_ "embed"
	"encoding/csv"
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
	"io"
	"log"
	"os"
	"strings"
)

//go:embed countryInfo.tsv
var countryInfo string

func LoadTsvFileAsCsv() *csv.Reader {
	// Create a string reader from the embedded file contents
	reader := strings.NewReader(countryInfo)
	csvReader := csv.NewReader(reader)

	csvReader.Comma = '\t'
	csvReader.Comment = '#'
	csvReader.FieldsPerRecord = -1

	return csvReader
}

// Pad row to a specific length by adding empty strings
func padRowToLength(row []string, targetLength int) []string {
	// Continue adding empty strings until desired length is reached
	for len(row) < targetLength {
		row = append(row, "")
	}
	return row
}

func ReadCountryAsDataFrame() dataframe.DataFrame {
	csv := LoadTsvFileAsCsv()

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
		maxLength := 19
		if length < maxLength {
			row = padRowToLength(row, maxLength)
		}

		records = append(records, row)

	}
	return dataframe.LoadRecords(records)
}

// ReadCountryAsDataFrameMostPerf reads and processes a TSV file into a DataFrame with improved performance and error handling
func ReadCountryAsDataFrameMostPerf(filePath string) (dataframe.DataFrame, error) {
	// Use bufio.Reader for more efficient I/O operations
	file, err := os.Open(filePath)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Create a buffered reader to reduce system calls and improve read performance
	bufferedReader := bufio.NewReader(file)

	// Configure CSV reader with optimized settings
	csvReader := csv.NewReader(bufferedReader)
	csvReader.Comma = '\t'         // TSV delimiter
	csvReader.Comment = '#'        // Ignore comment lines
	csvReader.FieldsPerRecord = 18 // Explicitly set expected number of columns
	csvReader.ReuseRecord = true   // Reuse memory for each record to reduce allocations

	// Pre-allocate slice with estimated capacity to reduce dynamic resizing
	records := make([][]string, 0, 1000)

	// Use a more efficient reading strategy with error checking
	for {
		row, err := csvReader.Read()
		if err != nil {
			// Break the loop on EOF, which is the expected end of file
			if errors.Is(err, io.EOF) {
				break
			}
			// Log or handle other parsing errors
			return dataframe.DataFrame{}, fmt.Errorf("error reading CSV: %w", err)
		}

		// Validate row length more efficiently
		// Removed separate length check as FieldsPerRecord handles this
		records = append(records, row)
	}

	// Convert records to DataFrame
	return dataframe.LoadRecords(records), nil
}

// ProcessCountryData usage with proper error handling
func ProcessCountryData() {
	df, err := ReadCountryAsDataFrameMostPerf("countryInfo.tsv")
	if err != nil {
		log.Fatalf("Failed to load country data: %v", err)
	}
	fmt.Printf("CountryDataFrame: %v\n", df)
}

func LoadCountryDataFrame() dataframe.DataFrame {
	df := ReadCountryAsDataFrame()

	//fil := df.Filter(dataframe.F{Colidx: 0, Colname: "Phone", Comparator: series.Eq, Comparando: "41"})
	//fmt.Printf(">>>>Loading country data: %v\n", fil)

	sub := df.Subset([]int{0, 2})
	////
	//fmt.Println(sub.String())
	//fmt.Println((df.Names()))
	return sub
}
func FilterCountryByCountryCodeDataFrame(countryCode string) dataframe.DataFrame {
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
	fil := df.Filter(dataframe.F{Colname: "Phone", Comparator: series.CompFunc, Comparando: stringContains(countryCode)})

	//fil := df.Filter(dataframe.F{Colname: "Phone", Comparator: series.Eq, Comparando: countryCode})
	return fil
}

func FilterCountryByCountryCode(countryCode string) gjson.Result {
	fil := FilterCountryByCountryCodeDataFrame(countryCode)
	buf := new(bytes.Buffer)
	err := fil.WriteJSON(buf)
	if err != nil {
		log.Printf("Expected success, got error: %v", err)
	}
	json := buf.String()
	result := gjson.Parse(json)

	return result
}
