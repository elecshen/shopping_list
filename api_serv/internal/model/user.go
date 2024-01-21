package model

type User struct {
	Id           int    `json:"-" db:"ID"`
	Username     string `json:"username" binding:"required"`
	Salt         []byte `json:"-" db:"Salt"`
	PasswordHash string `json:"password" binding:"required" db:"Hash"`
}
