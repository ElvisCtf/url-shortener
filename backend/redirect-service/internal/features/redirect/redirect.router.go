package redirect

import (
	"github.com/gin-gonic/gin"

	"example.com/shared"
)

type RedirectRouter struct {
	controller *RedirectController
}

func NewRedirectRouter(c *RedirectController) *RedirectRouter {
	return &RedirectRouter{controller: c}
}

func (r *RedirectRouter) SetupRoutes(rg *gin.RouterGroup, config *shared.Config) {
	group := rg.Group("")

	group.GET("/:code", r.controller.Redirect)
}
