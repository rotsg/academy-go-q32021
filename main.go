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
	filePath = "./data/universities.csv"
)

func main() {
	universitiesRepo := repo.Repo{
		FilePath: filePath,
	}
	universitiesUseCase := usecase.New(universitiesRepo)
	universitiesController := controller.New(universitiesUseCase)
	universitiesRoute := route.New(universitiesController)
	http.ListenAndServe(port, universitiesRoute.SetupRoutes())
}
