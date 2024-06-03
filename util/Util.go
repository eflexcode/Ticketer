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

type Organisation struct {
	CreatedAt                   time.Time
	OrganisationName            string `json:"organisation_name" validate:"required"`
	OrganisationEmail           string `json:"organisation_email" validate:"required"`
	OrganisationPassword        string `json:"organisation_password" validate:"required"`
	OrganisationAddress         string `json:"organisation_address" validate:"required"`
	OrganisationProfileImageUrl string `json:"organisation_profile_image_url"`
	OrganisationOverImageUrl    string `json:"organisation_cover_image_url"`
	OrganisationDescription     string `json:"organisation_description" validate:"required"`
}
