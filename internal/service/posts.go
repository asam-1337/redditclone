package service

import "github.com/asam-1337/reddit-clone.git/internal/repository"

type PostsService struct {
	repo repository.Posts
}

func NewPostsService(repo repository.Posts) *PostsService {
	return &PostsService{repo: repo}
}
