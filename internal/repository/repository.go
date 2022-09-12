package repository

import (
	"github.com/asam-1337/reddit-clone.git/internal/entity"
	"sync"
)

type Authorization interface {
	CreateUser(user *entity.User) error
}

type Posts interface {
	GetAll()
	GetByID()
	AddPost()
	AddComment()
}

type Repository struct {
	Authorization
	Posts
}

func NewRepository(mu *sync.Mutex) *Repository {
	return &Repository{
		Authorization: NewUserRepository(mu),
	}
}
