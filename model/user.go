package model

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

func (u *User) TableName() string {
	return "user"
}
