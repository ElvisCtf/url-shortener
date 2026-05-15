package router

import (
	"shorten-service/internal/handler"
	"shorten-service/internal/service"

    "github.com/gin-gonic/gin"
)

func SetupRouter(service *service.Shorten) *gin.Engine {
	router := gin.Default()

	// register POST /shorten
	handler.RegisterShorten(router, service)

	return router
}