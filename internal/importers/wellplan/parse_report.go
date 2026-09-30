package wellplan

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type tableKind int

const (
	kindUnknown tableKind = iota
	kindCase
	kindFluid
	kindRheology
	kindHoles
	kindFriction
	kindString
	kindGrades
	kindSurvey
	kindTDSettings
	kindOperations
	kindLimits
	kindLoads
	kindBitHydraulics
	kindGeothermal
)

// ParseReport parses a WellPlan engineering report (.docx, Russian or English).
func ParseReport(name string, data []byte) (*Report, error) {
	blocks, err := readDocx(data)
	if err != nil {
		return nil, err
	}
	r := newReport()
	r.Source = append(r.Source, sourceFile(name, "report", data))
	r.Language = "ru"

	var (
		previous  tableKind
		surveyIdx map[string]int
	)
	for _, b := range blocks {
		if !b.isTable() || len(b.rows) == 0 {
			continue
		}
		kind := classify(b.rows[0])
		if kind == kindUnknown && rowEmpty(b.rows[0]) {
			kind = previous // WellPlan splits long tables into a header table and a body table.
		}
		switch kind {
		case kindCase:
			r.parseCase(b.rows)
		case kindFluid:
			r.parseFluid(b.rows)
		case kindRheology:
			r.parseRheology(b.rows)
		case kindHoles:
			r.parseHoles(b.rows)
		case kindFriction:
			r.parseFriction(b.rows)
		case kindString:
			r.parseString(b.rows)
		case kindGrades:
			r.parseGrades(b.rows)
		case kindSurvey:
			if surveyIdx == nil {
				surveyIdx = columnIndex(b.rows[0], surveyColumns)
			}
			r.parseSurveyRows(b.rows, surveyIdx)
		case kindTDSettings:
			r.parseTDSettings(b.rows)
		case kindOperations:
			r.parseOperations(b.rows)
		case kindLimits:
			r.parseLimits(b.rows)
		case kindLoads:
			r.parseLoads(b.rows)
		case kindBitHydraulics:
			r.parseKV(b.rows, r.Hydraulics)
		case kindGeothermal:
			r.parseGeothermal(b.rows)
		}
		previous = kind // an unrelated table ends any continuation run
		if kind == kindCase && norm(b.rows[0][0]) == "company" {
			r.Language = "en"
		}
	}
	r.resolveGrades()
	r.fillStringDepths()
	if r.Case.Company == "" && len(r.String) == 0 {
		return nil, fmt.Errorf("%w: no case information or string table found", ErrInvalidFile)
	}
	return r, nil
}

func newReport() *Report {
	return &Report{
		Grades: map[string]float64{}, Friction: map[string]float64{},
		Hydraulics: map[string]string{}, Warnings: []string{},
		Survey: []Station{}, Holes: []HoleSection{}, String: []Component{},
	}
}

func classify(row []string) tableKind {
	if len(row) == 0 {
		return kindUnknown
	}
	first := norm(row[0])
	second := ""
	if len(row) > 1 {
		second = norm(row[1])
	}
	joined := norm(strings.Join(row, " "))
	switch {
	case first == "компания" || first == "company":
		return kindCase
	case (first == "жидкость" || first == "fluid") && len(row) >= 4:
		return kindFluid
	case strings.HasPrefix(first, "температура(") || strings.HasPrefix(first, "temperature("):
		return kindRheology
	case (first == "тип секции" || first == "sectiontype" || first == "section type") && len(row) >= 8:
		return kindHoles
	case (first == "тип секции" || first == "section type") && len(row) == 2:
		return kindFriction
	case (first == "тип" || first == "type") && (strings.HasPrefix(second, "длина") || strings.HasPrefix(second, "length")):
		return kindString
	case (first == "марка стали" || first == "grade") && strings.Contains(joined, "psi"):
		return kindGrades
	case containsLabel(surveyColumns["md"], first) && len(row) > 1 && containsLabel(surveyColumns["inc"], second):
		return kindSurvey
	case containsLabel(tdSettingLabels["bit_depth"], first):
		return kindTDSettings
	case strings.HasPrefix(first, "операции при бурении") || first == "drilling":
		return kindOperations
	case startsWith(row[0], "запас по затяжке", "overpull margin"):
		return kindLimits
	case startsWith(row[0], "условие нагружения", "load condition"):
		return kindLoads
	case containsLabel(tdSettingLabels["flow_rate"], first) && (strings.Contains(joined, "давление на стояке") || strings.Contains(joined, "stand pipe pressure")):
		return kindBitHydraulics
	case containsLabel(geoLabels["ambient"], first):
		return kindGeothermal
	}
	return kindUnknown
}

