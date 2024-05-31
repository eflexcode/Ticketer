package model

import "time"

type User struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
}