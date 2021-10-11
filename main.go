package main

import (
	"net/http"

	"github.com/rotsg/academy-go-q32021/controller"
	"github.com/rotsg/academy-go-q32021/repo"
	"github.com/rotsg/academy-go-q32021/route"
	"github.com/rotsg/academy-go-q32021/usecase"
)

const (
	port     = ":8080"
	filePath = "data/songs.csv"
)

func main() {
	songsRepo := repo.Repo{FilePath: filePath}
	songsUseCase := usecase.New(songsRepo)
	songsController := controller.New(songsUseCase)
	songsRoute := route.New(songsController)
	http.ListenAndServe(port, songsRoute.SetupRoutes())
}