func containsLabel(variants []string, n string) bool {
	for _, v := range variants {
		if v == n {
			return true
		}
	}
	return false
}

func rowEmpty(row []string) bool {
	for _, c := range row {
		if c != "" {
			return false
		}
	}
	return true
}

// pairs yields key/value cells laid out as "key | value | key | value".
func pairs(rows [][]string, fn func(key, value string)) {
	for _, row := range rows {
		for i := 0; i+1 < len(row); i += 2 {
			if row[i] != "" {
				fn(row[i], row[i+1])
			}
		}
	}
}

func (r *Report) parseCase(rows [][]string) {
	c := &r.Case
	pairs(rows, func(k, v string) {
		key, ok := label(caseLabels, k)
		if !ok {
			return
		}
		switch key {
		case "company":
			c.Company = v
		case "field":
			c.Field = v
		case "site":
			c.Site = v
		case "well":
			c.Well = v
		case "wellbore":
			c.Wellbore = v
		case "design":
			c.Design = v
		case "case":
			c.Case = v
		case "md":
			c.MD = numPtr(v)
		case "tvd":
			c.TVD = numPtr(v)
		case "air_gap":
			c.AirGap = numPtr(v)
		case "ground_elevation":
			c.GroundElevation = numPtr(v)
		case "datum":
			c.Datum = v
			if m := atNumber.FindStringSubmatch(v); m != nil {
				c.DatumElevation = numPtr(m[1])
			}
		case "well_type":
			c.WellType = v
		}
	})
}

func (r *Report) parseFluid(rows [][]string) {
	if r.Fluid == nil {
		r.Fluid = &Fluid{}
	}
	pairs(rows, func(k, v string) {
		switch key, _ := label(fluidLabels, k); key {
		case "name":
			r.Fluid.Name = v
		case "type":
			r.Fluid.Type = v
		case "base":
			r.Fluid.Base = v
		case "base_fluid":
			r.Fluid.BaseFluid = v
		case "rheology_model":
			r.Fluid.RheologyModel = v
		}
	})
}

// parseRheology reads the first data row: temperature, pressure, density, flag, PV, YP.
func (r *Report) parseRheology(rows [][]string) {
	if r.Fluid == nil {
		r.Fluid = &Fluid{}
	}
	for _, row := range rows[1:] {
		if len(row) < 6 {
			continue
		}
		if _, ok := parseNumber(row[0]); !ok {
			continue
		}
		f := r.Fluid
		f.Temperature, f.Pressure, f.Density = numPtr(row[0]), numPtr(row[1]), numPtr(row[2])
		f.PlasticViscosity, f.YieldPoint = numPtr(row[4]), numPtr(row[5])
		return
	}
}

func (r *Report) parseHoles(rows [][]string) {
	for _, row := range rows {
		if len(row) < 8 || classify(row) == kindHoles || row[0] == "" {
			continue
		}
		depth, ok1 := parseNumber(row[1])
		length, ok2 := parseNumber(row[2])
		id, ok3 := parseNumber(row[4])
		if !ok1 || !ok2 || !ok3 {
			r.warn("hole section %q skipped: depth, length or ID is not numeric", row[0])
			continue
		}
		h := HoleSection{Type: row[0], Depth: depth, Length: length, ShoeDepth: numPtr(row[3]), InnerDiameter: id,
			Drift: numPtr(row[5]), EffectiveDiameter: numPtr(row[6]), FrictionFactor: numPtr(row[7])}
		if len(row) > 8 {
			h.LinearCapacity = numPtr(row[8])
		}
		if len(row) > 9 {
			h.VolumeExcess = numPtr(row[9])
		}
		r.Holes = append(r.Holes, h)
	}
}

