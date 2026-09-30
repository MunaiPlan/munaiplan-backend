package wellplan

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	thousands = regexp.MustCompile(`^-?\d{1,3}(,\d{3})+(\.\d+)?$`)
	leadNum   = regexp.MustCompile(`^-?\d+(?:[.,]\d+)?`)
	atNumber  = regexp.MustCompile(`@\s*(-?\d+(?:[.,]\d+)?)`)
	// Units are often a separate paragraph in the cell: "Глубина по стволу" + "(m)".
	spaceParen = regexp.MustCompile(`\s+\(`)
)

// parseNumber reads WellPlan's printed numbers: "1,854", "55,000", "2296.02", "0,25", "-24".
// Units or text after the number are ignored ("30.00 m", "6.8046 atm").
func parseNumber(s string) (float64, bool) {
	s = strings.TrimSpace(strings.ReplaceAll(s, " ", ""))
	if s == "" || strings.EqualFold(s, "NA") {
		return 0, false
	}
	if field := strings.Fields(s); len(field) > 0 {
		s = field[0]
	}
	switch {
	case thousands.MatchString(s):
		s = strings.ReplaceAll(s, ",", "")
	default:
		m := leadNum.FindString(s)
		if m == "" {
			return 0, false
		}
		s = strings.Replace(m, ",", ".", 1)
	}
	v, err := strconv.ParseFloat(s, 64)
	return v, err == nil
}

func numPtr(s string) *float64 {
	if v, ok := parseNumber(s); ok {
		return &v
	}
	return nil
}

// norm lower-cases and collapses a label so RU/EN variants compare reliably.
func norm(s string) string {
	s = spaceParen.ReplaceAllString(strings.ToLower(clean(s)), "(")
	return strings.NewReplacer("ё", "е", ":", "").Replace(s)
}

// label maps every known RU/EN spelling to one canonical key.
func label(dict map[string][]string, text string) (string, bool) {
	n := norm(text)
	for key, variants := range dict {
		for _, v := range variants {
			if n == v {
				return key, true
			}
		}
	}
	return "", false
}

var caseLabels = map[string][]string{
	"company":          {"компания", "company"},
	"field":            {"месторождение", "project", "field"},
	"site":             {"кустовая площадка", "site"},
	"well":             {"скважина", "well"},
	"wellbore":         {"ствол", "wellbore"},
	"design":           {"траектория", "design"},
	"case":             {"кейс", "case"},
	"md":               {"забой по стволу", "hole md"},
	"tvd":              {"забой по вертикали", "hole tvd"},
	"air_gap":          {"воздушный зазор", "air gap"},
	"ground_elevation": {"ground elevation", "высота над уровнем моря"},
	"datum":            {"начало отсчета", "reference point"},
	"well_type":        {"тип скважины", "well type"},
}

var fluidLabels = map[string][]string{
	"name":           {"жидкость", "fluid"},
	"type":           {"тип", "type"},
	"base":           {"основа", "mud base type"},
	"base_fluid":     {"жидкость основы", "base fluid"},
	"rheology_model": {"реологическая модель", "rheology model"},
}

var tdSettingLabels = map[string][]string{
	"bit_depth":    {"глубина долота по стволу", "measured depth of bit"},
	"block_weight": {"вес блока", "block weight"},
	"flow_rate":    {"подача насосов", "pump rate"},
}

var geoLabels = map[string][]string{
	"ambient":  {"температура окружающей среды", "ambient temperature"},
	"at_depth": {"температура @ глубина", "temperature @ depth"},
	"gradient": {"градиент", "gradient"},
}

var operationLabels = map[string][]string{
	OpTrippingIn:        {"спуск", "tripping in"},
	OpTrippingOut:       {"подъём", "подъем", "tripping out"},
	OpRotatingOnBottom:  {"вращение на забое", "бурение ротором", "rotating on bottom"},
	OpSlideDrilling:     {"бурение гзд", "направленное бурение", "slide drilling"},
	OpRotatingOffBottom: {"вращение над забоем", "rotating off bottom"},
	OpBackReaming:       {"обратная проработка", "backreaming", "back reaming"},
}

// surveyColumns maps survey header cells (RU report, EN report, survey .txt) to station fields.
var surveyColumns = map[string][]string{
	"md":     {"глубина по стволу(m)", "md(m)", "md"},
	"inc":    {"зенитный угол(°)", "inc(°)", "incl."},
	"azi":    {"азимутальный угол(°)", "az(°)", "azim."},
	"tvd":    {"глубина по вертикали(m)", "tvd(m)", "tvd"},
	"dls":    {"пространст. интенсивн. искривления(°/30m)", "dls(°/30m)", "dogleg"},
	"vs":     {"смещ.-е(m)", "vsect(m)", "vertical section"},
	"ns":     {"ns(m)", "local n coord"},
	"ew":     {"ew(m)", "local e coord"},
	"subsea": {"sub-sea"},
	"gn":     {"global n coord"},
	"ge":     {"global e coord"},
}

// startsWith reports whether the first cell of a row begins with any of the prefixes.
func startsWith(cell string, prefixes ...string) bool {
	n := norm(cell)
	for _, p := range prefixes {
		if strings.HasPrefix(n, p) {
			return true
		}
	}
	return false
}
