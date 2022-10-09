package repository

import (
	"fmt"
	"github.com/asam-1337/reddit-clone.git/internal/entity"
	"github.com/jmoiron/sqlx"
)

type PostsRepository struct {
	db *sqlx.DB
}

func NewPostsRepository(db *sqlx.DB) *PostsRepository {
	return &PostsRepository{
		db: db,
	}
}

func (repo *PostsRepository) CreatePost(post *entity.Post) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (post_type, title, post_category, post_text, url, user_id, created) VALUES ($1, $2, $3, $4, $5, $6, $7, $8,) RETURNING id", usersTable)
	row := repo.db.QueryRow(query, post.Type, post.Title, post.Category, post.Text, post.URL, post.User.ID, post.Created)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (repo *PostsRepository) GetPostByID(postID int) (*entity.Post, error) {

}

func (repo *PostsRepository) GetAll() ([]*entity.Post, error) {

}

func (repo *PostsRepository) GetPostsByUsername(username string) ([]*entity.Post, error) {
	var posts []*entity.Post

	query := fmt.Sprintf("SELECT p.post_type, p.title, p.post_category, p.post_text, p.url, p.username, p.user_id, p.created FROM %s p INNER JOIN %s u on p.user_id = u.id WHERE p.username=$1", postsTable, usersTable)
	repo.db.Select(&posts, query, username)
}

func (repo *PostsRepository) GetPostsByCategory(category string) ([]*entity.Post, error) {

}

func (repo *PostsRepository) DeletePost(postID int) error {

}

func (repo *PostsRepository) Vote(postID int, vote *entity.Vote) (*entity.Post, error) {

}

func (repo *PostsRepository) Unvote(userID int, postID int) (*entity.Post, error) {

}
