package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/rotsg/academy-go-q32021/model"

	"github.com/gorilla/mux"
)

type Controller struct {
	useCase useCase
}

type useCase interface {
	Getter
	Updater
}

type Getter interface {
	GetUniversity(id int) (model.University, error)
}

type Updater interface {
	UpdateUniversities() error
}

func New(us useCase) Controller {
	return Controller{useCase: us}
}

// GetUniversity - Sends a http request to get a university by its id and return a json.
func (c Controller) GetUniversity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "response: invalid id")
		log.Println("ERROR: ", err)
		return
	}
	song, err := c.useCase.GetUniversity(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "response: %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(song)
}

// UpdateUniversities - Sends a http request to update the csv file of universities.
func (c Controller) UpdateUniversities(w http.ResponseWriter, r *http.Request) {
	err := c.useCase.UpdateUniversities()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "response: %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "response: the csv file of universities was successfully updated")
}
