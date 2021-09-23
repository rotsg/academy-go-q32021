package route

import (
	"github.com/gorilla/mux"
	"github.com/rotsg/first_deliverable/controller"
)

func New() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/getField/{id}", controller.GetField).Methods("GET")
	return router
}
