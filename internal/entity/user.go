package entity

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Admin    bool   `json:"admin,omitempty"`
	Password string `json:"-"`
}
