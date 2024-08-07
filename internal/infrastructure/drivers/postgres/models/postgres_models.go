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
	Organization   Organization `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE;" json:"organization"`
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
	Organization   Organization `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE;" json:"organization"`
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
	Company         Company   `gorm:"foreignKey:CompanyID;constraint:OnDelete:CASCADE;" json:"company"`
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
	Field   Field     `gorm:"foreignKey:FieldID;constraint:OnDelete:CASCADE;" json:"field"`
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
	Site                    Site       `gorm:"foreignKey:SiteID;constraint:OnDelete:CASCADE;" json:"site"`
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
	Well                           Well      `gorm:"foreignKey:WellID;constraint:OnDelete:CASCADE;" json:"well"`
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
	Wellbore     Wellbore     `gorm:"foreignKey:WellboreID;constraint:OnDelete:CASCADE;" json:"wellbore"`
	Cases        []Case       `gorm:"constraint:OnDelete:CASCADE;" json:"cases"`
	Trajectories []Trajectory `gorm:"constraint:OnDelete:CASCADE;" json:"trajectories"`
}

// Case model with UUID primary key and foreign key.
type Case struct {
	BaseModel
	CaseName        string    `json:"case_name"`
	CaseDescription string    `json:"case_description"`
	DrillDepth      float64   `json:"drill_depth"`
	PipeSize        float64   `json:"pipe_size"`
	DesignID        uuid.UUID `gorm:"type:uuid;not null" json:"design_id"`
	Design          Design    `gorm:"foreignKey:DesignID;constraint:OnDelete:CASCADE;" json:"design"`
}

// Trajectory model with UUID primary key and foreign key.
type Trajectory struct {
	BaseModel
	Name        string             `gorm:"type:varchar(255)" json:"name"`
	Description string             `json:"description"`
	DesignID    uuid.UUID          `gorm:"type:uuid;not null" json:"design_id"`
	Design      Design             `gorm:"foreignKey:DesignID;constraint:OnDelete:CASCADE;" json:"design"`
	Headers     []TrajectoryHeader `gorm:"constraint:OnDelete:CASCADE;" json:"headers"`
	Units       []TrajectoryUnit   `gorm:"constraint:OnDelete:CASCADE;" json:"units"`
}

// TrajectoryHeader model with UUID primary key and foreign key.
type TrajectoryHeader struct {
	ID               uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	TrajectoryID     uuid.UUID  `gorm:"type:uuid;not null" json:"trajectory_id"`
	Customer         string     `json:"customer"`
	Project          string     `json:"project"`
	ProfileType      string     `json:"profile_type"`
	Field            string     `json:"field"`
	YourRef          string     `json:"your_ref"`
	Structure        string     `json:"structure"`
	JobNumber        string     `json:"job_number"`
	Wellhead         string     `json:"wellhead"`
	KellyBushingElev float64    `json:"kelly_bushing_elev"`
	Profile          string     `json:"profile"`
	CreatedAt        time.Time  `gorm:"not null" json:"created_at"`
	Trajectory       Trajectory `gorm:"foreignKey:TrajectoryID;constraint:OnDelete:CASCADE;" json:"trajectory"`
}

// TrajectoryUnit model with UUID primary key and foreign key.
type TrajectoryUnit struct {
	ID              uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	TrajectoryID    uuid.UUID  `gorm:"type:uuid;not null" json:"trajectory_id"`
	MD              float64    `json:"md"`
	Incl            float64    `json:"incl"`
	Azim            float64    `json:"azim"`
	SubSea          float64    `json:"sub_sea"`
	TVD             float64    `json:"tvd"`
	LocalNCoord     float64    `json:"local_n_coord"`
	LocalECoord     float64    `json:"local_e_coord"`
	GlobalNCoord    float64    `json:"global_n_coord"`
	GlobalECoord    float64    `json:"global_e_coord"`
	Dogleg          float64    `json:"dogleg"`
	VerticalSection float64    `json:"vertical_section"`
	CreatedAt       time.Time  `gorm:"not null" json:"created_at"`
	Trajectory      Trajectory `gorm:"foreignKey:TrajectoryID;constraint:OnDelete:CASCADE;" json:"trajectory"`
}
