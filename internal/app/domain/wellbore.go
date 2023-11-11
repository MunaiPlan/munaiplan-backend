package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// TODO() Correct all tables in this file
type Wellbore struct {
	ID                      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	WellID                  primitive.ObjectID `json:"wellId" bson:"wellId"`
	Name                    string             `json:"name" bson:"name"`
	BottomHoleLocation      int                `json:"bottomHoleLocation" bson:"bottomHoleLocation"`   // Coordinates (should be a more complex type)
	LateralFromExistingWell bool               `json:"lateralFromExistingWell" bson:"lateralFromExistingWell"` // Is this a lateral from an existing well
	MainBore                primitive.ObjectID `json:"mainBore" bson:"mainBore"`                       // If this wellbore is a lateral, reference to the main wellbore
	Type                    string             `json:"type" bson:"type"`                               // Type of the wellbore
	// Other fields...
}