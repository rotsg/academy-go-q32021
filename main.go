package main

import (
	"net/http"

	"github.com/rotsg/academy-go-q32021/controller"
	"github.com/rotsg/academy-go-q32021/repo"
	"github.com/rotsg/academy-go-q32021/route"
	"github.com/rotsg/academy-go-q32021/service"
)

const (
	port     = ":8080"
	filePath = "data/songs.csv"
)

func main() {
	r := repo.Repo{FilePath: filePath}
	s := service.New(r)
	c := controller.New(s)
	ro := route.New(c)
	http.ListenAndServe(port, ro.SetupRoutes())
}
