package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// TODO() Correct all tables in this file


// Куст (под месторождением)
type Site struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Area        string             `json:"area" bson:"area"`
	Block       string             `json:"block" bson:"block"`
	Azimuth     int                `json:"azimuth" bson:"azimuth"` // Manual entry
	LicenseArea string             `json:"licenseArea" bson:"licenseArea"`
	Region      string             `json:"region" bson:"region"`
	Country     string             `json:"country" bson:"country"`
	// Other fields...
}