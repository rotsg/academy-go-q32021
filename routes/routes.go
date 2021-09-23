package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rotsg/first_deliverable/controllers"
)

func New() {
	router := mux.NewRouter()

	//specify endpoints
	router.HandleFunc("/getField/{id}", controllers.GetField).Methods("GET")

	//start and listen to requests
	http.ListenAndServe(":8080", router)
}
