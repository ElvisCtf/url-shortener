package repository

import (
    "sync"
	"time"

    "shorten-service/internal/model"
    "shorten-service/internal/util"
)

type MemoryRepo struct {
    mutex           sync.RWMutex
    counter         uint64
    store           map[string]model.Link
    reverseStore    map[string]string
}

func NewMemoryRepo() *MemoryRepo {
    return &MemoryRepo{
        store: make(map[string]model.Link),
        reverseStore: make(map[string]string),
    }
}

func (repo *MemoryRepo) Save(originalURL string) (string, error) {
    repo.mutex.Lock()
    defer repo.mutex.Unlock()

    code, ok := repo.reverseStore[originalURL]
    if ok {
        return code, nil
    }

    repo.counter++
    code = util.EncodeBase62(repo.counter)
    link := model.Link{
        Code:        code,
        OriginalURL: originalURL,
        CreatedAt:   time.Now(),
    }
    repo.store[link.Code] = link
    repo.reverseStore[originalURL] = code

    return code, nil
}
