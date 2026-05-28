package redirect

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/shared"
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
		slog.Error("RedirectController Redirect: failed to find original URL", "code", code, "error", err)
		ctx.Status(http.StatusNotFound)
	}
}
