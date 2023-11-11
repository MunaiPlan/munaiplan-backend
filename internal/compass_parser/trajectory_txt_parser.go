package compassparser

import (
	"strconv"
	"strings"
	"time"
)

// TODO() Correct all code in this file

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
	Inclination      float64
	Azimuth          float64
	SubSea           float64
	TVD              float64
	LocalNCoord      float64
	LocalECoord      float64
	GlobalNCoord     float64
	GlobalECoord     float64
	Dogleg           float64
	VerticalSection  float64
}

func parseHeader(headerLines []string) Header {
	header := Header{}
	for _, line := range headerLines {
		parts := strings.Split(line, "\t")
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "Customer":
			header.Customer = value
		case "Creation Date":
			date, _ := time.Parse("1/2/2006", value)
			header.CreationDate = date
		case "Project":
			header.Project = value
		case "Profile Type":
			header.ProfileType = value
		case "Field":
			header.Field = value
		case "Your Ref":
			header.YourRef = value
		case "Structure":
			header.Structure = value
		case "Job Number":
			header.JobNumber = value
		case "Wellhead":
			header.Wellhead = value
		case "Kelly Bushing Elev.":
			header.KellyBushingElev, _ = strconv.ParseFloat(value, 64)
		case "Profile":
			header.Profile = value
		case "Print Date":
			date, _ := time.Parse("1/2/2006", value)
			header.PrintDate = date
		}
	}
	return header
}

func parseData(dataLines []string) []Data {
	data := []Data{}
	for _, line := range dataLines {
		parts := strings.Fields(line)
		if len(parts) == 11 {
			md, _ := strconv.ParseFloat(parts[0], 64)
			inclination, _ := strconv.ParseFloat(parts[1], 64)
			azimuth, _ := strconv.ParseFloat(parts[2], 64)
			subSea, _ := strconv.ParseFloat(parts[3], 64)
			tvd, _ := strconv.ParseFloat(parts[4], 64)
			localNCoord, _ := strconv.ParseFloat(parts[5], 64)
			localECoord, _ := strconv.ParseFloat(parts[6], 64)
			globalNCoord, _ := strconv.ParseFloat(parts[7], 64)
			globalECoord, _ := strconv.ParseFloat(parts[8], 64)
			dogleg, _ := strconv.ParseFloat(parts[9], 64)
			verticalSection, _ := strconv.ParseFloat(parts[10], 64)

			dataPoint := Data{
				MD:              md,
				Inclination:     inclination,
				Azimuth:         azimuth,
				SubSea:          subSea,
				TVD:             tvd,
				LocalNCoord:     localNCoord,
				LocalECoord:     localECoord,
				GlobalNCoord:    globalNCoord,
				GlobalECoord:    globalECoord,
				Dogleg:          dogleg,
				VerticalSection: verticalSection,
			}

			data = append(data, dataPoint)
		}
	}
	return data
}

func trajectoryParser() 