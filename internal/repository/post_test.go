package repository

import (
	"fmt"
	"github.com/asam-1337/reddit-clone.git/internal/entity"
	_ "github.com/lib/pq"
	"testing"
)

func TestPostsRepository_GetPostsByCategory(t *testing.T) {
	db, err := NewPostgresDB(Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "docker",
		Password: "docker",
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		t.Fatalf("error occured in db open: %s", err.Error())
	}

	var posts []*entity.Post

	query := fmt.Sprintf("SELECT p.id, p.post_type, p.title, p.post_category, p.post_text, p.url, p.created, u.id, u.username FROM %s p INNER JOIN %s u on p.user_id = u.id WHERE u.username=$1", postsTable, usersTable)
	rows, err := db.Query(query, "asam12")
	if err != nil {
		t.Errorf(err.Error())
	}

	for rows.Next() {
		post := &entity.Post{}
		user := &entity.User{}
		fmt.Println(1)
		if err := rows.Scan(&post.ID, &post.Type, &post.Title, &post.Category, &post.Text, &post.URL, &post.Created, &user.ID, &user.Username); err != nil {
			t.Errorf(err.Error())
		}

		posts = append(posts, post)
	}
	rows.Close()
	fmt.Println(posts[0].ID)
}
