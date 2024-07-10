package domain

import "time"

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
