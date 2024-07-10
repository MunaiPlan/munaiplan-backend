package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Компания (под юзером)
// TODO() Correct all tables in this file
type Company struct {
	// General Information
	ID             primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Name           string               `json:"name" bson:"name"`
	Division       string               `json:"division" bson:"division"`
	Group          string               `json:"group" bson:"group"`
	Representative string               `json:"representative" bson:"representative"`
	Address        string               `json:"address" bson:"address"`
	Phone          string               `json:"phone" bson:"phone"`
	CreatedAt      time.Time            `json:"createdAt" bson:"createdAt"`
	ModifiedAt     time.Time            `json:"modifiedAt" bson:"modifiedAt"`
	Fields         []primitive.ObjectID `json:"fields" bson:"fields"`
}
