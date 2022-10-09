package repository

import (
	"fmt"
	"github.com/asam-1337/reddit-clone.git/internal/entity"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) CreateUser(username, password string) (int, error) {
	var id int

	insertQuery := fmt.Sprintf("INSERT INTO %s (username, password_hash) VALUES ($1, $2) RETURNING id", usersTable)
	row := repo.db.QueryRow(insertQuery, username, password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (repo *UserRepository) GetUserByUsernamePassword(username string, password string) (*entity.User, error) {
	user := &entity.User{
		Username: username,
		Password: password,
	}

	selectQuery := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	row := repo.db.QueryRow(selectQuery, username, password)

	if err := row.Scan(&user.ID); err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserRepository) GetUserByID(userID int) (*entity.User, error) {
	user := &entity.User{
		ID: userID,
	}

	query := fmt.Sprintf("SELECT username FROM %s WHERE id=$1", usersTable)
	err := repo.db.Get(user, query, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
