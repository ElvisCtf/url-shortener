package shorten

import (
	"errors"
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

	response, err := c.service.Create(request.OriginalURL)
	if errors.Is(err, ErrSSRF) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "URL is not allowed"})
		return
	}
	if err != nil {
		slog.Error("ShortenController Shorten: Failed to create shortened URL", "original_url", request.OriginalURL, "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	ctx.JSON(http.StatusOK, response)
}
