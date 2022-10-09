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

func (s *PostsService) CreatePost(post *entity.Post, userID int) (*entity.Post, error) {
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	post.User = user
	post.Created = time.Now().Format(time.RFC3339)
	vote := &entity.Vote{
		UserId: userID,
		Vote:   1,
	}
	post.Votes = append(post.Votes, vote)
	post.Views = 1
	post.Score = 1
	post.UpvotePercentage = 100

	post.ID, err = s.postRepo.CreatePost(post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (s *PostsService) GetPostByID(postID int) (*entity.Post, error) {
	post, err := s.postRepo.GetPostByID(postID)
	if err != nil {
		return nil, err
	}
	post.Views++

	return post, nil
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

func (s *PostsService) DeletePost(postID int) error {
	return s.postRepo.DeletePost(postID)
}

func (s *PostsService) CreateComment(userID int, postID int, body string) (*entity.Post, error) {
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	post, err := s.postRepo.GetPostByID(postID)
	if err != nil {
		return nil, err
	}

	comment := &entity.Comment{
		Body:    body,
		Created: time.Now().Format(time.RFC3339),
		User:    user,
	}

	post.Comments = append(post.Comments, comment)

	return post, nil
}

func updatePostVotes(post *entity.Post) {
	var upVotes, postVote int
	post.Score = 0

	for _, val := range post.Votes {
		postVote = val.Vote
		if postVote == 1 {
			upVotes += postVote
		}
		post.Score += val.Vote
	}

	if len(post.Votes) == 0 {
		post.UpvotePercentage = 0
		return
	}

	post.UpvotePercentage = upVotes / len(post.Votes) * 100
}

func (s *PostsService) Vote(postID int, vote *entity.Vote) (*entity.Post, error) {
	post, err := s.postRepo.Vote(postID, vote)
	if err != nil {
		return nil, err
	}

	updatePostVotes(post)

	return post, nil
}

func (s *PostsService) Unvote(userID int, postID int) (*entity.Post, error) {
	post, err := s.postRepo.Unvote(userID, postID)
	if err != nil {
		return nil, err
	}

	updatePostVotes(post)

	return post, nil
}

func getRandomSeed(n int) string {
	runes := make([]rune, n)

	for i := range runes {
		runes[i] = alphabet[rand.Intn(len(alphabet))]
	}

	return string(runes)
}
