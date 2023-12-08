package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Месторождение (под компанией)
type Field struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	// General Information
	Name           string `json:"name" bson:"name"`
	Description       string `json:"description" bson:"description"`
	Location          string `json:"location" bson:"location"`
	Representative string `json:"representative" bson:"representative"`
	Address        string `json:"address" bson:"address"`
	Phone          string `json:"phone" bson:"phone"`
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt" bson:"modifiedAt"`
	Fields      []primitive.ObjectID `json:"fields" bson:"fields"`
}