func (r *Report) parseFriction(rows [][]string) {
	for _, row := range rows[1:] {
		if len(row) == 2 {
			if v, ok := parseNumber(row[1]); ok {
				r.Friction[row[0]] = v
			}
		}
	}
}

// parseString reads component rows: type, length, depth, body OD, ID, avg joint length,
// tool joint length, OD, ID, weight, material, grade, class.
func (r *Report) parseString(rows [][]string) {
	for _, row := range rows {
		if len(row) < 10 || row[0] == "" || classify(row) == kindString {
			continue
		}
		length, ok1 := parseNumber(row[1])
		od, ok3 := parseNumber(row[3])
		if !ok1 || !ok3 {
			continue // sub-header rows
		}
		depth, ok2 := parseNumber(row[2])
		if !ok2 {
			depth = math.NaN() // filled from cumulative lengths in fillStringDepths
		}
		c := Component{Type: row[0], Length: length, Depth: depth, BodyOD: od, BodyID: numPtr(row[4]),
			AvgJointLength: numPtr(row[5]), JointLength: numPtr(row[6]), JointOD: numPtr(row[7]),
			JointID: numPtr(row[8]), Weight: numPtr(row[9])}
		if len(row) > 10 {
			c.Material = row[10]
		}
		if len(row) > 11 {
			c.Grade = row[11]
		}
		if len(row) > 12 {
			c.Class = row[12]
		}
		r.String = append(r.String, c)
	}
}

// fillStringDepths derives missing bottom depths as the running sum of lengths from surface.
func (r *Report) fillStringDepths() {
	for i := range r.String {
		c := &r.String[i]
		if !math.IsNaN(c.Depth) {
			continue
		}
		c.Depth = c.Length
		if i > 0 {
			c.Depth += r.String[i-1].Depth
		}
		r.warn("string depths were missing and are derived from cumulative component lengths")
	}
}

func (r *Report) parseGrades(rows [][]string) {
	for _, row := range rows[1:] {
		if len(row) >= 2 && row[0] != "" {
			if v, ok := parseNumber(row[1]); ok {
				r.Grades[row[0]] = v
			}
		}
	}
}

func (r *Report) resolveGrades() {
	for i := range r.String {
		c := &r.String[i]
		if c.Grade == "" {
			continue
		}
		if v, ok := r.Grades[c.Grade]; ok {
			c.MinYield = &v
		} else {
			r.warn("grade %q of %s has no minimum yield in the grade table", c.Grade, c.Type)
		}
	}
}

// columnIndex finds each known column in a header row.
func columnIndex(header []string, dict map[string][]string) map[string]int {
	idx := map[string]int{}
	for i, cell := range header {
		if key, ok := label(dict, cell); ok {
			if _, seen := idx[key]; !seen {
				idx[key] = i
			}
		}
	}
	return idx
}

func (r *Report) parseSurveyRows(rows [][]string, idx map[string]int) {
	get := func(row []string, key string) *float64 {
		i, ok := idx[key]
		if !ok || i >= len(row) {
			return nil
		}
		return numPtr(row[i])
	}
	for _, row := range rows {
		md, inc, azi, tvd := get(row, "md"), get(row, "inc"), get(row, "azi"), get(row, "tvd")
		if md == nil || inc == nil || azi == nil || tvd == nil {
			continue // header or blank rows
		}
		r.Survey = append(r.Survey, Station{MD: *md, Inc: *inc, Azi: *azi, TVD: *tvd,
			SubSea: get(row, "subsea"), NS: get(row, "ns"), EW: get(row, "ew"),
			GlobalN: get(row, "gn"), GlobalE: get(row, "ge"),
			DLS: get(row, "dls"), VerticalSection: get(row, "vs")})
	}
}

