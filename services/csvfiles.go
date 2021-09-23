package services

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"

	"github.com/rotsg/first_deliverable/models"
)

func ReadCsvFile(filePath string) (map[int]models.Data, error) {
	m := make(map[int]models.Data)
	csvfile, err := os.Open(filePath)
	if err != nil {
		return m, errors.New("couldn't open the csv file")
	}
	r := csv.NewReader(csvfile)
	for {
		data, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return m, errors.New("couldn't read the csv file")
		}
		price, _ := strconv.ParseFloat(data[1], 64)
		id, _ := strconv.Atoi(data[0])
		concept := data[2]
		m[id] = models.Data{id, price, concept}
	}
	return m, nil
}
