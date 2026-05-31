package redirect

import (
	"context"
	"log/slog"
	"time"
)

type RedirectService struct {
	repo *RedirectRepository
}

func NewRedirectService(repo *RedirectRepository) *RedirectService {
	return &RedirectService{repo: repo}
}

func (s *RedirectService) FindOriginalURL(code string) (string, error) {
	url, err := s.repo.FindByCode(code)
	if err != nil {
		return "", err
	}

	go func(c string) {
		s.LogClicks(c)
	}(code)

	return url, err
}

func (s *RedirectService) LogClicks(code string) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := s.repo.AddClicks(ctx, code); err != nil {
		slog.Error("RedirectService LogClicks: failed to increment click for", "code", code, "error", err)
	}
}
