package repository

import (
	"github.com/asam-1337/reddit-clone.git/internal/entity"
	"sync"
)

type UserRepository struct {
	users map[string]*entity.User
	mu    *sync.Mutex
}

func NewUserRepository(mu *sync.Mutex) *UserRepository {
	return &UserRepository{
		users: make(map[string]*entity.User, 1),
		mu:    mu,
	}
}

func (repo *UserRepository) AddUser(user *entity.User) error {
	_, err := repo.GetUserByID(user.ID)
	if err == nil {
		return &repoError{Err: "user already exist"}
	}

	repo.mu.Lock()
	repo.users[user.ID] = user
	repo.mu.Unlock()

	return nil
}

func (repo *UserRepository) GetUserByID(userID string) (*entity.User, error) {
	repo.mu.Lock()
	user, ok := repo.users[userID]
	repo.mu.Unlock()

	if !ok {
		return nil, &repoError{Err: "user does not exist"}
	}
	return user, nil
}
