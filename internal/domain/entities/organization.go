package entities

import (
	"time"
)

// Пользователь (База данных в книге)
// TODO() Correct all tables in this file
type Organization struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
	Address   string     `json:"address"`
	CreatedAt time.Time  `json:"registeredAt"`
	Companies []*Company `json:"companies"`
	Users     []*User    `json:"users"`
}
