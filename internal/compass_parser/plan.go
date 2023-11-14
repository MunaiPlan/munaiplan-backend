package compassparser

import "time"

type Header struct {
	Customer         string
	CreationDate     time.Time
	Project          string
	ProfileType      string
	Field            string
	YourRef          string
	Structure        string
	JobNumber        string
	Wellhead         string
	KellyBushingElev float64
	Profile          string
	PrintDate        time.Time
}

type Data struct {
	MD               float64
	Incl             float64
	Azim             float64
	SubSea           float64
	TVD              float64
	LocalNCoord      float64
	LocalECoord      float64
	GlobalNCoord     float64
	GlobalECoord     float64
	Dogleg           float64
	VerticalSection  float64
}

type Plan struct {
	Header Header
	Data             []Data
}