package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Organization model with UUID primary key.
type Organization struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Name      string         `gorm:"not null" json:"name"`
	Email     string         `gorm:"type:varchar(255);not null" json:"email"`
	Phone     string         `gorm:"type:varchar(20)" json:"phone"`
	Address   string         `json:"address"`
	Companies []Company      `gorm:"constraint:OnDelete:CASCADE;" json:"companies"`
	Users     []User         `gorm:"constraint:OnDelete:CASCADE;" json:"users"`
}

// User model with UUID primary key.
type User struct {
	ID             uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt      time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	OrganizationID uuid.UUID      `gorm:"type:uuid;not null" json:"organization_id"`
	Name           string         `gorm:"not null" json:"name"`
	Surname        string         `gorm:"not null" json:"surname"`
	Email          string         `gorm:"type:varchar(255);not null" json:"email"`
	Password       string         `gorm:"type:varchar(70);not null" json:"password"`
	Phone          string         `gorm:"type:varchar(20)" json:"phone"`
}

// Company model with UUID primary key and foreign key.
type Company struct {
	ID             uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt      time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	OrganizationID uuid.UUID      `gorm:"type:uuid;not null" json:"organization_id"`
	Name           string         `gorm:"type:varchar(255);not null" json:"name"`
	Division       string         `json:"division"`
	Group          string         `json:"group"`
	Representative string         `json:"representative"`
	Address        string         `json:"address"`
	Phone          string         `json:"phone"`
	Fields         []Field        `gorm:"constraint:OnDelete:CASCADE;" json:"fields"`
}

// Field model with UUID primary key and foreign key.
type Field struct {
	ID              uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CompanyID       uuid.UUID      `gorm:"type:uuid;not null" json:"company_id"`
	Name            string         `gorm:"type:varchar(255);not null" json:"name"`
	Description     string         `json:"description"`
	ReductionLevel  string         `json:"reduction_level"`
	ActiveFieldUnit string         `json:"active_field_unit"`
	Sites           []Site         `gorm:"constraint:OnDelete:CASCADE;" json:"sites"`
}

// Site model with UUID primary key and foreign key.
type Site struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	FieldID   uuid.UUID      `gorm:"type:uuid;not null" json:"field_id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	Area      float64        `json:"area"`
	Block     string         `json:"block"`
	Azimuth   float64        `json:"azimuth"`
	Country   string         `json:"country"`
	State     string         `json:"state"`
	Region    string         `json:"region"`
	Wells     []Well         `gorm:"constraint:OnDelete:CASCADE;" json:"wells"`
}

// Well model with UUID primary key and foreign key.
type Well struct {
	ID                      uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt               time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt               time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt               gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	SiteID                  uuid.UUID      `gorm:"type:uuid;not null" json:"site_id"`
	Name                    string         `gorm:"type:varchar(255);not null" json:"name"`
	Description             string         `json:"description"`
	Location                string         `json:"location"`
	UniversalWellIdentifier string         `json:"universal_well_identifier"`
	Type                    string         `json:"type"`
	WellNumber              string         `json:"well_number"`
	WorkingGroup            string         `json:"working_group"`
	ActiveWellUnit          string         `json:"active_well_unit"`
	Wellbores               []Wellbore     `gorm:"constraint:OnDelete:CASCADE;" json:"wellbores"`
}

// Wellbore model with UUID primary key and foreign key.
type Wellbore struct {
	ID                             uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt                      time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt                      time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt                      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	WellID                         uuid.UUID      `gorm:"type:uuid;not null" json:"well_id"`
	Name                           string         `gorm:"type:varchar(255);not null" json:"name"`
	BottomHoleLocation             string         `json:"bottom_hole_location"`
	WellboreDepth                  float64        `json:"wellbore_depth"`
	AverageHookLoad                float64        `json:"average_hook_load"`
	RiserPressure                  float64        `json:"riser_pressure"`
	AverageInletFlow               float64        `json:"average_inlet_flow"`
	AverageColumnRotationFrequency float64        `json:"average_column_rotation_frequency"`
	MaximumColumnRotationFrequency float64        `json:"maximum_column_rotation_frequency"`
	AverageWeightOnBit             float64        `json:"average_weight_on_bit"`
	MaximumWeightOnBit             float64        `json:"maximum_weight_on_bit"`
	AverageTorque                  float64        `json:"average_torque"`
	MaximumTorque                  float64        `json:"maximum_torque"`
	DownStaticFriction             float64        `json:"down_static_friction"`
	DepthInterval                  float64        `json:"depth_interval"`
	Designs                        []Design       `gorm:"constraint:OnDelete:CASCADE;" json:"designs"`
}

// Design model with UUID primary key and foreign key.
type Design struct {
	ID           uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	WellboreID   uuid.UUID      `gorm:"type:uuid;not null" json:"wellbore_id"`
	PlanName     string         `json:"plan_name"`
	Stage        string         `json:"stage"`
	Version      string         `json:"version"`
	ActualDate   time.Time      `json:"actual_date"`
	Trajectories []Trajectory   `gorm:"constraint:OnDelete:CASCADE;" json:"trajectories"`
}

// Trajectory model with UUID primary key and foreign key.
type Trajectory struct {
	ID          uuid.UUID          `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt   time.Time          `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time          `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt     `gorm:"index" json:"deleted_at"`
	DesignID    uuid.UUID          `gorm:"type:uuid;not null" json:"design_id"`
	Name        string             `gorm:"type:varchar(255)" json:"name"`
	Description string             `json:"description"`
	Headers     []TrajectoryHeader `gorm:"constraint:OnDelete:CASCADE;" json:"headers"`
	Units       []TrajectoryUnit   `gorm:"constraint:OnDelete:CASCADE;" json:"units"`
	Cases       []Case             `gorm:"constraint:OnDelete:CASCADE;" json:"cases"`
}

