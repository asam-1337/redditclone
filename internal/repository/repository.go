package repository

import (
	"github.com/asam-1337/reddit-clone.git/internal/entity"
	"sync"
)

type repoError struct {
	Err string `json:"repo_error"`
}

func (e repoError) Error() string {
	return e.Err
}

type Authorization interface {
	AddUser(user *entity.User) error
	GetUserByID(userID string) (*entity.User, error)
}

type Posts interface {
	AddPost(post *entity.Post)
	GetPostByID(postID string) (*entity.Post, error)
	GetPostsByUsername(username string) ([]*entity.Post, error)
	GetPostsByCategory(category string) ([]*entity.Post, error)
	GetAll() ([]*entity.Post, error)
	DeletePost(postID string) error

	Vote(postID string, vote *entity.Vote) (*entity.Post, error)
	Unvote(userID, postID string) (*entity.Post, error)
}

type Repository struct {
	Authorization
	Posts
}

func NewRepository(mu *sync.Mutex) *Repository {
	return &Repository{
		Authorization: NewUserRepository(mu),
		Posts:         NewPostsRepository(mu),
	}
}
