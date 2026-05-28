package link

import (
	"example.com/admin-service/internal/features/auth"
	"example.com/shared"

	"github.com/gin-gonic/gin"
)

type LinkRouter struct {
	controller *LinkController
}

func NewLinkRouter(c *LinkController) *LinkRouter {
	return &LinkRouter{controller: c}
}

func (r *LinkRouter) SetupRoutes(rg *gin.RouterGroup, config *shared.Config) {
	linkGroup := rg.Group("/api/v1")

	linkGroup.GET("/links", auth.RequireAccessToken(config, true), r.controller.GetLinks)
}
