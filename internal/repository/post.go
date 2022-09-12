package repository

import (
	"github.com/asam-1337/reddit-clone.git/internal/entity"
	"sync"
)

type PostsMemRepository struct {
	posts []entity.Post
	mu    *sync.Mutex
}
