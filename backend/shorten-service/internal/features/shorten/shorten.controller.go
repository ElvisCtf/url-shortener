package shorten

import (
	"log/slog"
	"net/http"

	"example.com/shared"

	"github.com/gin-gonic/gin"
)

type ShortenController struct {
	service *ShortenService
	config  *shared.Config
}

func NewShortenController(s *ShortenService, config *shared.Config) *ShortenController {
	return &ShortenController{service: s, config: config}
}

func (c *ShortenController) Shorten(ctx *gin.Context) {
	var request ShortenRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		slog.Error("ShortenController Shorten: Failed to bind JSON", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := c.service.Create(request.OriginalURL)
	if response != nil {
		ctx.JSON(http.StatusOK, response)
	} else {
		slog.Error("ShortenController Shorten:Failed to create shortened URL", "original_url", request.OriginalURL)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
}
