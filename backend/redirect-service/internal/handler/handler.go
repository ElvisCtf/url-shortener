package handler

import (
	"example.com/redirect-service/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRedirect(router *gin.Engine, service *service.Redirect) {
	router.GET("/:code", func(c *gin.Context) {
		code := c.Param("code")
		originalURL, error := service.FindOriginalURL(code)

		if error == nil && originalURL != "" {
			c.Redirect(http.StatusFound, originalURL)
		} else {
			c.Status(http.StatusNotFound)
		}
	})
}
