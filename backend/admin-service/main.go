package main

import (
	"example.com/admin-service/internal/features/auth"
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
	authRepo := auth.NewAuthRepository(db)
	authService := auth.NewAuthService(authRepo, config)
	authController := auth.NewAuthController(authService, config)
	authRouter := auth.NewAuthRouter(authController)

	// Setup Routes
	authRouter.SetupRoutes(&router.RouterGroup, config)

	router.Run(addr)
}
