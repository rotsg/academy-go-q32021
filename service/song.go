package service

import (
	"errors"

	"github.com/rotsg/academy-go-q32021/model"
)

type Service struct {
	repo Repo
}

type Repo interface {
	Getter
}

type Getter interface {
	GetSongs() (map[int]model.Song, error)
}

func New(r Repo) Service {
	return Service{repo: r}
}

// GetSong - Finds a song by its id from a map of songs.
func (s Service) GetSong(id int) (model.Song, error) {
	data, err := s.repo.GetSongs()
	if err != nil {
		return model.Song{}, err
	}
	field, ok := data[id]
	if ok {
		return field, nil
	}
	return model.Song{}, errors.New("song not found")
}
