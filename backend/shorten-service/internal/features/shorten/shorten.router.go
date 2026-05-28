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

	// /auth/admins requires JWT (only root admin can register a new admin)
	group.POST("/shorten", r.controller.Shorten)
}
