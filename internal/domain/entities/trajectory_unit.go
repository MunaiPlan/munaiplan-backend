package domain

import "time"

type TrajectoryUnit struct {
	ID              string    `json:"id"`
	MD              float64   `json:"md"`
	Incl            float64   `json:"incl"`
	Azim            float64   `json:"azim"`
	SubSea          float64   `json:"sub_sea"`
	TVD             float64   `json:"tvd"`
	LocalNCoord     float64   `json:"local_n_coord"`
	LocalECoord     float64   `json:"local_e_coord"`
	GlobalNCoord    float64   `json:"global_n_coord"`
	GlobalECoord    float64   `json:"global_e_coord"`
	Dogleg          float64   `json:"dogleg"`
	VerticalSection float64   `json:"vertical_section"`
	CreatedAt       time.Time `json:"created_at"`
}
