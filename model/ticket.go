package model

type Ticket struct {
	ID           uint  `json:"id" gorm:"primaryKey"`
	Event        Event `json:"event"`
	TicketNumber int   `json:"ticket_number"`
}
