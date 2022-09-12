package repository

import (
	"github.com/asam-1337/reddit-clone.git/internal/entity"
	"math/rand"
	"sync"
)

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	idLength    = 32
)

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

type UserRepository struct {
	users []*entity.User
	mu    *sync.Mutex
}

func NewUserRepository(mu *sync.Mutex) *UserRepository {
	return &UserRepository{
		users: make([]*entity.User, 1),
		mu:    mu,
	}
}

func (repo *UserRepository) CreateUser(user *entity.User) error {
	user.ID = RandStringRunes(idLength)

	repo.mu.Lock()
	repo.users = append(repo.users, user)
	repo.mu.Unlock()
	return nil
}
