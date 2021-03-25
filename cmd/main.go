package main

import (
	"log"

	greendo "github.com/Dponya/GreenDo"
	"github.com/Dponya/GreenDo/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(greendo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
