package redirect

import (
	"example.com/shared"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RedirectController struct {
	service *RedirectService
	config  *shared.Config
}

func NewRedirectController(s *RedirectService, config *shared.Config) *RedirectController {
	return &RedirectController{service: s, config: config}
}

func (c *RedirectController) Redirect(ctx *gin.Context) {
	code := ctx.Param("code")

	originalURL, err := c.service.FindOriginalURL(code)

	if err == nil && originalURL != "" {
		ctx.Redirect(http.StatusFound, originalURL)
	} else {
		ctx.Status(http.StatusNotFound)
	}
}
