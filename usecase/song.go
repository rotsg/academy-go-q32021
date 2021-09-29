package usecase

import (
	"errors"

	"github.com/rotsg/bootcamp_challenge/model"
	"github.com/rotsg/bootcamp_challenge/service"
)

const filePath = "data/songs.csv"

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
