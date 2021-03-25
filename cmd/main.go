package main

import (
	"log"

	greendo "github.com/Dponya/GreenDo"
	"github.com/Dponya/GreenDo/pkg/handler"
	"github.com/Dponya/GreenDo/pkg/repository"
	"github.com/Dponya/GreenDo/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(greendo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
