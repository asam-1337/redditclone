package service

import (
	"crypto/md5"
	"fmt"
	"github.com/asam-1337/reddit-clone.git/internal/repository"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

const (
	salt            = ""
	tokenSigningKey = "asdwgfdgkert8"
)

type jwtUser struct {
	Username string `json:"username"`
	ID       int    `json:"id,string"`
}

type tokenClaims struct {
	jwt.StandardClaims
	User jwtUser `json:"user"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(username, password string) (string, error) {
	id, err := s.repo.CreateUser(username, password)
	if err != nil {
		return "", err
	}

	return s.GenerateToken(id, username)
}

func (s *AuthService) Authenticate(username, password string) (string, error) {
	user, err := s.repo.GetUserByUsernamePassword(username, password)
	if err != nil {
		return "", serviceError{"user not found"}
	}

	return s.GenerateToken(user.ID, user.Username)
}

func (s *AuthService) GenerateToken(userID int, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
		},
		jwtUser{
			Username: username,
			ID:       userID,
		},
	})

	return token.SignedString([]byte(tokenSigningKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		method, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok && method.Alg() != "HS256" {
			return nil, &serviceError{"bad sign method"}
		}

		return []byte(tokenSigningKey), nil
	}

	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, keyFunc)
	if err != nil {
		log.Println("service: ParseToken: ", err.Error())
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, serviceError{"token claims are not type of *tokenClaims"}
	}

	log.Println("service: ParseToken: Parsed token")

	return claims.User.ID, nil

}

func generateMd5Hash(data string) string {
	hash := md5.New()
	hash.Write([]byte(data))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
