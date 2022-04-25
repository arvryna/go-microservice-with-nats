package model

import "time"

type User struct {
	Id        int `gorm:"primaryKey"`
	Name      string
	Balance   int64
	Token     string
	CreatedAt time.Time
}
