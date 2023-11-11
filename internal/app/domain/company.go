package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO() Correct all tables in this file
type Company struct {
	ID             primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Name           string               `json:"name" bson:"name"`
	Division       string               `json:"division" bson:"division"`
	Group          string               `json:"group" bson:"group"`
	Representative string               `json:"representative" bson:"representative"`
	Address        string               `json:"address" bson:"address"`
	Phone          string               `json:"phone" bson:"phone"`
	AdminIDs       []primitive.ObjectID `json:"adminIds" bson:"adminIds,omitempty"`
	AuditInfo      AuditInfo            `json:"auditInfo" bson:"auditInfo"`
}

type AuditInfo struct {
	CreatedAt    time.Time `json:"createdAt" bson:"createdAt"`
	ModifiedAt   time.Time `json:"modifiedAt" bson:"modifiedAt"`
	ModifiedBy   string    `json:"modifiedBy" bson:"modifiedBy"`
	ModifiedByID primitive.ObjectID `json:"modifiedById" bson:"modifiedById,omitempty"`
}