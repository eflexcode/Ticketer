package model

import "time"

type Organisation struct {
	ID                          uint    `json:"id" gorm:"primaryKey"`
	OrganisationName            string  `json:"organisation_name"`
	OrganisationEmail           string  `json:"organisation_email"`
	OrganisationPassword        string  `json:"organisation_password"`
	OrganisationAddress         string  `json:"organisation_address"`
	OrganisationProfileImageUrl string  `json:"organisation_profile_image_url"`
	OrganisationOverImageUrl    string  `json:"organisation_cover_image_url"`
	OrganisationDescription     string  `json:"organisation_description"`
	Events                      []Event `json:"events" gorm:"foreignKey:ID;polymorphic:Owner"`
	CreatedAt                   time.Time
}
