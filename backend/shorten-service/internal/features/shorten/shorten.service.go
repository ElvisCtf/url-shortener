package shorten

import (
	"example.com/shared"
)

type ShortenService struct {
	repo   *ShortenRepository
	config *shared.Config
}

func NewShortenService(repo *ShortenRepository, config *shared.Config) *ShortenService {
	return &ShortenService{repo: repo, config: config}
}

func (s *ShortenService) Create(originalURL string) *ShortenResponse {
	code, err := s.repo.Save(originalURL)
	if err == nil {
		return &ShortenResponse{
			OriginalURL: originalURL,
			ShortenURL:  s.config.Redirect + "/" + code,
		}
	} else {
		return nil
	}
}
