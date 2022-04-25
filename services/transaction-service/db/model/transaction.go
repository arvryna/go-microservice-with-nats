package model

import "time"

type Transaction struct {
	Id                int `gorm:"primaryKey"`
	TransactionAmount int64
	IsUp              bool
	UserId            int
	Before            int64
	After             int64
	CreatedAt         time.Time
}
