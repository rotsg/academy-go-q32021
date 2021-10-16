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
	Updater
}

type Getter interface {
	GetUniversities() (map[int]model.University, error)
}

type Updater interface {
	UpdateUniversities() error
}

func New(r repo) UseCase {
	return UseCase{repo: r}
}

// GetUniversity - Finds a university by its id from a map of universities.
func (u UseCase) GetUniversity(id int) (model.University, error) {
	universities, err := u.repo.GetUniversities()
	if err != nil {
		return model.University{}, err
	}
	university, ok := universities[id]
	if ok {
		return university, nil
	}
	return model.University{}, errors.New("university not found")
}

// UpdateUniversities - Updates the csv file of universities.
func (u UseCase) UpdateUniversities() error {
	err := u.repo.UpdateUniversities()
	if err != nil {
		return err
	}
	return nil
}
