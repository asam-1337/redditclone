package repository

import (
	"fmt"
	"github.com/asam-1337/reddit-clone.git/internal/entity"
	"sync"
)

type SessionsRepository struct {
	data map[string]*entity.Session
	mu   *sync.Mutex
}

func NewSessionsRepository(mu *sync.Mutex) *SessionsRepository {
	return &SessionsRepository{
		data: make(map[string]*entity.Session, 1),
		mu:   mu,
	}
}

func (repo *SessionsRepository) CreateSession(user *entity.User) error {
	user.ID = ""
	return nil
}

func (repo *SessionsRepository) GetSession(userID string) (*entity.Session, error) {
	session, ok := repo.data[userID]
	if !ok {
		return nil, fmt.Errorf(`user is not authorized`)
	}
	return session, nil
}
