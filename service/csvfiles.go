package service

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/rotsg/academy-go-q32021/model"
)

const message = "something went wrong"

// GetCsvData - Reads a csv file from the path parameter and returns a map of songs.
func GetCsvData(filePath string) (map[int]model.Song, error) {
	m := make(map[int]model.Song)
	csvfile, err := os.Open(filePath)
	if err != nil {
		log.Println("ERROR: ", err)
		return m, errors.New(message)
	}
	r := csv.NewReader(csvfile)
	for {
		data, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("ERROR: ", err)
			return m, errors.New(message)
		}
		id, err := strconv.Atoi(data[0])
		if err != nil {
			log.Println("ERROR: ", err)
			return m, errors.New(message)
		}
		m[id] = model.Song{
			Id:     id,
			Name:   data[1],
			Artist: data[2],
		}
	}
	return m, nil
}
