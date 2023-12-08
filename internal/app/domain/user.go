package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Пользователь (База данных в книге)
// TODO() Correct all tables in this file
type User struct {
	ID           primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Name         string               `json:"name" bson:"name"`
	Email        string               `json:"email" bson:"email"`
	Phone        string               `json:"phone" bson:"phone"`
	Password     string               `json:"password" bson:"password"`
	RegisteredAt time.Time            `json:"registeredAt" bson:"registeredAt"`
	LastVisitAt  time.Time            `json:"lastVisitAt" bson:"lastVisitAt"`
	Companies      []primitive.ObjectID `json:"companies" bson:"companies"`
}