package repository

import (
	"fmt"
	"github.com/asam-1337/reddit-clone.git/internal/entity"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type postDTO struct {
	ID               int    `db:"id"`
	Type             string `db:"post_type"`
	Title            string `db:"title"`
	Category         string `db:"post_category"`
	Text             string `db:"post_text"`
	URL              string `db:"url"`
	Score            int    `db:"score"`
	Views            int    `db:"views"`
	UpvotePercentage int    `db:"upvote_percentage"`
	UserID           int    `db:"user_id"`
	Created          string `db:"created"`
}

type commentDTO struct {
	ID      int    `db:"id"`
	Body    string `db:"body"`
	UserID  int    `db:"user_id"`
	Created string `db:"created"`
}

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

	query := fmt.Sprintf("INSERT INTO %s (post_type, title, post_category, post_text, url, user_id, created) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", postsTable)
	row := repo.db.QueryRow(query, post.Type, post.Title, post.Category, post.Text, post.URL, post.User.ID, post.Created)

	if err := row.Scan(&id); err != nil {
		logrus.Info("create error", err.Error())
		return 0, err
	}

	return id, nil
}

func (repo *PostsRepository) GetPostByID(postID int) (*entity.Post, error) {
	post := &entity.Post{
		ID:   postID,
		User: &entity.User{},
	}

	var err error
	query := fmt.Sprintf("SELECT p.post_type, p.title, p.post_category, p.post_text, p.url, p.created, u.id, u.username FROM %s p INNER JOIN %s u ON p.user_id=u.id WHERE p.id=$1", postsTable, usersTable)
	row := repo.db.QueryRow(query, postID)

	if err := row.Scan(&post.Type, &post.Title, &post.Category, &post.Text, &post.URL, &post.Created, &post.User.ID, &post.User.Username); err != nil {
		return nil, err
	}

	post.Comments, err = repo.GetComments(postID)
	if err != nil {
		return nil, err
	}

	post.Votes, err = repo.GetVotes(post.ID)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (repo *PostsRepository) GetAll() ([]*entity.Post, error) {
	var posts []*entity.Post

	query := fmt.Sprintf("SELECT p.id, p.post_type, p.title, p.post_category, p.post_text, p.url, p.created, u.id, u.username FROM %s p INNER JOIN %s u on p.user_id = u.id", postsTable, usersTable)
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		post := &entity.Post{
			User: &entity.User{},
		}

		if err := rows.Scan(&post.ID, &post.Type, &post.Title, &post.Category, &post.Text, &post.URL, &post.Created, &post.User.ID, &post.User.Username); err != nil {
			return nil, err
		}

		post.Comments, err = repo.GetComments(post.ID)
		if err != nil {
			return nil, err
		}

		post.Votes, err = repo.GetVotes(post.ID)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, err
}

func (repo *PostsRepository) GetPostsByUsername(username string) ([]*entity.Post, error) {
	var posts []*entity.Post

	query := fmt.Sprintf("SELECT p.id, p.post_type, p.title, p.post_category, p.post_text, p.url, p.created, u.id, u.username FROM %s p INNER JOIN %s u on p.user_id = u.id WHERE u.username=$1", postsTable, usersTable)
	rows, err := repo.db.Query(query, username)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		post := &entity.Post{
			User: &entity.User{},
		}

		if err := rows.Scan(&post.ID, &post.Type, &post.Title, &post.Category, &post.Text, &post.URL, &post.Created, &post.User.ID, &post.User.Username); err != nil {
			return nil, err
		}

		post.Comments, err = repo.GetComments(post.ID)
		if err != nil {
			return nil, err
		}

		post.Votes, err = repo.GetVotes(post.ID)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (repo *PostsRepository) GetPostsByCategory(category string) ([]*entity.Post, error) {
	var posts []*entity.Post

	query := fmt.Sprintf("SELECT p.id, p.post_type, p.title, p.post_category, p.post_text, p.url, p.created, u.id, u.username FROM %s p INNER JOIN %s u on p.user_id = u.id WHERE p.post_category=$1", postsTable, usersTable)
	rows, err := repo.db.Query(query, category)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		post := &entity.Post{
			User: &entity.User{},
		}

		if err := rows.Scan(&post.ID, &post.Type, &post.Title, &post.Category, &post.Text, &post.URL, &post.Created, &post.User.ID, &post.User.Username); err != nil {
			return nil, err
		}

		post.Comments, err = repo.GetComments(post.ID)
		if err != nil {
			return nil, err
		}

		post.Votes, err = repo.GetVotes(post.ID)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (repo *PostsRepository) DeletePost(postID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", postsTable)
	_, err := repo.db.Exec(query, postID)
	if err != nil {
		return err
	}

	return nil
}

func (repo *PostsRepository) GetComments(postID int) ([]*entity.Comment, error) {
	var comments []*entity.Comment

	query := fmt.Sprintf("SELECT ct.id, ct.body, ct.created, u.id, u.username FROM %s ct INNER JOIN %s u ON ct.user_id=u.id WHERE ct.post_id=$1", commentsTable, usersTable)
	rows, err := repo.db.Query(query, postID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		comment := &entity.Comment{
			User: &entity.User{},
		}

		err = rows.Scan(&comment.ID, &comment.Body, &comment.Created, &comment.User.ID, &comment.User.Username)
		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}

func (repo *PostsRepository) GetVotes(postID int) ([]*entity.Vote, error) {
	var votes []*entity.Vote

	query := fmt.Sprintf("SELECT user_id, vote FROM %s WHERE post_id=$1", votesTable)
	err := repo.db.Select(&votes, query, postID)
	if err != nil {
		return nil, err
	}

	return votes, nil
}

func (repo *PostsRepository) Vote(userID int, postID int, vote int) error {
	query := fmt.Sprintf("INSERT INTO %s (user_id, post_id, vote) VALUES ($1, $2, $3)", votesTable)
	_, err := repo.db.Exec(query, userID, postID, vote)
	if err == nil {
		return nil
	}

	query = fmt.Sprintf("UPDATE %s SET vote=$1 WHERE user_id=$2 AND post_id=$3", votesTable)
	_, err = repo.db.Exec(query, vote, userID, postID)
	if err != nil {
		return err
	}

	return nil
}

func (repo *PostsRepository) Unvote(userID int, postID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id=$1 AND post_id=$2", votesTable)
	_, err := repo.db.Exec(query, userID, postID)
	if err != nil {
		return err
	}

	return nil
}
