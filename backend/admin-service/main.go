package main

import (
	"admin-service/internal/features/auth"
	"admin-service/internal/util"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := util.LoadConfig()
	if err != nil {
		panic("Failed to load config: " + err.Error())
	}
	addr := config.Addr
	gin.SetMode(config.GinMode)

	router := gin.Default()

	// Dependency Injection
	authService := auth.NewAuthService()
	authController := auth.NewAuthController(authService)
	authRouter := auth.NewAuthRouter(authController)

	// Setup Routes
	authRouter.SetupRoutes(&router.RouterGroup)

	router.Run(addr)
}
