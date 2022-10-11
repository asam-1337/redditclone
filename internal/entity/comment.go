package entity

const (
	upvote   = -1
	downvote = 1
)

type Comment struct {
	ID      int    `json:"id,string"`
	Body    string `json:"body"`
	User    *User  `json:"author"`
	Created string `json:"created"`
}

type Vote struct {
	UserId int `json:"user,string" db:"user_id"`
	Vote   int `json:"vote" db:"vote"`
}
