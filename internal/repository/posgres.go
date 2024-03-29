package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable    = "users"
	postsTable    = "posts"
	commentsTable = "comments"
	votesTable    = "votes"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	conf := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)

	db, err := sqlx.Open("postgres", conf)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(15)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
