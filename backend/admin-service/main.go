package main

import (
	"github.com/gin-gonic/gin"

	"example.com/admin-service/internal/features/auth"
	"example.com/admin-service/internal/features/link"
	"example.com/shared"
	"example.com/shared/postgres"
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
	authRepo := auth.NewAuthRepository(db)
	authService := auth.NewAuthService(authRepo, config)
	authController := auth.NewAuthController(authService, config)
	authRouter := auth.NewAuthRouter(authController)

	linkRepo := link.NewLinkRepository(db)
	linkService := link.NewLinkService(linkRepo, config)
	linkController := link.NewLinkController(linkService, config)
	linkRouter := link.NewLinkRouter(linkController)

	// Setup Routes
	authRouter.SetupRoutes(&router.RouterGroup, config)
	linkRouter.SetupRoutes(&router.RouterGroup, config)

	router.Run(addr)
}
