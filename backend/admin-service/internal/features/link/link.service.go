package link

import (
	"math"

	"example.com/shared"
)

type LinkService struct {
	repo   *LinkRepository
	config *shared.Config
}

func NewLinkService(repo *LinkRepository, config *shared.Config) *LinkService {
	return &LinkService{repo: repo, config: config}
}

func (s *LinkService) GetLinks(query LinksQuery) (*LinksResponse, error) {
	if query.Page <= 0 {
		query.Page = 1
	}
	if query.PageSize <= 0 {
		query.PageSize = 10
	}

	links, total, err := s.repo.GetPaginated(query.Page, query.PageSize, query.Order)
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(query.PageSize)))

	return &LinksResponse{
		Data:       links,
		Total:      total,
		Page:       query.Page,
		PageSize:   query.PageSize,
		TotalPages: totalPages,
	}, nil
}

func (s *LinkService) BulkUpdateActive(ids []uint, active bool) error {
	return s.repo.BulkUpdateActive(ids, active)
}
