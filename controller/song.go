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
	service service
}

type service interface {
	Getter
}

type Getter interface {
	GetSong(id int) (model.Song, error)
}

func New(s service) Controller {
	return Controller{service: s}
}

// GetSong - Sends a http request to get a song by its id and return a json.
func (c Controller) GetSong(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "response: invalid id")
		log.Println("ERROR: ", err)
		return
	}
	song, err := c.service.GetSong(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "response: %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(song)
}
