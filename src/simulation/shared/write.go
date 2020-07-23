package shared

import (
	"os"
	"encoding/csv"
	"strconv"
)

type Record struct {
    ID   int
    Type string
    Year string
}



func write_csv(fname string, headers []string){
	m := make(map[int]Record)

	file, _ := os.Create(fname + ".csv")
	// checkError("Error:", err)
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// // define column headers
	// headers := []string{
	// 	"id",
	// 	"type",
	// 	"year",
	// }

	// write column headers
	writer.Write(headers)

	var idString string

	for key := range m {

		r := make([]string, 0, 1+len(headers)) // capacity of 4, 1 + the number of properties your struct has & the number of column headers you are passing

		// convert the Record.ID to a string in order to pass into []string
		idString = strconv.Itoa(m[key].ID)

		r = append(
			r,
			idString,
			m[key].Type,
			m[key].Year,
		)

		writer.Write(r)
	}
}