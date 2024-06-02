package model

import "time"

type User struct {
	ID                 uint `json:"id" gorm:"primaryKey"`
	CreatedAt          time.Time
	Email              string   `json:"email" gorm:"unique_index"`
	Password           string   `json:"password"`
	Username           string   `json:"username"`
	ProfileImageUrl    string   `json:"profile_image_url"`
	CoverImageUrl      string   `json:"cover_image_url"`
	Tickets            []Ticket `json:"tickets" gorm:"foreignKey:ID"`
	EventsInterestedIn []string `json:"events_interested_in" gorm:"type:text[]"`
}
