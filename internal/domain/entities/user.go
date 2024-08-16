package entities

import (
	"time"
)

// Пользователь (База данных в книге)
// TODO() Correct all tables in this file
type User struct {
	ID             string    `json:"id"`
	OrganizationID string    `json:"organizationId"`
	Name           string    `json:"name"`
	Surname        string    `json:"surname"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	Password       string    `json:"password"`
	CreatedAt      time.Time `json:"registeredAt"`
}
