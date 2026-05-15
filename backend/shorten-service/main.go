package main

import (
	"log"

	"shorten-service/internal/repository"
	"shorten-service/internal/router"
	"shorten-service/internal/service"
	"shorten-service/internal/util"
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
