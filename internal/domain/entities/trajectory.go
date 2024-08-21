package entities

import "time"

type Trajectory struct {
	ID          string              `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Headers     []*TrajectoryHeader `json:"headers"`
	Units       []*TrajectoryUnit   `json:"units"`
	Cases       []*Case             `json:"cases"`
	CreatedAt   time.Time           `json:"created_at"`
}

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

type TrajectoryHeader struct {
	ID               string    `json:"id"`
	Customer         string    `json:"customer"`
	Project          string    `json:"project"`
	ProfileType      string    `json:"profile_type"`
	Field            string    `json:"field"`
	YourRef          string    `json:"your_ref"`
	Structure        string    `json:"structure"`
	JobNumber        string    `json:"job_number"`
	Wellhead         string    `json:"wellhead"`
	KellyBushingElev float64   `json:"kelly_bushing_elev"`
	Profile          string    `json:"profile"`
	CreatedAt        time.Time `json:"created_at"`
}
