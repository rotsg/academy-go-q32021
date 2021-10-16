package universities

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/rotsg/academy-go-q32021/model"
)

const (
	message = "something went wrong"
	api     = "http://universities.hipolabs.com/search?country=Mexico"
)

// GetUniversities - Consumes the external API and returns an array of universities.
func GetUniversities() ([]model.University, error) {
	var universities []model.University

	resp, err := http.Get(api)
	if err != nil {
		log.Println("ERROR: ", err)
		return universities, errors.New(message)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(bodyBytes, &universities)
	if err != nil {
		log.Println("ERROR: ", err)
		return universities, errors.New(message)
	}
	return universities, nil
}
