package models

type User struct {
	ID       int64  `json:"id" gorm:"autoIncrement"`
	Username string `json:"username"`
	Password string `json:"password"`
}
