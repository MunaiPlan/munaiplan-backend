package requests

type EffectiveTensionFromMLModelRequest struct {
	MD                    []float64 `json:"md" binding:"required"`
	Incl                  []float64 `json:"incl" binding:"required"`
	Azim                  []float64 `json:"azim" binding:"required"`
	SubSea                []float64 `json:"sub_sea" binding:"required"`
	TVD                   []float64 `json:"tvd" binding:"required"`
	LocalNCoord           []float64 `json:"local_n_coord" binding:"required"`
	LocalECoord           []float64 `json:"local_e_coord" binding:"required"`
	GlobalNCoord          []float64 `json:"global_n_coord" binding:"required"`
	GlobalECoord          []float64 `json:"global_e_coord" binding:"required"`
	Dogleg                []float64 `json:"dogleg" binding:"required"`
	VerticalSection       []float64 `json:"vertical_section" binding:"required"`
	BodyOD                []float64 `json:"body_od" binding:"required"`
	BodyID                []float64 `json:"body_id" binding:"required"`
	BodyAvgJointLength    []float64 `json:"body_avg_joint_length" binding:"required"`
	StabilizerLength      []float64 `json:"stabilizer_length" binding:"required"`
	StabilizerOD          []float64 `json:"stabilizer_od" binding:"required"`
	StabilizerID          []float64 `json:"stabilizer_id" binding:"required"`
	Weight                []float64 `json:"weight" binding:"required"`
	CoefficientOfFriction []float64 `json:"coefficient_of_friction" binding:"required"`
	MinimumYieldStrength  []float64 `json:"minimum_yield_stress" binding:"required"`
}
