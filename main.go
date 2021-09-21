package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Data struct {
	price   float64
	concept string
}

func readCsvFile(filePath string) (map[int64]Data, error) {
	m := make(map[int64]Data)
	// Open the file
	csvfile, err := os.Open(filePath)
	if err != nil {
		return m, errors.New("Couldn't open the csv file")
	}
	// Parse the file
	r := csv.NewReader(csvfile)
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return m, errors.New("Couldn't read the csv file")
		}
		// Save data
		price, _ := strconv.ParseFloat(record[1], 64)
		ID, _ := strconv.ParseInt(record[0], 10, 64)
		concept := record[2]
		m[ID] = Data{price, concept}
	}
	return m, nil
}

func main() {
	records, err := readCsvFile("csv/FEDEX.csv")
	if err != nil {
		fmt.Println(err)
	} else {
		for key, value := range records {
			fmt.Println("Key:", key, "Value:", value)
		}
	}

}
