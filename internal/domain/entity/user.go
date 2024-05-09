package entity

import (
	"time"
)

type User struct {
	ID           int64     `gorm:"column:id;primary_key"`
	Username     string    `gorm:"column:username"`
	HashPassword string    `gorm:"column:hash_password"`
	Avatar       string    `gorm:"column:avatar"`
	CreatedAt    time.Time `gorm:"column:created_at"`
}

type NewUser struct {
	Username string `gorm:"column:username"`
	Password string `gorm:"column:hash_password"`
}

type UserOutput struct {
	Id        int64
	Username  string
	Avatar    string
	CreatedAt time.Time
}
