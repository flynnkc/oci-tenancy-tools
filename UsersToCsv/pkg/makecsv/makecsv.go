package makecsv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

var (
	debug bool
	file  string
)

// MapToCsv takes a map with string keys and slice values. Keys are used for headers and values become row data.
func MapToCsv(m map[string][]string) {
	// Initialize array
	spreadsheet := initArray(sizeArray(m))
	headers := makeHeaders(m)
	spreadsheet[0] = headers

	for col, head := range headers {
		for row := 1; row < len(spreadsheet); row++ {
			if row < len(m[head]) {
				spreadsheet[row][col] = m[head][row]
			} else {
				spreadsheet[row][col] = ""
			}
		}
	}

	f, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	//	defer log.Fatal(f.Close())

	c := csv.NewWriter(f)
	log.Fatal(c.WriteAll(spreadsheet))
	fmt.Println("Finished!")
}

// SetEnvironment is required to set variables shared from main package
func SetEnvironment(b bool, f string) {
	debug = b
	file = f
}

// makeHeaders returns map keys as a slice of type string
func makeHeaders(m map[string][]string) []string {
	var s []string
	for k := range m {
		s = append(s, k)
	}
	return s
}

// sliceMaxLength returns an int that is the max length of a map of slices
func sizeArray(m map[string][]string) (int, int) {
	row := 0
	for _, v := range m {
		if len(v) > row {
			row = len(v)
		}
	}
	col := len(m)
	return row, col
}

// initArray is going to return a blank, initialized 2d string array
func initArray(row, col int) [][]string {
	s := make([][]string, row)
	for i := range s {
		s[i] = make([]string, col)
	}

	return s
}
