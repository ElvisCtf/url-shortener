package link

import (
	"example.com/shared"
	"github.com/gin-gonic/gin"
	"net/http"
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := c.service.GetLinks(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch links"})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
