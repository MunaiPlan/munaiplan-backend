package entities

import "time"

// OrganizationSummary is the administrator's view of a tenant.
type OrganizationSummary struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UserCount int64     `json:"user_count"`
}
