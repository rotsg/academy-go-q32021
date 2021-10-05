package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/rotsg/academy-go-q32021/usecase"

	"github.com/gorilla/mux"
)

// GetSong - Sends a http request to get a song by its id and return a json.
func GetSong(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "response: invalid id")
		log.Println("ERROR: ", err)
		return
	}
	song, err := usecase.GetSong(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "response: %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(song)
}
