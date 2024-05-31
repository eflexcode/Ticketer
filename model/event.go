package model

import "time"

type Event struct {
	ID                      uint      `json:"id" gorm:"primaryKey"`
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
	OrganisationId          string    `json:"organisation_id"`
}
