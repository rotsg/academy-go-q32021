package usecase

import (
	"errors"

	"github.com/rotsg/first_deliverable/model"
	"github.com/rotsg/first_deliverable/service"
)

const filePath = "data/FEDEX.csv"

func GetField(id int) (model.Data, error) {
	data, err := service.ReadCsvFile(filePath)
	if err != nil {
		return model.Data{}, err
	}
	field, ok := data[id]
	if ok {
		return field, nil
	}
	return model.Data{}, errors.New("field not found")
}
