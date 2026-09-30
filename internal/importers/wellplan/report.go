// Package wellplan parses Landmark WellPlan exports (the .docx engineering report and the
// tab-separated survey report) into one normalised Report. It has no database or HTTP
// dependencies so it can be tested in isolation and reused by other entry points.
//
// Units are kept exactly as WellPlan's metric exports print them and are not converted:
// depths and lengths in m, diameters in mm, linear weight in kg/m, angles in degrees,
// dogleg in °/30 m, yield stress in psi, density in kg/m³, torque in kN·m, loads in tonnes.
package wellplan

// Units documents the unit of every numeric field in a Report.
var Units = map[string]string{
	"depth": "m", "length": "m", "diameter": "mm", "linear_weight": "kg/m",
	"angle": "deg", "dogleg": "deg/30m", "yield_stress": "psi", "density": "kg/m3",
	"torque": "kN-m", "load": "tonne", "pressure": "atm", "temperature": "degC",
	"plastic_viscosity": "cP", "yield_point": "lbf/100ft2", "flow_rate": "L/s",
}

// Report is everything MunaiPlan can use from one WellPlan case.
type Report struct {
	Source     []SourceFile       `json:"sources"`
	Language   string             `json:"language,omitempty"` // "ru" or "en" for .docx reports
	Case       CaseInfo           `json:"case"`
	Survey     []Station          `json:"survey"`
	SurveyInfo SurveyHeader       `json:"survey_header"`
	Holes      []HoleSection      `json:"hole_sections"`
	String     []Component        `json:"string"`
	Grades     map[string]float64 `json:"grades_min_yield_psi"`
	Fluid      *Fluid             `json:"fluid,omitempty"`
	Geothermal *Geothermal        `json:"geothermal,omitempty"`
	Friction   map[string]float64 `json:"friction_factors"`
	TorqueDrag *TorqueDragResults `json:"torque_drag,omitempty"`
	Hydraulics map[string]string  `json:"hydraulics,omitempty"`
	Warnings   []string           `json:"warnings"`
}

type SourceFile struct {
	Name   string `json:"name"`
	Kind   string `json:"kind"` // "report" or "survey"
	SHA256 string `json:"sha256"`
	Bytes  int64  `json:"bytes"`
}

// CaseInfo is the WellPlan hierarchy and headline values for the case.
type CaseInfo struct {
	Company         string   `json:"company"`
	Field           string   `json:"field"`
	Site            string   `json:"site"`
	Well            string   `json:"well"`
	Wellbore        string   `json:"wellbore"`
	Design          string   `json:"design"`
	Case            string   `json:"case"`
	MD              *float64 `json:"md,omitempty"`
	TVD             *float64 `json:"tvd,omitempty"`
	AirGap          *float64 `json:"air_gap,omitempty"`
	GroundElevation *float64 `json:"ground_elevation,omitempty"`
	DatumElevation  *float64 `json:"datum_elevation,omitempty"`
	Datum           string   `json:"datum,omitempty"`
	WellType        string   `json:"well_type,omitempty"`
}

// SurveyHeader holds the key/value header of a WellPlan survey export.
type SurveyHeader struct {
	Customer         string   `json:"customer,omitempty"`
	Project          string   `json:"project,omitempty"`
	ProfileType      string   `json:"profile_type,omitempty"`
	Field            string   `json:"field,omitempty"`
	YourRef          string   `json:"your_ref,omitempty"`
	Structure        string   `json:"structure,omitempty"`
	JobNumber        string   `json:"job_number,omitempty"`
	Wellhead         string   `json:"wellhead,omitempty"`
	KellyBushingElev *float64 `json:"kelly_bushing_elev,omitempty"`
	Profile          string   `json:"profile,omitempty"`
}

// Station is one survey point. Optional values are nil when the source does not provide them.
type Station struct {
	MD              float64  `json:"md"`
	Inc             float64  `json:"inc"`
	Azi             float64  `json:"azi"`
	TVD             float64  `json:"tvd"`
	SubSea          *float64 `json:"sub_sea,omitempty"`
	NS              *float64 `json:"ns,omitempty"`
	EW              *float64 `json:"ew,omitempty"`
	GlobalN         *float64 `json:"global_n,omitempty"`
	GlobalE         *float64 `json:"global_e,omitempty"`
	DLS             *float64 `json:"dls,omitempty"`
	VerticalSection *float64 `json:"vertical_section,omitempty"`
}

