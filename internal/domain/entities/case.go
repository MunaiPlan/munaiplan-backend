package entities

import "time"

type Case struct {
	ID                string              `json:"id"`
	CaseName          string              `json:"case_name"`
	CaseDescription   string              `json:"case_description"`
	DrillDepth        float64             `json:"drill_depth"`
	PipeSize          float64             `json:"pipe_size"`
	IsComplete        bool                `json:"is_complete"`
	CreatedAt         time.Time           `json:"created_at"`
	Fluids            []*Fluid            `json:"fluids"`
	Strings           []*String           `json:"strings"`
	Holes             []*Hole             `json:"holes"`
	PorePressures     []*PorePressure     `json:"pore_pressures"`
	FractureGradients []*FractureGradient `json:"fracture_gradients"`
	Rigs              []*Rig              `json:"rigs"`
}
