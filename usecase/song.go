package usecase

import (
	"errors"

	"github.com/rotsg/academy-go-q32021/model"
	"github.com/rotsg/academy-go-q32021/service"
)

const filePath = "data/songs.csv"

// GetSong - Finds a song by its id from a map of songs.
func GetSong(id int) (model.Song, error) {
	data, err := service.GetCsvData(filePath)
	if err != nil {
		return model.Song{}, err
	}
	field, ok := data[id]
	if ok {
		return field, nil
	}
	return model.Song{}, errors.New("song not found")
}
