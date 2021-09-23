package usecases

import (
	"errors"

	"github.com/rotsg/first_deliverable/models"
	"github.com/rotsg/first_deliverable/services"
)

const filePath = "data/FEDEX.csv"

func GetField(id int) (models.Data, error) {
	data, err := services.ReadCsvFile(filePath)
	if err == nil {
		field, ok := data[id]
		if ok {
			return field, nil
		}
		return models.Data{}, errors.New("field not found")
	}
	return models.Data{}, err
}
