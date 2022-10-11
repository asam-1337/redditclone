package entity

type User struct {
	ID       int    `json:"id,string" db:"id"`
	Username string `json:"username" db:"username"`
	Admin    bool   `json:"admin,omitempty" db:"admin"`
	Password string `json:"-" db:"password_hash"`
}
