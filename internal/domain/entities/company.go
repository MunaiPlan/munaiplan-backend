package entities

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
	Fields         []*Field  `json:"fields"`
}
