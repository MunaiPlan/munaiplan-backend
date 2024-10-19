package requests

type EffectiveTensionFromMLModelRequest struct {
	MD                    []float64 `json:"MD" binding:"required"`
	Incl                  []float64 `json:"Incl" binding:"required"`
	Azim                  []float64 `json:"Azim" binding:"required"`
	SubSea                []float64 `json:"Sub_Sea" binding:"required"`
	TVD                   []float64 `json:"TVD" binding:"required"`
	LocalNCoord           []float64 `json:"Local_N_Coord" binding:"required"`
	LocalECoord           []float64 `json:"Local_E_Coord" binding:"required"`
	GlobalNCoord          []float64 `json:"Global_N_Coord" binding:"required"`
	GlobalECoord          []float64 `json:"Global_E_Coord" binding:"required"`
	Dogleg                []float64 `json:"Dogleg" binding:"required"`
	VerticalSection       []float64 `json:"Vertical_Section" binding:"required"`
	BodyOD                []float64 `json:"Body_OD" binding:"required"`
	BodyID                []float64 `json:"Body_ID" binding:"required"`
	BodyAvgJointLength    []float64 `json:"Body_AvgJointLength" binding:"required"`
	StabilizerLength      []float64 `json:"Stabilizer_Length" binding:"required"`
	StabilizerOD          []float64 `json:"Stabilizer_OD" binding:"required"`
	StabilizerID          []float64 `json:"Stabilizer_ID" binding:"required"`
	Weight                []float64 `json:"Weight" binding:"required"`
	CoefficientOfFriction []float64 `json:"Coefficient_of_Friction" binding:"required"`
	MinimumYieldStrength  []float64 `json:"Minimum_Yield_Strength" binding:"required"`
}