func (r *Report) td() *TorqueDragResults {
	if r.TorqueDrag == nil {
		r.TorqueDrag = &TorqueDragResults{WOB: map[string]float64{}, BitTorque: map[string]float64{},
			TripSpeed: map[string]float64{}, Limits: map[string]string{}, Loads: []LoadCondition{}}
	}
	return r.TorqueDrag
}

func (r *Report) parseTDSettings(rows [][]string) {
	td := r.td()
	pairs(rows, func(k, v string) {
		switch key, _ := label(tdSettingLabels, k); key {
		case "bit_depth":
			td.BitDepth = numPtr(v)
		case "block_weight":
			td.BlockWeight = numPtr(v)
		case "flow_rate":
			td.FlowRate = numPtr(v)
		}
	})
}

func (r *Report) parseOperations(rows [][]string) {
	td := r.td()
	tripping := false
	for _, row := range rows {
		if len(row) < 2 {
			continue
		}
		if startsWith(row[0], "операции при спо", "tripping") && !startsWith(row[0], "tripping in", "tripping out") {
			tripping = true
			continue
		}
		op, ok := label(operationLabels, row[0])
		if !ok {
			continue
		}
		if tripping {
			if v, ok := parseNumber(row[1]); ok {
				td.TripSpeed[op] = v
			}
			continue
		}
		if v, ok := parseNumber(row[1]); ok {
			td.WOB[op] = v
		}
		if len(row) > 2 {
			if v, ok := parseNumber(row[2]); ok {
				td.BitTorque[op] = v
			}
		}
	}
}

func (r *Report) parseLimits(rows [][]string) {
	td := r.td()
	for _, row := range rows {
		if len(row) >= 2 && row[0] != "" {
			td.Limits[row[0]] = strings.TrimSpace(strings.Join(row[1:], " "))
		}
	}
}

// parseLoads reads the load summary. The last nine cells are numeric results; the cells
// between the label and them are failure flags.
func (r *Report) parseLoads(rows [][]string) {
	td := r.td()
	for _, row := range rows {
		if len(row) < 10 {
			continue
		}
		op, ok := label(operationLabels, row[0])
		if !ok {
			continue
		}
		n := len(row)
		lc := LoadCondition{Operation: op, Label: row[0],
			SurfaceTorque: numPtr(row[n-9]), WindupWithBit: numPtr(row[n-8]), WindupWithoutBit: numPtr(row[n-7]),
			HookLoad: numPtr(row[n-6]), Stretch: numPtr(row[n-5]),
			NeutralFromSurface: numPtr(row[n-2]), NeutralFromBit: numPtr(row[n-1])}
		for _, flag := range row[1 : n-9] {
			if flag != "" {
				lc.Failures = append(lc.Failures, flag)
			}
		}
		td.Loads = append(td.Loads, lc)
	}
}

func (r *Report) parseKV(rows [][]string, into map[string]string) {
	pairs(rows, func(k, v string) {
		if v != "" {
			into[k] = v
		}
	})
}

func (r *Report) parseGeothermal(rows [][]string) {
	g := &Geothermal{}
	pairs(rows, func(k, v string) {
		switch key, _ := label(geoLabels, k); key {
		case "ambient":
			g.AmbientTemperature = numPtr(v)
		case "at_depth":
			parts := strings.SplitN(v, "@", 2)
			g.TemperatureAtDepth = numPtr(parts[0])
			if len(parts) == 2 {
				g.Depth = numPtr(parts[1])
			}
		case "gradient":
			g.Gradient = numPtr(v)
		}
	})
	r.Geothermal = g
}

func (r *Report) warn(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	for _, w := range r.Warnings {
		if w == msg {
			return
		}
	}
	r.Warnings = append(r.Warnings, msg)
}

// sortedKeys is used for deterministic output in tests and previews.
func sortedKeys(m map[string]float64) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
