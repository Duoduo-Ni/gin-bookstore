package model

// User 结构体
type User struct {
	ID       int    `json:"iD"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
