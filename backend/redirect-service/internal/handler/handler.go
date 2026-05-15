package handler

import (
	"net/http"
	"redirect-service/internal/service"

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