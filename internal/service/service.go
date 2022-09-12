package service

import (
	"github.com/asam-1337/reddit-clone.git/internal/entity"
	"github.com/asam-1337/reddit-clone.git/internal/repository"
)

type Authorization interface {
	CreateUser(user *entity.User) error
}

type Posts interface {
}

type Service struct {
	Authorization
	Posts
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Posts:         NewPostsService(repo.Posts),
	}
}
