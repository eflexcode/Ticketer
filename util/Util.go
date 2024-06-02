package util

import "time"

type User struct {
	CreatedAt       time.Time
	Email           string `json:"email" validate:"required"`
	Password        string `json:"password" validate:"required"`
	Username        string `json:"username" validate:"required"`
	ProfileImageUrl string `json:"profile_image_url"`
	CoverImageUrl   string `json:"cover_image_url"`
}
