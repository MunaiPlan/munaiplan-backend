package entities

import (
	"time"
)

// Ствол Скважины (под скважиной)
// TODO() Correct all tables in this file
type Wellbore struct {
	ID                             string    `json:"id"`
	Name                           string    `json:"name"`
	BottomHoleLocation             string    `json:"bottom_hole_location"`
	WellboreDepth                  float64   `json:"wellbore_depth"`
	AverageHookLoad                float64   `json:"average_hook_load"`
	RiserPressure                  float64   `json:"riser_pressure"`
	AverageInletFlow               float64   `json:"average_inlet_flow"`
	AverageColumnRotationFrequency float64   `json:"average_column_rotation_frequency"`
	MaximumColumnRotationFrequency float64   `json:"maximum_column_rotation_frequency"`
	AverageWeightOnBit             float64   `json:"average_weight_on_bit"`
	MaximumWeightOnBit             float64   `json:"maximum_weight_on_bit"`
	AverageTorque                  float64   `json:"average_torque"`
	MaximumTorque                  float64   `json:"maximum_torque"`
	DownStaticFriction             float64   `json:"down_static_friction"`
	DepthInterval                  float64   `json:"depth_interval"`
	Designs                        []*Design `json:"designs"`
	CreatedAt                      time.Time `json:"created_at"`
}
