package main

import (
	"log"

	"example.com/shorten-service/internal/repository"
	"example.com/shorten-service/internal/router"
	"example.com/shorten-service/internal/service"
	"example.com/shorten-service/internal/util"
)

func main() {
	addr := util.Env("ADDR", ":8080")
	baseURL := util.Env("BASE_URL", "http://localhost:8081")
	storage := util.Env("STORAGE", "")

	repo := repository.NewRepo(storage)
	service := service.NewShorten(baseURL, repo)
	router := router.SetupRouter(service)

	if err := router.Run(addr); err != nil {
		log.Fatalf("startup service failed, err: %v\n", err)
	}
}
