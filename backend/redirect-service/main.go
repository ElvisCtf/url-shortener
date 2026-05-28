package main

import (
	"example.com/redirect-service/internal/features/redirect"
	"example.com/shared"
	"example.com/shared/postgres"
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
	redirectRepo := redirect.NewRedirectRepository(db)
	redirectService := redirect.NewRedirectService(redirectRepo)
	redirectController := redirect.NewRedirectController(redirectService, config)
	redirectRouter := redirect.NewRedirectRouter(redirectController)

	// Setup Routes
	redirectRouter.SetupRoutes(&router.RouterGroup, config)

	router.Run(addr)
}
