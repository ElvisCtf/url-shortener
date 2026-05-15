package model

import "time"

type Link struct {
    ID          uint64
    Code        string
    OriginalURL string
    CreatedAt   time.Time
}