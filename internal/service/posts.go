package service

import (
	"github.com/asam-1337/reddit-clone.git/internal/entity"
	"github.com/asam-1337/reddit-clone.git/internal/repository"
	"math/rand"
	"time"
)

const idLen = 32

var alphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type PostsService struct {
	userRepo repository.Authorization
	postRepo repository.Posts
}

func NewPostsService(userRepo repository.Authorization, postRepo repository.Posts) *PostsService {
	return &PostsService{
		userRepo: userRepo,
		postRepo: postRepo,
	}
}

func (s *PostsService) CreatePost(post *entity.Post, userID string) (*entity.Post, error) {
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	post.User = user
	post.ID = getRandomSeed(idLen)
	post.Created = time.Now().Format(time.RFC3339)

	s.postRepo.AddPost(post)

	return post, nil
}

func (s *PostsService) GetPostByID(postID string) (*entity.Post, error) {
	return s.postRepo.GetPostByID(postID)
}

func (s *PostsService) GetAll() ([]*entity.Post, error) {
	return s.postRepo.GetAll()
}

func (s *PostsService) GetPostsByUsername(username string) ([]*entity.Post, error) {
	return s.postRepo.GetPostsByUsername(username)
}

func (s *PostsService) GetPostsByCategory(category string) ([]*entity.Post, error) {
	return s.postRepo.GetPostsByCategory(category)
}

func (s *PostsService) DeletePost(postID string) error {
	return s.postRepo.DeletePost(postID)
}

func (s *PostsService) CreateComment(userID, postID, body string) (*entity.Post, error) {
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	post, err := s.postRepo.GetPostByID(postID)
	if err != nil {
		return nil, err
	}

	comment := &entity.Comment{
		ID:      getRandomSeed(24),
		Body:    body,
		Created: time.Now().Format(time.RFC3339),
		User:    user,
	}

	post.Comments = append(post.Comments, comment)

	return post, nil
}

func getRandomSeed(n int) string {
	runes := make([]rune, n)

	for i := range runes {
		runes[i] = alphabet[rand.Intn(len(alphabet))]
	}

	return string(runes)
}
