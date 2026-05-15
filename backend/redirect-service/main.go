package main

import (
	"log"

	"redirect-service/internal/repository"
	"redirect-service/internal/router"
	"redirect-service/internal/service"
	"redirect-service/internal/util"
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