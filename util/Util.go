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

type Ticket struct {
	EventID      int       `json:"event_id"`
	TicketNumber int       `json:"ticket_number"`
	BuyDate      time.Time `json:"buy_date"`
	BoughtBy     string    `json:"bought_by"`
	BoughtFor    string    `json:"bought_for"`
}

type Event struct {
	NumberOFTicketPrinted   int       `json:"number_of_ticket_printed"`
	NumberOfTicketSold      int       `json:"number_of_ticket_sold"`
	NumberOfTicketAvailable int       `json:"number_of_ticket_available"`
	EvenName                string    `json:"even_name"`
	EventDescription        string    `json:"event_description"`
	EventCoverImage         string    `json:"event_cover_image"`
	EventAddress            string    `json:"event_address"`
	EventDate               time.Time `json:"event_date"`
	TicketStartSalesDate    time.Time `json:"ticket_start_sales_date"`
	TicketEndSalesDate      time.Time `json:"ticket_end_sales_date"`
	OrganisationId          int       `json:"organisation_id" validate:"required"`
}
