package repo

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/rotsg/academy-go-q32021/model"
)

type Repo struct {
	FilePath string
}

const message = "something went wrong"

// GetSongs - Reads a csv file from the path parameter and returns a map of songs.
func (r Repo) GetSongs() (map[int]model.Song, error) {
	m := make(map[int]model.Song)
	csvfile, err := os.Open(r.FilePath)
	if err != nil {
		log.Println("ERROR: ", err)
		return m, errors.New(message)
	}
	nr := csv.NewReader(csvfile)
	for {
		data, err := nr.Read()
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

func (r Repo) UpdateUniversities(universities []model.University) error {
	csvFile, err := os.Create(r.FilePath)
	if err != nil {
		log.Println("ERROR: ", err)
		return errors.New(message)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)

	for index, university := range universities {
		var row []string
		row = append(row, strconv.Itoa(index+1))
		row = append(row, university.Name)
		row = append(row, university.Country)
		row = append(row, university.WebPage[0])
		writer.Write(row)
	}
	writer.Flush()
	return nil
}
