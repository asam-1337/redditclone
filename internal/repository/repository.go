package repository

import (
	"github.com/asam-1337/reddit-clone.git/internal/entity"
	"github.com/jmoiron/sqlx"
)

type repoError struct {
	Err string `json:"repo_error"`
}

func (e repoError) Error() string {
	return e.Err
}

type Authorization interface {
	CreateUser(username, password string) (int, error)
	GetUserByUsername(username string) (*entity.User, error)
	GetUserByID(userID int) (*entity.User, error)
}

type Posts interface {
	CreatePost(post *entity.Post) (int, error)
	GetPostByID(postID int) (*entity.Post, error)
	GetPostsByUsername(username string) ([]*entity.Post, error)
	GetPostsByCategory(category string) ([]*entity.Post, error)
	GetAll() ([]*entity.Post, error)
	DeletePost(postID int) error

	GetVotes(postID int) ([]*entity.Vote, error)
	Vote(userID int, postID int, vote int) error
	Unvote(userID int, postID int) error
}

type Repository struct {
	Authorization
	Posts
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewUserRepository(db),
		Posts:         NewPostsRepository(db),
	}
}
