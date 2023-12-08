package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Скважина (под кустом)
// TODO() Correct all tables in this file
type Well struct {
	ID                     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CompanyID              primitive.ObjectID `json:"companyId" bson:"companyId"`
	Name                   string             `json:"name" bson:"name"`
	LegalName              string             `json:"legalName" bson:"legalName"`
	Description            string             `json:"description" bson:"description"`
	LocationDescription    string             `json:"locationDescription" bson:"locationDescription"` // Need clarification
	UWI                    string             `json:"uwi" bson:"uwi"`                                 // Universal Well Identifier
	WellNumber             int                `json:"wellNumber" bson:"wellNumber"`
	GhostLevelAboveMSL     float64            `json:"ghostLevelAboveMSL" bson:"ghostLevelAboveMSL"`   // Mean Sea Level Rise
	ReferenceLevel         float64            `json:"referenceLevel" bson:"referenceLevel"`           // Altitude (inherits data)
	DefaultReferenceLevel  bool               `json:"defaultReferenceLevel" bson:"defaultReferenceLevel"` // Default reference level
	Elevation              float64            `json:"elevation" bson:"elevation"`                     // Elevation above the System reference level
	DrillName              string             `json:"drillName" bson:"drillName"`                     // Name of the Drill
	DateParameters         time.Time          `json:"dateParameters" bson:"dateParameters"`           // Current date for parameters
	Onshore                bool               `json:"onshore" bson:"onshore"`                         // Onshore well checkbox
	Offshore               bool               `json:"offshore" bson:"offshore"`                       // Offshore well checkbox
	WaterDepth             float64            `json:"waterDepth" bson:"waterDepth"`                   // For Offshore wells
	Subsea                 bool               `json:"subsea" bson:"subsea"`                           // For Subsea wells
	WellheadElevation      float64            `json:"wellheadElevation" bson:"wellheadElevation"`     // For Subsea wells
	// Other fields...
}
