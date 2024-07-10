package domain

import (
	"time"
)

// Пользователь (База данных в книге)
// TODO() Correct all tables in this file
type User struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"registeredAt"`
	Companies []*Company `json:"companies"`
}
