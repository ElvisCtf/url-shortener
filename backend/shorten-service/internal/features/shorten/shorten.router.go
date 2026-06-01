package shorten

import (
	shared "example.com/shared"

	"github.com/gin-gonic/gin"
)

type ShortenRouter struct {
	controller *ShortenController
}

func NewShortenRouter(c *ShortenController) *ShortenRouter {
	return &ShortenRouter{controller: c}
}

func (r *ShortenRouter) SetupRoutes(rg *gin.RouterGroup, config *shared.Config) {
	group := rg.Group("api/v1")

	group.POST("/shorten", RateLimitByIP(), r.controller.Shorten)
}
