package entity

const (
	upvote   = -1
	downvote = 1
)

type Vote struct {
	User string `json:"user"`
	Vote int    `json:"vote"`
}
