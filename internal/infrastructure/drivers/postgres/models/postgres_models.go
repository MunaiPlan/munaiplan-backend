package models

import (
	"time"

	_ "ariga.io/atlas-provider-gorm/gormschema"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BaseModel defines common fields for all tables.
type BaseModel struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// Organization model with UUID primary key.
type Organization struct {
	BaseModel
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"type:varchar(255);not null" json:"email"`
	Phone     string    `gorm:"type:varchar(20)" json:"phone"`
	Address   string    `json:"address"`
	Companies []Company `gorm:"constraint:OnDelete:CASCADE;" json:"companies"`
	Users     []User    `gorm:"constraint:OnDelete:CASCADE;" json:"users"`
}

// User model with UUID primary key.
type User struct {
	BaseModel
	OrganizationID uuid.UUID    `gorm:"type:uuid;not null" json:"organization_id"`
	Name           string       `gorm:"not null" json:"name"`
	Surname        string       `gorm:"not null" json:"surname"`
	Email          string       `gorm:"type:varchar(255);not null" json:"email"`
	Password       string       `gorm:"type:varchar(70);not null" json:"password"`
	Phone          string       `gorm:"type:varchar(20)" json:"phone"`
}

// Company model with UUID primary key and foreign key.
type Company struct {
	BaseModel
	OrganizationID uuid.UUID    `gorm:"type:uuid;not null" json:"organization_id"`
	Name           string       `gorm:"type:varchar(255);not null" json:"name"`
	Division       string       `json:"division"`
	Group          string       `json:"group"`
	Representative string       `json:"representative"`
	Address        string       `json:"address"`
	Phone          string       `json:"phone"`
	Fields         []Field      `gorm:"constraint:OnDelete:CASCADE;" json:"fields"`
}

// Field model with UUID primary key and foreign key.
type Field struct {
	BaseModel
	CompanyID       uuid.UUID `gorm:"type:uuid;not null" json:"company_id"`
	Name            string    `gorm:"type:varchar(255);not null" json:"name"`
	Description     string    `json:"description"`
	ReductionLevel  string    `json:"reduction_level"`
	ActiveFieldUnit string    `json:"active_field_unit"`
	Sites           []Site    `gorm:"constraint:OnDelete:CASCADE;" json:"sites"`
}

// Site model with UUID primary key and foreign key.
type Site struct {
	BaseModel
	FieldID uuid.UUID `gorm:"type:uuid;not null" json:"field_id"`
	Name    string    `gorm:"type:varchar(255);not null" json:"name"`
	Area    float64   `json:"area"`
	Block   string    `json:"block"`
	Azimuth float64   `json:"azimuth"`
	Country string    `json:"country"`
	State   string    `json:"state"`
	Region  string    `json:"region"`
	Wells   []Well    `gorm:"constraint:OnDelete:CASCADE;" json:"wells"`
}

// Well model with UUID primary key and foreign key.
type Well struct {
	BaseModel
	SiteID                  uuid.UUID  `gorm:"type:uuid;not null" json:"site_id"`
	Name                    string     `gorm:"type:varchar(255);not null" json:"name"`
	Description             string     `json:"description"`
	Location                string     `json:"location"`
	UniversalWellIdentifier string     `json:"universal_well_identifier"`
	Type                    string     `json:"type"`
	WellNumber              string     `json:"well_number"`
	WorkingGroup            string     `json:"working_group"`
	ActiveWellUnit          string     `json:"active_well_unit"`
	Wellbores               []Wellbore `gorm:"constraint:OnDelete:CASCADE;" json:"wellbores"`
}

// Wellbore model with UUID primary key and foreign key.
type Wellbore struct {
	BaseModel
	Name                           string    `gorm:"type:varchar(255);not null" json:"name"`
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
	WellID                         uuid.UUID `gorm:"type:uuid;not null" json:"well_id"`
	Designs                        []Design  `gorm:"constraint:OnDelete:CASCADE;" json:"designs"`
}

// Design model with UUID primary key and foreign key.
type Design struct {
	BaseModel
	PlanName     string       `json:"plan_name"`
	Stage        string       `json:"stage"`
	Version      string       `json:"version"`
	ActualDate   time.Time    `json:"actual_date"`
	WellboreID   uuid.UUID    `gorm:"type:uuid;not null" json:"wellbore_id"`
	Trajectories []Trajectory `gorm:"constraint:OnDelete:CASCADE;" json:"trajectories"`
}

// Trajectory model with UUID primary key and foreign key.
type Trajectory struct {
	BaseModel
	Name        string             `gorm:"type:varchar(255)" json:"name"`
	Description string             `json:"description"`
	DesignID    uuid.UUID          `gorm:"type:uuid;not null" json:"design_id"`
	Headers     []TrajectoryHeader `gorm:"constraint:OnDelete:CASCADE;" json:"headers"`
	Units       []TrajectoryUnit   `gorm:"constraint:OnDelete:CASCADE;" json:"units"`
	Cases       []Case             `gorm:"constraint:OnDelete:CASCADE;" json:"cases"`
}

// TrajectoryHeader model with UUID primary key and foreign key.
type TrajectoryHeader struct {
	ID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	TrajectoryID     uuid.UUID `gorm:"type:uuid;not null" json:"trajectory_id"`
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
	CreatedAt        time.Time `gorm:"not null" json:"created_at"`
}

// TrajectoryUnit model with UUID primary key and foreign key.
type TrajectoryUnit struct {
	ID              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	TrajectoryID    uuid.UUID `gorm:"type:uuid;not null" json:"trajectory_id"`
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
	CreatedAt       time.Time `gorm:"not null" json:"created_at"`
}

// Case model with UUID primary key and foreign key.
type Case struct {
	BaseModel
	CaseName        string     `json:"case_name"`
	CaseDescription string     `json:"case_description"`
	DrillDepth      float64    `json:"drill_depth"`
	PipeSize        float64    `json:"pipe_size"`
	TrajectoryID    uuid.UUID  `gorm:"type:uuid;not null" json:"trajectory_id"`
	Holes           []Hole     `gorm:"constraint:OnDelete:CASCADE;" json:"holes"`
	Strings         []String   `gorm:"constraint:OnDelete:CASCADE;" json:"strings"`
	Fluids          []Fluid    `gorm:"constraint:OnDelete:CASCADE;" json:"fluids"`
}

// Hole model
type Hole struct {
	BaseModel
	CaseID                    uuid.UUID `gorm:"type:uuid;not null" json:"case_id"`
	Case                      Case      `gorm:"foreignKey:CaseID;constraint:OnDelete:CASCADE;" json:"case"`
	MDTop                     float64   `gorm:"not null" json:"md_top"`
	MDBase                    float64   `gorm:"not null" json:"md_base"`
	Length                    float64   `gorm:"not null" json:"length"`
	ShoeMD                    float64   `gorm:"not null" json:"shoe_md"`
	OD                        float64   `gorm:"not null" json:"od"`
	VD                        float64   `gorm:"not null" json:"vd"`
	CaisingInternalDiameter   float64   `gorm:"not null" json:"caising_internal_diameter"`
	DriftInternalDiameter     float64   `gorm:"not null" json:"drift_internal_diameter"`
	EffectiveHoleDiameter     float64   `gorm:"not null" json:"effective_hole_diameter"`
	Weight                    float64   `gorm:"not null" json:"weight"`
	Grade                     string    `gorm:"type:string;not null" json:"grade"`
	MinYieldStrength          float64   `gorm:"not null" json:"min_yield_strength"`
	BurstRating               float64   `gorm:"not null" json:"burst_rating"`
	CollapseRating            float64   `gorm:"not null" json:"collapse_rating"`
	FrictionFactorCasing      float64   `gorm:"not null" json:"friction_factor_casing"`
	LinearCapacityCasing      float64   `gorm:"not null" json:"linear_capacity_casing"`
	DescriptionCasing         string    `gorm:"type:string;not null" json:"description_casing"`
	ManufacturerCasing        string    `gorm:"type:string;not null" json:"manufacturer_casing"`
	ModelCasing               string    `gorm:"type:string;not null" json:"model_casing"`
	OpenHoleMDTop             float64   `gorm:"not null" json:"open_hole_md_top"`
	OpenHoleMDBase            float64   `gorm:"not null" json:"open_hole_md_base"`
	OpenHoleLength            float64   `gorm:"not null" json:"open_hole_length"`
	OpenHoleInternalDiameter  float64   `gorm:"not null" json:"open_hole_internal_diameter"`
	EffectiveDiameter         float64   `gorm:"not null" json:"effective_diameter"`
	FrictionFactorOpenHole    float64   `gorm:"not null" json:"friction_factor_open_hole"`
	LinearCapacityOpenHole    float64   `gorm:"not null" json:"linear_capacity_open_hole"`
	VolumeExcess              float64   `gorm:"not null" json:"volume_excess"`
	DescriptionOpenHole       string    `gorm:"type:string;not null" json:"description_open_hole"`
	TrippingInCasing          float64   `gorm:"not null" json:"tripping_in_casing"`
	TrippingOutCasing         float64   `gorm:"not null" json:"tripping_out_casing"`
	RotatingOnBottomCasing    float64   `gorm:"not null" json:"rotating_on_bottom_casing"`
	SlideDrillingCasing       float64   `gorm:"not null" json:"slide_drilling_casing"`
	BackReamingCasing         float64   `gorm:"not null" json:"back_reaming_casing"`
	RotatingOffBottomCasing   float64   `gorm:"not null" json:"rotating_off_bottom_casing"`
	TrippingInOpenHole        float64   `gorm:"not null" json:"tripping_in_open_hole"`
	TrippingOutOpenHole       float64   `gorm:"not null" json:"tripping_out_open_hole"`
	RotatingOnBottomOpenHole  float64   `gorm:"not null" json:"rotating_on_bottom_open_hole"`
	SlideDrillingOpenHole     float64   `gorm:"not null" json:"slide_drilling_open_hole"`
	BackReamingOpenHole       float64   `gorm:"not null" json:"back_reaming_open_hole"`
	RotatingOffBottomOpenHole float64   `gorm:"not null" json:"rotating_off_bottom_open_hole"`
}

// Strings table
type String struct {
	BaseModel
	Name     string    `gorm:"type:text;not null" json:"name"`
	Depth    float64   `gorm:"not null" json:"depth"`
	CaseID   uuid.UUID `gorm:"type:uuid;not null" json:"case_id"`
	Case     Case      `gorm:"foreignKey:CaseID;constraint:OnDelete:CASCADE;" json:"case"`
	Sections []Section `gorm:"constraint:OnDelete:CASCADE;" json:"sections"`
}

// SectionTypes table
type SectionType struct {
	BaseModel
	Name       string             `gorm:"type:varchar(255);not null" json:"name"`
	Attributes []SectionAttribute `gorm:"constraint:OnDelete:CASCADE;" json:"attributes"`
}

// SectionAttributes table
type SectionAttribute struct {
	BaseModel
	Name          string           `gorm:"type:varchar(255);not null" json:"name"`
	Unit          string           `gorm:"type:varchar(50)" json:"unit"`
	ValueTypeID   uuid.UUID        `gorm:"type:uuid;not null" json:"value_type_id"`
	ValueType     SectionValueType `gorm:"constraint:OnDelete:CASCADE;" json:"value_type"`
	SectionTypeID uuid.UUID        `gorm:"type:uuid;not null" json:"section_type_id"`
	SectionType   SectionType      `gorm:"constraint:OnDelete:CASCADE;" json:"section_type"`
}

// Sections table
type Section struct {
	BaseModel
	StringID      uuid.UUID   `gorm:"type:uuid;not null" json:"string_id"`
	String        String      `gorm:"foreignKey:StringID;constraint:OnDelete:CASCADE;" json:"string"`
	SectionTypeID uuid.UUID   `gorm:"type:uuid;not null" json:"section_type_id"`
	SectionType   SectionType `gorm:"foreignKey:SectionTypeID;constraint:OnDelete:CASCADE;" json:"section_type"`
	SectionValues []SectionValue `gorm:"constraint:OnDelete:CASCADE;" json:"section_values"`
}

// SectionValues table
type SectionValue struct {
	BaseModel
	SectionID   uuid.UUID        `gorm:"type:uuid;not null" json:"section_id"`
	Section     Section          `gorm:"foreignKey:SectionID;constraint:OnDelete:CASCADE;" json:"section"`
	AttributeID uuid.UUID        `gorm:"type:uuid;not null" json:"attribute_id"`
	Attribute   SectionAttribute `gorm:"foreignKey:AttributeID;constraint:OnDelete:CASCADE;" json:"attribute"`
	Value       string           `gorm:"type:text;not null" json:"value"`
}

// SectionValueTypes table
type SectionValueType struct {
	BaseModel
	Name string `gorm:"type:varchar(50);not null;unique" json:"name"` // e.g., "text", "float", "integer"
}

// Languages table
type Language struct {
	BaseModel
	Name string `gorm:"type:varchar(50);not null;unique" json:"name"`
}

// SectionAttributesU18n table (for translations)
type SectionAttributeU18n struct {
	BaseModel
	AttributeID uuid.UUID        `gorm:"type:uuid;not null" json:"attribute_id"`
	Attribute   SectionAttribute `gorm:"foreignKey:AttributeID;constraint:OnDelete:CASCADE;" json:"attribute"`
	LanguageID  uuid.UUID        `gorm:"type:uuid;not null" json:"language_id"`
	Language    Language         `gorm:"foreignKey:LanguageID;constraint:OnDelete:CASCADE;" json:"language"`
	Name        string           `gorm:"type:varchar(255);not null" json:"name"`
	Unit        string           `gorm:"type:varchar(50)" json:"unit"`
}

// Fluids table
type Fluid struct {
	BaseModel
	CaseID          uuid.UUID `gorm:"type:uuid;not null" json:"case_id"`
	Case            Case      `gorm:"foreignKey:CaseID;constraint:OnDelete:CASCADE;" json:"case"`
	Name            string    `gorm:"type:text;not null" json:"name"`
	Description     string    `gorm:"type:text" json:"description"`
	Density         float64   `gorm:"type:decimal;not null" json:"density"`
	FluidBaseTypeID uuid.UUID `gorm:"type:uuid;not null" json:"fluid_base_type_id"`
	FluidBaseType   FluidType `gorm:"foreignKey:FluidBaseTypeID;constraint:OnDelete:CASCADE;" json:"fluid_base_type"`
	BaseFluidID     uuid.UUID `gorm:"type:uuid;not null" json:"base_fluid_id"`
	BaseFluid       FluidType `gorm:"foreignKey:BaseFluidID;constraint:OnDelete:CASCADE;" json:"base_fluid"`
}

// FluidTypes table
type FluidType struct {
	BaseModel
	Name string `gorm:"type:text;not null" json:"name"`
}
