package route

import (
	"github.com/gorilla/mux"
	"github.com/rotsg/bootcamp_challenge/controller"
)

func New() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/songs/{id}", controller.GetSong).Methods("GET")
	return router
}
