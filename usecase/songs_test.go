package usecase

import (
	"errors"
	"testing"

	"github.com/rotsg/academy-go-q32021/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var songs = map[int]model.Song{
	1: {
		Id:     1,
		Name:   "My Universe",
		Artist: "Colplay ft. BTS",
	},
	2: {
		Id:     2,
		Name:   "Boy With Luv",
		Artist: "BTS ft. Halsey",
	},
}

var song = model.Song{
	Id:     1,
	Name:   "My Universe",
	Artist: "Colplay ft. BTS",
}

type MockSongsRepo struct {
	mock.Mock
}

func (sr MockSongsRepo) GetSongs() (map[int]model.Song, error) {
	args := sr.Called()
	return args.Get(0).(map[int]model.Song), args.Error(1)
}

func TestSongsUseCase_GetSong(t *testing.T) {
	testCases := []struct {
		name         string
		response     map[int]model.Song
		expectedSong model.Song
		hasError     bool
		error        error
		songId       int
	}{
		{
			name:         "error in repo",
			response:     nil,
			expectedSong: model.Song{},
			hasError:     true,
			error:        errors.New("something went wrong"),
			songId:       1,
		},
		{
			name:         "found",
			response:     songs,
			expectedSong: song,
			hasError:     false,
			error:        nil,
			songId:       1,
		},
		{
			name:         "not found",
			response:     songs,
			expectedSong: model.Song{},
			hasError:     true,
			error:        errors.New("song not found"),
			songId:       3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := MockSongsRepo{}
			mock.On("GetSongs").Return(tc.response, tc.error)
			us := New(mock)
			song, err := us.GetSong(tc.songId)
			assert.EqualValues(t, song, tc.expectedSong)
			if tc.hasError {
				assert.EqualError(t, err, tc.error.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
