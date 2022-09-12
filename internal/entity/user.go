package entity

type User struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"username"`
	Admin    bool   `json:"admin,omitempty"`
	Password string `json:"password,omitempty"`
}
