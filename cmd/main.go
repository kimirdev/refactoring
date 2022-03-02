package main

import (
	"log"
	"refactoring"
	"refactoring/handler"
	"refactoring/repository"
	"refactoring/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(refactoring.Server)

	if err := srv.Run("3333", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
