package repo

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/rotsg/academy-go-q32021/model"
	"github.com/rotsg/academy-go-q32021/universities"
)

type Repo struct {
	FilePath string
}

const message = "something went wrong"

// GetUniversities - Reads a csv file from the path parameter and returns a map of universities.
func (r Repo) GetUniversities() (map[int]model.University, error) {
	universities := make(map[int]model.University)
	csvfile, err := os.Open(r.FilePath)
	if err != nil {
		log.Println("ERROR: ", err)
		return universities, errors.New(message)
	}
	nr := csv.NewReader(csvfile)
	for {
		data, err := nr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("ERROR: ", err)
			return universities, errors.New(message)
		}
		id, err := strconv.Atoi(data[0])
		if err != nil {
			log.Println("ERROR: ", err)
			return universities, errors.New(message)
		}
		universities[id] = model.University{
			Id:      id,
			Name:    data[1],
			Country: data[2],
		}
	}
	return universities, nil
}

// UpdateUniversities -  Saves the result of the external API in the csv file
func (r Repo) UpdateUniversities() error {
	csvFile, err := os.Create(r.FilePath)
	if err != nil {
		log.Println("ERROR: ", err)
		return errors.New(message)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	universities, err := universities.GetUniversities()
	if err != nil {
		log.Println("ERROR: ", err)
		return errors.New(message)
	}

	for index, university := range universities {
		var row []string
		row = append(row, strconv.Itoa(index+1))
		row = append(row, university.Name)
		row = append(row, university.Country)
		writer.Write(row)
	}
	writer.Flush()
	return nil
}
