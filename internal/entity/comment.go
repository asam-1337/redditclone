package entity

import "time"

type Comment struct {
	ID      string        `json:"id"`
	Body    string        `json:"body"`
	Author  *User         `json:"author"`
	Created time.Duration `json:"created"`
}
