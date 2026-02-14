package repository

import (
    "errors"
)

var saveErr = errors.New("Cannot save")
var notFoundErr = errors.New("not found")

type Repository interface {
    Save(originalURL string) (string, error)
}

func NewRepo(storage string) Repository {
    if storage == "postgres" {
        return NewPostgreRepo()
    } else {
        return NewMemoryRepo()
    }
}