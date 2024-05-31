package model

type Event struct {
	ID                      uint   `json:"id" gorm:"primaryKey"`
	NumberOFTicketPrinted   int    `json:"number_of_ticket_printed"`
	NumberOfTicketSold      int    `json:"number_of_ticket_sold"`
	NumberOfTicketAvailable int    `json:"number_of_ticket_available"`
	EvenName                string `json:"even_name"`
	EventDescription        string `json:"event_description"`
	EventCoverImage         string `json:"event_cover_image"`
}