// HoleSection is a cased or open-hole interval from the report's hole-section table.
type HoleSection struct {
	Type              string   `json:"type"` // "Casing", "Open Hole", "Liner", ...
	Depth             float64  `json:"depth"`
	Length            float64  `json:"length"`
	ShoeDepth         *float64 `json:"shoe_depth,omitempty"`
	InnerDiameter     float64  `json:"inner_diameter"`
	Drift             *float64 `json:"drift,omitempty"`
	EffectiveDiameter *float64 `json:"effective_diameter,omitempty"`
	FrictionFactor    *float64 `json:"friction_factor,omitempty"`
	LinearCapacity    *float64 `json:"linear_capacity,omitempty"` // L/m
	VolumeExcess      *float64 `json:"volume_excess,omitempty"`   // %
}

// IsOpenHole reports whether the section is uncased.
func (h HoleSection) IsOpenHole() bool { return h.Type == "Open Hole" }

// Component is one element of the work string, listed from surface to bit.
// Depth is the measured depth of the component's bottom, as WellPlan prints it.
type Component struct {
	Type           string   `json:"type"`
	Length         float64  `json:"length"`
	Depth          float64  `json:"depth"`
	BodyOD         float64  `json:"body_od"`
	BodyID         *float64 `json:"body_id,omitempty"`
	AvgJointLength *float64 `json:"avg_joint_length,omitempty"`
	JointLength    *float64 `json:"joint_length,omitempty"` // "Stabilizer / Tool Joint" length
	JointOD        *float64 `json:"joint_od,omitempty"`
	JointID        *float64 `json:"joint_id,omitempty"`
	Weight         *float64 `json:"weight,omitempty"`
	Material       string   `json:"material,omitempty"`
	Grade          string   `json:"grade,omitempty"`
	Class          string   `json:"class,omitempty"`
	MinYield       *float64 `json:"min_yield,omitempty"` // psi, resolved from the grade table
}

type Fluid struct {
	Name             string   `json:"name"`
	Type             string   `json:"type,omitempty"`
	Base             string   `json:"base,omitempty"`
	BaseFluid        string   `json:"base_fluid,omitempty"`
	RheologyModel    string   `json:"rheology_model,omitempty"`
	Temperature      *float64 `json:"temperature,omitempty"`
	Pressure         *float64 `json:"pressure,omitempty"`
	Density          *float64 `json:"density,omitempty"`
	PlasticViscosity *float64 `json:"plastic_viscosity,omitempty"`
	YieldPoint       *float64 `json:"yield_point,omitempty"`
}

type Geothermal struct {
	AmbientTemperature *float64 `json:"ambient_temperature,omitempty"`
	TemperatureAtDepth *float64 `json:"temperature_at_depth,omitempty"`
	Depth              *float64 `json:"depth,omitempty"`    // TVD of the temperature reading
	Gradient           *float64 `json:"gradient,omitempty"` // degC/100m
}

// Operation codes shared by reference results and, later, MunaiPlan predictions.
const (
	OpTrippingIn        = "tripping_in"
	OpTrippingOut       = "tripping_out"
	OpRotatingOnBottom  = "rotating_on_bottom"
	OpSlideDrilling     = "slide_drilling"
	OpRotatingOffBottom = "rotating_off_bottom"
	OpBackReaming       = "back_reaming"
)

// LoadCondition is one row of WellPlan's Torque & Drag load summary: the reference answer.
type LoadCondition struct {
	Operation          string   `json:"operation"`
	Label              string   `json:"label"`
	Failures           []string `json:"failures,omitempty"`
	SurfaceTorque      *float64 `json:"surface_torque"` // kN-m at the rotary table
	WindupWithBit      *float64 `json:"windup_with_bit,omitempty"`
	WindupWithoutBit   *float64 `json:"windup_without_bit,omitempty"`
	HookLoad           *float64 `json:"hook_load"`         // tonne, "measured weight"
	Stretch            *float64 `json:"stretch,omitempty"` // m
	NeutralFromSurface *float64 `json:"neutral_point_from_surface,omitempty"`
	NeutralFromBit     *float64 `json:"neutral_point_from_bit,omitempty"`
}

// TorqueDragResults are WellPlan's own outputs and settings for the case.
type TorqueDragResults struct {
	BitDepth    *float64           `json:"bit_depth,omitempty"`
	BlockWeight *float64           `json:"block_weight,omitempty"`
	FlowRate    *float64           `json:"flow_rate,omitempty"`
	WOB         map[string]float64 `json:"wob,omitempty"`        // tonne per drilling operation
	BitTorque   map[string]float64 `json:"bit_torque,omitempty"` // kN-m per drilling operation
	TripSpeed   map[string]float64 `json:"trip_speed,omitempty"` // m/min per tripping operation
	Loads       []LoadCondition    `json:"load_summary"`
	Limits      map[string]string  `json:"mechanical_limits,omitempty"`
}