// TrajectoryHeader model with UUID primary key and foreign key.
type TrajectoryHeader struct {
	ID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`
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
}

// TrajectoryUnit model with UUID primary key and foreign key.
type TrajectoryUnit struct {
	ID              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
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
}

// Case model with UUID primary key and foreign key.
type Case struct {
	ID                uuid.UUID          `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt         time.Time          `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time          `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt         gorm.DeletedAt     `gorm:"index" json:"deleted_at"`
	CaseName          string             `json:"case_name"`
	CaseDescription   string             `json:"case_description"`
	DrillDepth        float64            `json:"drill_depth"`
	PipeSize          float64            `json:"pipe_size"`
	IsComplete        bool               `gorm:"not null;default:false" json:"is_complete"`
	TrajectoryID      uuid.UUID          `gorm:"type:uuid;not null" json:"trajectory_id"`
	Holes             []Hole             `gorm:"constraint:OnDelete:CASCADE;" json:"holes"`
	Strings           []String           `gorm:"constraint:OnDelete:CASCADE;" json:"strings"`
	Fluids            []Fluid            `gorm:"constraint:OnDelete:CASCADE;" json:"fluids"`
	PorePressures     []PorePressure     `gorm:"constraint:OnDelete:CASCADE;" json:"pore_pressures"`
	FractureGradients []FractureGradient `gorm:"constraint:OnDelete:CASCADE;" json:"fracture_gradients"`
	Rigs              []Rig              `gorm:"constraint:OnDelete:CASCADE;" json:"rigs"`
}

// Hole model
type Hole struct {
	ID                        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt                 time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt                 time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt                 gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	CaseID                    uuid.UUID      `gorm:"type:uuid;not null" json:"case_id"`
	OpenHoleMDTop             float64        `json:"open_hole_md_top"`
	OpenHoleMDBase            float64        `json:"open_hole_md_base"`
	OpenHoleLength            float64        `json:"open_hole_length"`
	OpenHoleVD                float64        `json:"open_hole_vd"`
	EffectiveDiameter         float64        `json:"effective_diameter"`
	FrictionFactorOpenHole    float64        `json:"friction_factor_open_hole"`
	LinearCapacityOpenHole    float64        `json:"linear_capacity_open_hole"`
	VolumeExcess              *float64       `json:"volume_excess,omitempty"`
	DescriptionOpenHole       *string        `json:"description_open_hole,omitempty"`
	TrippingInCasing          float64        `json:"tripping_in_casing"`
	TrippingOutCasing         float64        `json:"tripping_out_casing"`
	RotatingOnBottomCasing    float64        `json:"rotating_on_bottom_casing"`
	SlideDrillingCasing       float64        `json:"slide_drilling_casing"`
	BackReamingCasing         float64        `json:"back_reaming_casing"`
	RotatingOffBottomCasing   float64        `json:"rotating_off_bottom_casing"`
	TrippingInOpenHole        float64        `json:"tripping_in_open_hole"`
	TrippingOutOpenHole       float64        `json:"tripping_out_open_hole"`
	RotatingOnBottomOpenHole  float64        `json:"rotating_on_bottom_open_hole"`
	SlideDrillingOpenHole     float64        `json:"slide_drilling_open_hole"`
	BackReamingOpenHole       float64        `json:"back_reaming_open_hole"`
	RotatingOffBottomOpenHole float64        `json:"rotating_off_bottom_open_hole"`
	Caisings                  []Caising      `gorm:"constraint:OnDelete:CASCADE;" json:"caisings"`
}

