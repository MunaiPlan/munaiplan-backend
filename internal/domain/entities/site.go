package entities

// TODO() Correct all tables in this file

// Куст (под месторождением)
type Site struct {
	ID      string  `json:"field_id"`
	Name    string  `json:"name"`
	Area    float64 `json:"area"`
	Block   string  `json:"block"`
	Azimuth float64 `json:"azimuth"`
	Country string  `json:"country"`
	State   string  `json:"state"`
	Region  string  `json:"region"`
	Wells   []*Well `json:"wells"`
}
