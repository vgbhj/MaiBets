package models

import (
	"time"
)

type User struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Username  string `json:"username" gorm:"unique"`
	Password  string `json:"password"`
	CreatedAt time.Time
}
