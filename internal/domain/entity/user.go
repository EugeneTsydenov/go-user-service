package entity

import (
	"time"
)

type User struct {
	Id           int64
	Username     string
	HashPassword string
	Avatar       string
	CreatedAt    time.Time
}

type NewUser struct {
	Username string `gorm:"column:username"`
	Password string `gorm:"column:hash_password"`
}
