package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	controller controller
}

type controller interface {
	Getter
	Updater
}

type Getter interface {
	GetUniversity(w http.ResponseWriter, r *http.Request)
}

type Updater interface {
	UpdateUniversities(w http.ResponseWriter, r *http.Request)
}

func New(c controller) Route {
	return Route{controller: c}
}

func (r Route) SetupRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/universities/{id}", r.controller.GetUniversity).Methods("GET")
	router.HandleFunc("/universities/", r.controller.UpdateUniversities).Methods("GET")
	return router
}
