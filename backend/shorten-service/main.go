package main

import (
	"example.com/shared"
	"example.com/shared/postgres"
	"example.com/shorten-service/internal/features/shorten"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := shared.LoadConfig()
	if err != nil {
		panic("Failed to load config: " + err.Error())
	}
	addr := config.Addr
	gin.SetMode(config.GinMode)

	router := gin.Default()

	db, err := postgres.InitGormDB(config)
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	// Dependency Injection
	shortenRepo := shorten.NewShortenRepository(db)
	shortenService := shorten.NewShortenService(shortenRepo, config)
	shortenController := shorten.NewShortenController(shortenService, config)
	shortenRouter := shorten.NewShortenRouter(shortenController)

	// Setup Routes
	shortenRouter.SetupRoutes(&router.RouterGroup, config)

	router.Run(addr)
}
