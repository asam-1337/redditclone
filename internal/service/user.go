package service

import (
	"crypto/md5"
	"fmt"
	"github.com/asam-1337/reddit-clone.git/internal/entity"
	"github.com/asam-1337/reddit-clone.git/internal/repository"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

const (
	salt            = "vcmn324mdsf92mfksdsdf3"
	tokenSigningKey = "asdwgfdgkert8"
)

type jwtUser struct {
	Username string `json:"username"`
	ID       string `json:"id"`
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

func (s *AuthService) CreateUser(user *entity.User) (string, error) {
	username := user.Username
	password := user.Password

	user.ID = generateMd5Hash(username)
	user.Password = generateMd5Hash(user.Password)

	err := s.repo.AddUser(user)
	if err != nil {
		return "", err
	}

	return s.GenerateToken(username, password)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	userID := generateMd5Hash(username)

	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return "", serviceError{"invalid login"}
	}

	if user.Password != generateMd5Hash(password) {
		return "", serviceError{"invalid password"}
	}

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

func (s *AuthService) ParseToken(accessToken string) (string, string, error) {
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
		return "", "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", "", serviceError{"token claims are not type of *tokenClaims"}
	}

	log.Println("service: ParseToken: Parsed token")

	return claims.User.ID, claims.User.Username, nil

}

func generateMd5Hash(data string) string {
	hash := md5.New()
	hash.Write([]byte(data))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
