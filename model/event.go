package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID                      primitive.ObjectID `bson:"_id"`
	NumberOFTicketPrinted   int                `json:"number_of_ticket_printed"`
	NumberOfTicketSold      int                `json:"number_of_ticket_sold"`
	NumberOfTicketAvailable int                `json:"number_of_ticket_available"`
	EvenName                string             `json:"even_name"`
	EventDescription        string             `json:"event_description"`
	EventCoverImage         string             `json:"event_cover_image"`
	EventAddress            string             `json:"event_address"`
	EventDate               string             `json:"event_date"`
	TicketStartSalesDate    string             `json:"ticket_start_sales_date"`
	TicketEndSalesDate      string             `json:"ticket_end_sales_date"`
	OrganisationId          int                `json:"organisation_id"`
}
