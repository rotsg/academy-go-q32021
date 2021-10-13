package main

import (
	"github.com/rotsg/academy-go-q32021/repo"
	"github.com/rotsg/academy-go-q32021/universities"
)

const (
	port     = ":8080"
	filePath = "./data/universities.csv"
)

func main() {
	uni := universities.University{}
	newUniversities, _ := uni.GetUniversities()
	songsRepo := repo.Repo{FilePath: filePath}
	songsRepo.UpdateUniversities(newUniversities)
	/*
		songsUseCase := usecase.New(songsRepo)
		songsController := controller.New(songsUseCase)
		songsRoute := route.New(songsController)
		http.ListenAndServe(port, songsRoute.SetupRoutes())*/
}
