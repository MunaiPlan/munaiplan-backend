package wellplan

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"unicode/utf8"

	"golang.org/x/text/encoding/charmap"
)

var surveyHeaderLabels = map[string][]string{
	"customer":     {"customer"},
	"project":      {"project"},
	"profile_type": {"profile type"},
	"field":        {"field"},
	"your_ref":     {"your ref"},
	"structure":    {"structure"},
	"job_number":   {"job number"},
	"wellhead":     {"wellhead"},
	"kb":           {"kelly bushing elev."},
	"profile":      {"profile"},
}

// ParseSurvey parses a WellPlan survey report exported as tab-separated text.
// Only metric exports (MD in metres) are accepted; nothing is converted.
func ParseSurvey(name string, data []byte) (*Report, error) {
	if len(data) > MaxUploadBytes {
		return nil, fmt.Errorf("%w: file exceeds %d MB", ErrInvalidFile, MaxUploadBytes>>20)
	}
	text := string(data)
	if !utf8.Valid(data) {
		// WellPlan on Russian Windows writes CP1251.
		decoded, err := charmap.Windows1251.NewDecoder().Bytes(data)
		if err != nil {
			return nil, fmt.Errorf("%w: unknown text encoding", ErrInvalidFile)
		}
		text = string(decoded)
	}
	lines := strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n")

	r := newReport()
	r.Source = append(r.Source, sourceFile(name, "survey", data))
	headerAt := -1
	for i, line := range lines {
		cells := splitTabs(line)
		if len(cells) > 3 && norm(cells[0]) == "md" {
			headerAt = i
			break
		}
		for j := 0; j+1 < len(cells); j += 2 {
			r.setSurveyHeader(cells[j], cells[j+1])
		}
	}
	if headerAt < 0 || headerAt+1 >= len(lines) {
		return nil, fmt.Errorf("%w: survey column header (MD, Incl., Azim., ...) not found", ErrInvalidFile)
	}
	header := splitTabs(lines[headerAt])
	idx := columnIndex(header, surveyColumns)
	for _, key := range []string{"md", "inc", "azi", "tvd"} {
		if _, ok := idx[key]; !ok {
			return nil, fmt.Errorf("%w: survey column %s is missing", ErrInvalidFile, key)
		}
	}
	units := splitTabs(lines[headerAt+1])
	if md := idx["md"]; md < len(units) && norm(units[md]) != "(m)" {
		return nil, fmt.Errorf("%w: survey depth unit %q is not metres; export the survey in metric units", ErrInvalidFile, units[md])
	}
	rows := make([][]string, 0, len(lines))
	for _, line := range lines[headerAt+2:] {
		rows = append(rows, splitTabs(line))
	}
	r.parseSurveyRows(rows, idx)
	if len(r.Survey) < 2 {
		return nil, fmt.Errorf("%w: survey contains fewer than two stations", ErrInvalidFile)
	}
	return r, nil
}

func (r *Report) setSurveyHeader(k, v string) {
	h := &r.SurveyInfo
	switch key, _ := label(surveyHeaderLabels, k); key {
	case "customer":
		h.Customer = v
	case "project":
		h.Project = v
	case "profile_type":
		h.ProfileType = v
	case "field":
		h.Field = v
	case "your_ref":
		h.YourRef = v
	case "structure":
		h.Structure = v
	case "job_number":
		h.JobNumber = v
	case "wellhead":
		h.Wellhead = v
	case "kb":
		h.KellyBushingElev = numPtr(v)
	case "profile":
		h.Profile = v
	}
}

func splitTabs(line string) []string {
	cells := strings.Split(line, "\t")
	for i := range cells {
		cells[i] = clean(cells[i])
	}
	for len(cells) > 0 && cells[len(cells)-1] == "" {
		cells = cells[:len(cells)-1]
	}
	return cells
}

func sourceFile(name, kind string, data []byte) SourceFile {
	return SourceFile{Name: name, Kind: kind, SHA256: fmt.Sprintf("%x", sha256.Sum256(data)), Bytes: int64(len(data))}
}
