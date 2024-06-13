package model

import (
	"time"
)

type User struct {
	//ID                 primitive.ObjectID `bson:"_id"`
	CreatedAt          time.Time
	Email              string   `json:"email"`
	Password           string   `json:"password"`
	Username           string   `json:"username"`
	ProfileImageUrl    string   `json:"profile_image_url"`
	CoverImageUrl      string   `json:"cover_image_url"`
	Tickets            []string `json:"tickets"`
	EventsInterestedIn []string `json:"events_interested_in"`
}
