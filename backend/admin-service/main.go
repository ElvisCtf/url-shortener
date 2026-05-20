package main

import (
	"admin-service/internal/features/auth"
	"admin-service/internal/util"
	"github.com/gin-gonic/gin"
)

func main() {
	addr := util.Env("ADDR", ":8082")

	router := gin.Default()

	authService := auth.NewAuthService()
	authController := auth.NewAuthController(authService)
	authRouter := auth.NewAuthRouter(authController)
	authRouter.SetupRoutes(&router.RouterGroup)

	router.Run(addr)
}
