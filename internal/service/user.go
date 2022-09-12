package service

import (
	"github.com/asam-1337/reddit-clone.git/internal/entity"
	"github.com/asam-1337/reddit-clone.git/internal/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user *entity.User) error {
	return s.repo.CreateUser(user)
}
