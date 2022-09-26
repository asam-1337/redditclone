package service

import (
	"github.com/asam-1337/reddit-clone.git/internal/entity"
	"github.com/asam-1337/reddit-clone.git/internal/repository"
)

type serviceError struct {
	Err string `json:"service_error"`
}

func (e serviceError) Error() string {
	return e.Err
}

type Authorization interface {
	CreateUser(username, password string) (string, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (string, string, error)
}

type Posts interface {
	CreatePost(post *entity.Post, userID string) (*entity.Post, error)
	GetPostByID(postID string) (*entity.Post, error)
	GetPostsByUsername(username string) ([]*entity.Post, error)
	GetPostsByCategory(category string) ([]*entity.Post, error)
	GetAll() ([]*entity.Post, error)
	DeletePost(postID string) error

	CreateComment(userID, postID, comment string) (*entity.Post, error)

	Vote(postID string, vote *entity.Vote) (*entity.Post, error)
	Unvote(userID, postID string) (*entity.Post, error)
}

type Service struct {
	Authorization
	Posts
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Posts:         NewPostsService(repo.Authorization, repo.Posts),
	}
}
