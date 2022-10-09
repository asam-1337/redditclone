package entity

type User struct {
	ID       int    `json:"id,string"`
	Username string `json:"username"`
	Admin    bool   `json:"admin,omitempty"`
	Password string `json:"-"`
}
