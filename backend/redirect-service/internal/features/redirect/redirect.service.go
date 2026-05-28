package redirect

type RedirectService struct {
	repo *RedirectRepository
}

func NewRedirectService(repo *RedirectRepository) *RedirectService {
	return &RedirectService{repo: repo}
}

func (r *RedirectService) FindOriginalURL(code string) (string, error) {
	return r.repo.FindByCode(code)
}
