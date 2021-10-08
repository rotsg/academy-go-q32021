package usecase

import (
	"errors"

	"github.com/rotsg/academy-go-q32021/model"
)

type UseCase struct {
	repo repo
}

type repo interface {
	Getter
}

type Getter interface {
	GetSongs() (map[int]model.Song, error)
}

func New(r repo) UseCase {
	return UseCase{repo: r}
}

// GetSong - Finds a song by its id from a map of songs.
func (u UseCase) GetSong(id int) (model.Song, error) {
	songs, err := u.repo.GetSongs()
	if err != nil {
		return model.Song{}, err
	}
	song, ok := songs[id]
	if ok {
		return song, nil
	}
	return model.Song{}, errors.New("song not found")
}
