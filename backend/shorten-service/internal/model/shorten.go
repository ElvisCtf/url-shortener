package model

type ShortenRequest struct {
    OriginalURL string `json:"original_url" binding:"required,url"`
}

type ShortenResponse struct {
    OriginalURL string `json:"original_url"`
    ShortenURL  string `json:"shorten_url"`
}
