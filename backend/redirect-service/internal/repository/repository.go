package repository

import "errors"

var notFoundErr = errors.New("not found")

type Repository interface {
	FindByCode(code string) (string, error)
}

func NewRepo() Repository {
    return newPostgreRepo()
}