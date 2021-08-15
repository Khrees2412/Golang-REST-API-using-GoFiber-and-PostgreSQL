package models

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
}

type Book struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Author string `json:"author"`
	UserID int    `json:"user_id"`
}
