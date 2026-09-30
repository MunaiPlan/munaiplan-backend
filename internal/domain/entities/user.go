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
	Password       string    `json:"-"`
	Role           string    `json:"role"`
	CreatedAt      time.Time `json:"registeredAt"`
}

// Roles. Administrators provision organizations and accounts; users work within one organization.
const (
	RoleUser  = "user"
	RoleAdmin = "admin"
)
