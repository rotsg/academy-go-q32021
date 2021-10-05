package route

import (
	"github.com/rotsg/academy-go-q32021/controller"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/songs/{id}", controller.GetSong).Methods("GET")
	return router
}