// Caising model with UUID primary key and foreign key.
type Caising struct {
	ID                    uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt             time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt             time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt             gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	HoleID                uuid.UUID      `gorm:"type:uuid;not null" json:"hole_id"`
	MDTop                 float64        `json:"md_top"`
	MDBase                float64        `json:"md_base"`
	Length                float64        `json:"length"`
	ShoeMD                *float64       `json:"shoe_md,omitempty"`
	OD                    float64        `json:"od"`
	VD                    float64        `json:"vd"`
	DriftID               float64        `json:"drift_id"`
	EffectiveHoleDiameter float64        `json:"effective_hole_diameter"`
	Weight                float64        `json:"weight"`
	Grade                 string         `json:"grade"`
	MinYieldStrength      float64        `json:"min_yield_strength"`
	BurstRating           float64        `json:"burst_rating"`
	CollapseRating        float64        `json:"collapse_rating"`
	FrictionFactorCaising float64        `json:"friction_factor_caising"`
	LinearCapacityCaising float64        `json:"linear_capacity_caising"`
	DescriptionCaising    *string        `json:"description_caising,omitempty"`
	ManufacturerCaising   *string        `json:"manufacturer_caising,omitempty"`
	ModelCaising          *string        `json:"model_caising,omitempty"`
}

// String represents the Strings table.
type String struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CaseID    uuid.UUID      `gorm:"type:uuid;not null" json:"case_id"`
	Name      string         `gorm:"type:text;not null" json:"name"`
	Depth     float64        `gorm:"not null" json:"depth"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Sections  []Section      `gorm:"foreignKey:StringID;constraint:OnDelete:CASCADE;" json:"sections"`
}

// Section represents the Sections table, which is associated with a specific String.
type Section struct {
	ID                  uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	StringID            uuid.UUID      `gorm:"type:uuid;not null" json:"string_id"`
	Description         *string        `gorm:"type:text" json:"description,omitempty"`
	Manufacturer        *string        `gorm:"type:text" json:"manufacturer,omitempty"`
	Type                string         `gorm:"type:text;not null" json:"type"`
	BodyMD              float64        `gorm:"not null" json:"body_md"`
	BodyLength          float64        `gorm:"not null" json:"body_length"`
	BodyOD              float64        `gorm:"not null" json:"body_od"`
	BodyID              float64        `gorm:"not null" json:"body_id"`
	AvgJointLength      *float64       `json:"avg_joint_length,omitempty"`
	StabilizerLength    *float64       `json:"stabilizer_length,omitempty"`
	StabilizerOD        *float64       `json:"stabilizer_od,omitempty"`
	StabilizerID        *float64       `json:"stabilizer_id,omitempty"`
	Weight              *float64       `json:"weight,omitempty"`
	Material            *string        `gorm:"type:text" json:"material,omitempty"`
	Grade               *string        `gorm:"type:text" json:"grade,omitempty"`
	Class               *int           `json:"class,omitempty"`
	FrictionCoefficient *float64       `json:"friction_coefficient,omitempty"`
	MinYieldStrength    *float64       `json:"min_yield_strength,omitempty"`
	CreatedAt           time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt           time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// LibrarySection represents the LibrarySections table, which stores reusable sections.
type LibrarySection struct {
	ID                  uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Description         *string        `gorm:"type:text" json:"description,omitempty"`
	Manufacturer        *string        `gorm:"type:text" json:"manufacturer,omitempty"`
	Type                string         `gorm:"type:text;not null" json:"type"`
	BodyOD              float64        `gorm:"not null" json:"body_od"`
	BodyID              float64        `gorm:"not null" json:"body_id"`
	AvgJointLength      *float64       `json:"avg_joint_length,omitempty"`
	StabilizerLength    *float64       `json:"stabilizer_length,omitempty"`
	StabilizerOD        *float64       `json:"stabilizer_od,omitempty"`
	StabilizerID        *float64       `json:"stabilizer_id,omitempty"`
	Weight              *float64       `json:"weight,omitempty"`
	Material            *string        `gorm:"type:text" json:"material,omitempty"`
	Grade               *string        `gorm:"type:text" json:"grade,omitempty"`
	Class               *int           `json:"class,omitempty"`
	FrictionCoefficient *float64       `json:"friction_coefficient,omitempty"`
	MinYieldStrength    *float64       `json:"min_yield_strength,omitempty"`
	CreatedAt           time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt           time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// Fluid model with UUID primary key and foreign keys.
type Fluid struct {
	ID              uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CaseID          uuid.UUID      `gorm:"type:uuid;not null" json:"case_id"`
	Case            Case           `gorm:"foreignKey:CaseID;constraint:OnDelete:CASCADE;" json:"case"`
	Name            string         `gorm:"type:text;not null" json:"name"`
	Description     string         `gorm:"type:text" json:"description"`
	Density         float64        `gorm:"not null" json:"density"`
	FluidBaseTypeID uuid.UUID      `gorm:"type:uuid;not null" json:"fluid_base_type_id"`
	FluidBaseType   FluidType      `gorm:"foreignKey:FluidBaseTypeID;constraint:OnDelete:CASCADE;" json:"fluid_base_type"`
	BaseFluidID     uuid.UUID      `gorm:"type:uuid;not null" json:"base_fluid_id"`
	BaseFluid       FluidType      `gorm:"foreignKey:BaseFluidID;constraint:OnDelete:CASCADE;" json:"base_fluid"`
}

// FluidType model with UUID primary key.
type FluidType struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Name      string         `gorm:"type:text;not null" json:"name"`
}

// PorePressure model with UUID primary key and foreign key.
type PorePressure struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CaseID    uuid.UUID      `gorm:"type:uuid;not null" json:"case_id"`
	Case      Case           `gorm:"foreignKey:CaseID;constraint:OnDelete:CASCADE;" json:"-"`
	TVD       float64        `gorm:"not null" json:"tvd"`
	Pressure  float64        `gorm:"not null" json:"pressure"`
	EMW       float64        `gorm:"not null" json:"emw"`
}

// FractureGradient model with UUID primary key and foreign key.
type FractureGradient struct {
	ID                   uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt            time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt            time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt            gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CaseID               uuid.UUID      `gorm:"type:uuid;not null" json:"case_id"`
	Case                 Case           `gorm:"foreignKey:CaseID;constraint:OnDelete:CASCADE;" json:"-"`
	TemperatureAtSurface float64        `gorm:"not null" json:"temperature_at_surface"`
	TemperatureAtWellTVD float64        `gorm:"not null" json:"temperature_at_well_tvd"`
	TemperatureGradient  float64        `gorm:"not null" json:"temperature_gradient"`
	WellTVD              float64        `gorm:"not null" json:"well_tvd"`
}

// Rig model with UUID primary key and foreign key.
type Rig struct {
	ID                                uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt                         time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt                         time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt                         gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CaseID                            uuid.UUID      `gorm:"type:uuid;not null" json:"case_id"`
	Case                              Case           `gorm:"foreignKey:CaseID;constraint:OnDelete:CASCADE;" json:"-"`
	BlockRating                       *float64       `json:"block_rating,omitempty"`
	TorqueRating                      *float64       `json:"torque_rating,omitempty"`
	RatedWorkingPressure              float64        `gorm:"not null" json:"rated_working_pressure"`
	BopPressureRating                 float64        `gorm:"not null" json:"bop_pressure_rating"`
	SurfacePressureLoss               float64        `gorm:"not null;default:0" json:"surface_pressure_loss"`
	StandpipeLength                   *float64       `json:"standpipe_length,omitempty"`
	StandpipeInternalDiameter         *float64       `json:"standpipe_internal_diameter,omitempty"`
	HoseLength                        *float64       `json:"hose_length,omitempty"`
	HoseInternalDiameter              *float64       `json:"hose_internal_diameter,omitempty"`
	SwivelLength                      *float64       `json:"swivel_length,omitempty"`
	SwivelInternalDiameter            *float64       `json:"swivel_internal_diameter,omitempty"`
	KellyLength                       *float64       `json:"kelly_length,omitempty"`
	KellyInternalDiameter             *float64       `json:"kelly_internal_diameter,omitempty"`
	PumpDischargeLineLength           *float64       `json:"pump_discharge_line_length,omitempty"`
	PumpDischargeLineInternalDiameter *float64       `json:"pump_discharge_line_internal_diameter,omitempty"`
	TopDriveStackupLength             *float64       `json:"top_drive_stackup_length,omitempty"`
	TopDriveStackupInternalDiameter   *float64       `json:"top_drive_stackup_internal_diameter,omitempty"`
}
