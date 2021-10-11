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
}

type Getter interface {
	GetSong(w http.ResponseWriter, r *http.Request)
}

func New(c controller) Route {
	return Route{controller: c}
}

func (r Route) SetupRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/songs/{id}", r.controller.GetSong).Methods("GET")
	return router
}
