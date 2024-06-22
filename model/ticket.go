package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ticket struct {
	ID        primitive.ObjectID `bson:"_id"`
	EventID   string             `json:"event_id"`
	TicketID  string             `json:"ticket_id"`
	BuyDate   string             `json:"buy_date"`
	BoughtBy  string             `json:"bought_by"`
	BoughtFor string             `json:"bought_for"`
}
