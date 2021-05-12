package models

type User struct {
	ID           int    `gorm:"index"`
	Username     string `gorm:"index:,unique"`
	Email        string `gorm:"index:,unique"`
	PasswordHash string
	Permissions  uint
}
