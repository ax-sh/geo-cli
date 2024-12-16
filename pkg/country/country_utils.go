package country

import "github.com/go-gota/gota/dataframe"

func MoveColumns(df dataframe.DataFrame) dataframe.DataFrame {
	// Column to move
	columnsToMove := []string{"Country", "Capital"}
	// Get all column names
	columns := df.Names()

	// Create a new column order with the target columns at the front
	newOrder := make([]string, 0)
	seen := make(map[string]bool)

	// Add the columns to move at the start
	for _, col := range columnsToMove {
		newOrder = append(newOrder, col)
		seen[col] = true
	}

	// Add the remaining columns in their original order
	for _, col := range columns {
		if !seen[col] {
			newOrder = append(newOrder, col)
		}
	}

	df = df.Select(newOrder)
	return df
}

func NormalizeCountryDataFrame(df dataframe.DataFrame) dataframe.DataFrame {
	fil := df.Drop([]string{"EquivalentFipsCode", "Postal Code Regex", "Postal Code Format"}).Drop("Area(in sq km)")
	return MoveColumns(fil)
}
