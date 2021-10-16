package usecase

import (
	"errors"
	"testing"

	"github.com/rotsg/academy-go-q32021/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var universities = map[int]model.University{
	1: {
		Id:      1,
		Name:    "CETYS Universidad",
		Country: "Mexico",
	},
	2: {
		Id:      2,
		Name:    "Centro de Estudios Universitarios Monterrey",
		Country: "Country",
	},
}

var university = model.University{
	Id:      1,
	Name:    "CETYS Universidad",
	Country: "Mexico",
}

type MockUniversitiesRepo struct {
	mock.Mock
}

func (ur MockUniversitiesRepo) GetUniversities() (map[int]model.University, error) {
	args := ur.Called()
	return args.Get(0).(map[int]model.University), args.Error(1)
}

func (ur MockUniversitiesRepo) UpdateUniversities() error {
	args := ur.Called()
	return args.Error(0)
}

func TestSongsUseCase_GetUniversity(t *testing.T) {
	testCases := []struct {
		name               string
		response           map[int]model.University
		expectedUniversity model.University
		hasError           bool
		error              error
		universityId       int
	}{
		{
			name:               "error in repo",
			response:           nil,
			expectedUniversity: model.University{},
			hasError:           true,
			error:              errors.New("something went wrong"),
			universityId:       1,
		},
		{
			name:               "found",
			response:           universities,
			expectedUniversity: university,
			hasError:           false,
			error:              nil,
			universityId:       1,
		},
		{
			name:               "not found",
			response:           universities,
			expectedUniversity: model.University{},
			hasError:           true,
			error:              errors.New("song not found"),
			universityId:       3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := MockUniversitiesRepo{}
			mock.On("GetUniversities").Return(tc.response, tc.error)
			us := New(mock)
			university, err := us.GetUniversity(tc.universityId)
			assert.EqualValues(t, university, tc.expectedUniversity)
			if tc.hasError {
				assert.EqualError(t, err, tc.error.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestSongsUseCase_UpdateUniversities(t *testing.T) {
	testCases := []struct {
		name     string
		hasError bool
		error    error
	}{
		{
			name:     "error in repo",
			hasError: true,
			error:    errors.New("something went wrong"),
		},
		{
			name:     "updated",
			hasError: false,
			error:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := MockUniversitiesRepo{}
			mock.On("UpdateUniversities").Return(tc.error)
			us := New(mock)
			err := us.UpdateUniversities()
			if tc.hasError {
				assert.EqualError(t, err, tc.error.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
