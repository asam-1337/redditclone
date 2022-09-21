package repository

import (
	"fmt"
	"github.com/asam-1337/reddit-clone.git/internal/entity"
	"sync"
)

type PostsRepository struct {
	posts map[string]*entity.Post
	mu    *sync.Mutex
}

func NewPostsRepository(mu *sync.Mutex) *PostsRepository {
	return &PostsRepository{
		mu:    mu,
		posts: make(map[string]*entity.Post, 1),
	}
}

func (repo *PostsRepository) AddPost(post *entity.Post) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	repo.posts[post.ID] = post
}

func (repo *PostsRepository) GetPostByID(postID string) (*entity.Post, error) {

	repo.mu.Lock()
	post, ok := repo.posts[postID]
	repo.mu.Unlock()

	if !ok {
		return nil, fmt.Errorf("post %s does not exist", postID)
	}

	return post, nil
}

func (repo *PostsRepository) GetAll() ([]*entity.Post, error) {
	posts := make([]*entity.Post, 0, len(repo.posts))

	repo.mu.Lock()
	defer repo.mu.Unlock()

	for _, val := range repo.posts {
		posts = append(posts, val)
	}

	return posts, nil
}

func (repo *PostsRepository) GetPostsByUsername(username string) ([]*entity.Post, error) {
	posts := make([]*entity.Post, 0, 1)

	repo.mu.Lock()
	defer repo.mu.Unlock()

	for _, val := range repo.posts {
		if val.User.Username == username {
			posts = append(posts, val)
		}
	}

	return posts, nil
}

func (repo *PostsRepository) GetPostsByCategory(category string) ([]*entity.Post, error) {
	posts := make([]*entity.Post, 0, 1)

	repo.mu.Lock()
	defer repo.mu.Unlock()

	for _, val := range repo.posts {
		if val.Category == category {
			posts = append(posts, val)
		}
	}

	return posts, nil
}

func (repo *PostsRepository) DeletePost(postID string) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	delete(repo.posts, postID)
	return nil
}
