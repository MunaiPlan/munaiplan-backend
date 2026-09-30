package wellplan

import (
	"errors"
	"fmt"
	"math"
	"path/filepath"
	"sort"
	"strings"
)

// ErrNoSurvey means neither the report nor an attached survey file contains stations.
var ErrNoSurvey = errors.New("no survey stations: attach the survey export (.txt) for this case")

// Combine merges an optional report and an optional survey export into one validated case.
// An attached survey replaces the report's survey table because it carries sub-sea and
// global coordinates that the report omits.
func Combine(report, survey *Report) (*Report, error) {
	if report == nil && survey == nil {
		return nil, fmt.Errorf("%w: no file supplied", ErrInvalidFile)
	}
	r := report
	if r == nil {
		r = survey
		r.caseFromSurveyHeader()
	} else if survey != nil {
		if len(r.Survey) > 0 {
			r.warn("the attached survey file replaces the report's own survey table (%d stations)", len(r.Survey))
		}
		r.Survey, r.SurveyInfo = survey.Survey, survey.SurveyInfo
		r.Source = append(r.Source, survey.Source...)
		r.Warnings = append(r.Warnings, survey.Warnings...)
	}
	if err := r.finalize(); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *Report) caseFromSurveyHeader() {
	h := r.SurveyInfo
	r.Case.Company, r.Case.Field, r.Case.Site = h.Customer, h.Field, h.Structure
	r.Case.Well, r.Case.Design = h.Wellhead, h.Profile
}

func (r *Report) finalize() error {
	r.normaliseSurvey()
	if len(r.Survey) < 2 {
		return ErrNoSurvey
	}
	r.deriveSubSea()
	r.defaultNames()
	r.checkString()
	r.applyFrictionFactors()
	if len(r.Holes) == 0 {
		r.warn("no hole sections: casing/open-hole intervals and friction factors are unknown")
	}
	if r.Fluid == nil || r.Fluid.Density == nil {
		r.warn("no fluid density in the report")
	}
	return nil
}

// normaliseSurvey sorts stations by MD and removes exact duplicate depths.
func (r *Report) normaliseSurvey() {
	if !sort.SliceIsSorted(r.Survey, func(i, j int) bool { return r.Survey[i].MD < r.Survey[j].MD }) {
		sort.SliceStable(r.Survey, func(i, j int) bool { return r.Survey[i].MD < r.Survey[j].MD })
		r.warn("survey stations were not in MD order and have been sorted")
	}
	out := r.Survey[:0]
	for i, s := range r.Survey {
		if i > 0 && s.MD == out[len(out)-1].MD {
			r.warn("duplicate survey station at MD %.2f m removed", s.MD)
			continue
		}
		out = append(out, s)
	}
	r.Survey = out
	for _, s := range r.Survey {
		if s.Inc < 0 || s.Inc > 180 || s.Azi < 0 || s.Azi >= 360.0001 || math.IsNaN(s.MD) {
			r.warn("survey station at MD %.2f m has an out-of-range angle (inc %.2f, azi %.2f)", s.MD, s.Inc, s.Azi)
		}
	}
}

// deriveSubSea fills sub-sea TVD from the datum elevation when the source lacks it
// (sub-sea = TVD - datum elevation above sea level; WellPlan prints datums as "@ -19m").
func (r *Report) deriveSubSea() {
	missingSubSea, missingGlobal := false, false
	for _, s := range r.Survey {
		missingSubSea = missingSubSea || s.SubSea == nil
		missingGlobal = missingGlobal || s.GlobalN == nil || s.GlobalE == nil
	}
	if missingSubSea {
		elevation := r.Case.DatumElevation
		if elevation == nil {
			elevation = r.SurveyInfo.KellyBushingElev
		}
		if elevation != nil {
			for i := range r.Survey {
				if r.Survey[i].SubSea == nil {
					v := r.Survey[i].TVD - *elevation
					r.Survey[i].SubSea = &v
				}
			}
			r.warn("sub-sea depth derived as TVD − datum elevation (%.2f m)", *elevation)
		} else {
			r.warn("sub-sea depth unknown: no datum elevation; stored as 0")
		}
	}
	if missingGlobal {
		r.warn("global (grid) coordinates are not in the source; stored as 0 — attach the survey .txt export to include them")
	}
}

func (r *Report) defaultNames() {
	base := "Imported case"
	if len(r.Source) > 0 {
		base = strings.TrimSuffix(filepath.Base(r.Source[0].Name), filepath.Ext(r.Source[0].Name))
	}
	fill := func(value *string, fallback, what string) {
		if isBlankName(*value) {
			*value = fallback
			r.warn("%s name missing in the source; using %q", what, fallback)
		}
	}
	c := &r.Case
	fill(&c.Company, "Импортированные кейсы", "company")
	fill(&c.Field, "Imported field", "field")
	fill(&c.Site, c.Field, "site")
	fill(&c.Well, base, "well")
	fill(&c.Wellbore, c.Well, "wellbore")
	fill(&c.Design, "Design #1", "design")
	fill(&c.Case, base, "case")
}

// isBlankName treats names lost to encoding ("????") as missing.
func isBlankName(s string) bool {
	return strings.Trim(s, "? .-") == ""
}

func (r *Report) checkString() {
	if len(r.String) == 0 {
		r.warn("no work string: the case cannot be used for Torque & Drag until a string is added")
		return
	}
	for i := 1; i < len(r.String); i++ {
		if r.String[i].Depth < r.String[i-1].Depth {
			r.warn("string component %d (%s) is shallower than the one above it", i+1, r.String[i].Type)
		}
	}
	bottom := r.String[len(r.String)-1].Depth
	if r.TorqueDrag != nil && r.TorqueDrag.BitDepth != nil && math.Abs(bottom-*r.TorqueDrag.BitDepth) > 1 {
		r.warn("string bottom %.2f m differs from the analysed bit depth %.2f m", bottom, *r.TorqueDrag.BitDepth)
	}
	last := r.Survey[len(r.Survey)-1].MD
	if bottom > last+1 {
		r.warn("string reaches %.2f m, deeper than the last survey station (%.2f m)", bottom, last)
	}
}

// applyFrictionFactors fills hole-section friction from the friction-factor table.
func (r *Report) applyFrictionFactors() {
	for i := range r.Holes {
		h := &r.Holes[i]
		if h.FrictionFactor == nil {
			if v, ok := r.Friction[h.Type]; ok {
				h.FrictionFactor = &v
			}
		}
	}
}
