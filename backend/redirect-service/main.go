package main

import (
	"log"

	"example.com/redirect-service/internal/repository"
	"example.com/redirect-service/internal/router"
	"example.com/redirect-service/internal/service"
	"example.com/redirect-service/internal/util"
)

func main() {
	addr := util.Env("ADDR", ":8081")

	repo := repository.NewRepo()
	service := service.NewRedirect(repo)
	router := router.SetupRouter(service)

	if err := router.Run(addr); err != nil {
		log.Fatalf("startup service failed, err: %v\n", err)
	}
}
