package router

import (
	"example.com/redirect-service/internal/handler"
	"example.com/redirect-service/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter(service *service.Redirect) *gin.Engine {
	router := gin.Default()

	// register GET /{code}
	handler.RegisterRedirect(router, service)

	return router
}
