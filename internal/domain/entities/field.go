package entities

// Месторождение (под компанией)
type Field struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	ReductionLevel  string  `json:"reduction_level"`
	ActiveFieldUnit string  `json:"active_field_unit"`
	Sites           []*Site `json:"sites"`
}
