package service

import "redirect-service/internal/repository"

type Redirect struct {
	repo repository.Repository
}

func NewRedirect(repo repository.Repository) *Redirect {
	return &Redirect{repo: repo}
}

func (r *Redirect) FindOriginalURL(code string) (string, error) {
	return r.repo.FindByCode(code)
}
