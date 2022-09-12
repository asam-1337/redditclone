package entity

type Post struct {
	ID               string   `json:"id"`
	Type             string   `json:"type"`
	Title            string   `json:"title"`
	Category         string   `json:"category"`
	Score            int      `json:"score"`
	Views            int      `json:"views"`
	UpvotePercentage int      `json:"upvotePercentage"`
	User             *User    `json:"author"`
	Comments         *Comment `json:"comments,omitempty"`
	Votes            *[]Vote  `json:"votes"`
}
