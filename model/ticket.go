package model

import "time"

type Ticket struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Event        Event     `json:"event"`
	TicketNumber int       `json:"ticket_number"`
	BuyDate      time.Time `json:"buy_date"`
	BoughtBy     string    `json:"bought_by"`
	BoughtFor    string    `json:"bought_for"`
}
