package domain

import (
	"time"
)

// Компания (под юзером)
// TODO() Correct all tables in this file
type Company struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Division       string    `json:"division"`
	Group          string    `json:"group"`
	Representative string    `json:"representative"`
	Address        string    `json:"address"`
	Phone          string    `json:"phone"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Fields         []*Field  `json:"fields"`
}
