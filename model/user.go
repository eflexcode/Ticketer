package model

import "time"

type User struct {
	ID                uint `json:"id" gorm:"primaryKey"`
	CreatedAt         time.Time
	Email             string   `json:"email"`
	Username          string   `json:"username"`
	ProfileImageUrl   string   `json:"profile_image_url"`
	CoverImageUrl     string   `json:"cover_image_url"`
	Tickets           []Ticket `json:"tickets"`
	EventInterestedIn []string `json:"event_interested_in"`
}
