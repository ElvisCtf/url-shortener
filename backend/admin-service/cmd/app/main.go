package main

import (
	"example.com/admin-service/internal/features/auth"
	"example.com/admin-service/internal/util"
	"example.com/admin-service/internal/util/database"

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

	db, err := database.InitGormDB(config)
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
