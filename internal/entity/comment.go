package entity

const (
	upvote   = -1
	downvote = 1
)

type Comment struct {
	ID      string `json:"id"`
	Body    string `json:"body"`
	User    *User  `json:"author"`
	Created string `json:"created"`
}

type Vote struct {
	UserId string `json:"user"`
	Vote   int    `json:"vote"`
}
