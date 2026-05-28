package link

import "example.com/shared/postgres"

type LinksQuery struct {
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
	Order    string `form:"order"`
}

type LinksResponse struct {
	Data       []postgres.Link `json:"data"`
	Total      int64           `json:"total"`
	Page       int             `json:"page"`
	PageSize   int             `json:"page_size"`
	TotalPages int             `json:"total_pages"`
}
