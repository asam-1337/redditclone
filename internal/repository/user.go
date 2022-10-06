package repository

import (
	"fmt"
	"github.com/asam-1337/reddit-clone.git/internal/entity"
	"github.com/jmoiron/sqlx"
	"log"
	"sync"
)

type UserRepository struct {
	db    *sqlx.DB
	users map[string]*entity.User
	mu    *sync.Mutex
}

func NewUserRepository(mu *sync.Mutex, db *sqlx.DB) *UserRepository {
	return &UserRepository{
		users: make(map[string]*entity.User, 1),
		mu:    mu,
		db:    db,
	}
}

func (repo *UserRepository) AddUser(user *entity.User) error {
	var id int

	selectQuery := fmt.Sprintf("SELECT id FROM %s WHERE username=$1", usersTable)
	row := repo.db.QueryRow(selectQuery, user.Username)

	if err := row.Scan(&id); err == nil {
		return &repoError{"user already exist"}
	}

	insertQuery := fmt.Sprintf("INSERT INTO %s (username, password_hash) values ($1, $2) RETURNING id", usersTable)
	row = repo.db.QueryRow(insertQuery, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		log.Printf("databse error: %s", err.Error())
		return err
	}

	fmt.Println(id)
	//////////////////////
	_, err := repo.GetUserByID(user.ID)
	if err == nil {
		return &repoError{Err: "user already exist"}
	}

	repo.mu.Lock()
	repo.users[user.ID] = user
	repo.mu.Unlock()

	return nil
}

func (repo *UserRepository) GetUserByID(userID string) (*entity.User, error) {
	//////////////
	repo.mu.Lock()
	user, ok := repo.users[userID]
	repo.mu.Unlock()

	if !ok {
		return nil, &repoError{Err: "user does not exist"}
	}
	return user, nil
}
