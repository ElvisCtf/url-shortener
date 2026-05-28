package link

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/shared"
)

type LinkController struct {
	service *LinkService
	config  *shared.Config
}

func NewLinkController(s *LinkService, config *shared.Config) *LinkController {
	return &LinkController{service: s, config: config}
}

func (c *LinkController) GetLinks(ctx *gin.Context) {
	var query LinksQuery

	// Bind query parameters (?page=1&page_size=10&order=id)
	if err := ctx.ShouldBindQuery(&query); err != nil {
		slog.Error("LinkController GetLinks: Failed to bind query parameters", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := c.service.GetLinks(query)
	if err != nil {
		slog.Error("LinkController GetLinks: Failed to fetch links", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch links"})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *LinkController) BulkUpdateActive(ctx *gin.Context) {
	var req BulkActiveUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		slog.Error("LinkController BulkUpdateActive: Failed to bind JSON", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.BulkUpdateActive(req.IDs, *req.Active); err != nil {
		slog.Error("LinkController BulkUpdateActive: Failed to update links", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update links"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true})
}
