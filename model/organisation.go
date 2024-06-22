package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Organisation struct {
	ID                          primitive.ObjectID `bson:"_id"`
	OrganisationName            string             `json:"organisation_name"`
	OrganisationEmail           string             `json:"organisation_email"`
	OrganisationPassword        string             `json:"organisation_password"`
	OrganisationAddress         string             `json:"organisation_address"`
	OrganisationProfileImageUrl string             `json:"organisation_profile_image_url"`
	OrganisationOverImageUrl    string             `json:"organisation_cover_image_url"`
	OrganisationDescription     string             `json:"organisation_description"`
	Events                      []string           `json:"events" gorm:"foreignKey:ID;polymorphic:Owner"`
	CreatedAt                   time.Time
}
