package service

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/munaiplan/munaiplan-backend/internal/application/types/requests"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
)

// CaseDataError lists everything that must be fixed before a case can be predicted (HTTP 422).
type CaseDataError struct{ Problems []string }

func (e *CaseDataError) Error() string {
	return "кейс не готов к расчёту Torque & Drag: " + strings.Join(e.Problems, "; ")
}

// maxStations bounds one prediction request.
const maxStations = 5000

type frictionInterval struct {
	top, base, factor float64
}

// buildFeatures assembles the model's 20 per-station inputs. Rules:
//
//   - Stations are ordered by MD; exact duplicate depths are dropped.
//   - Only stations along the string (MD <= string bottom, i.e. the bit) are used: WellPlan's
//     depth-based T&D output, and so the model, only covers the string in the hole.
//   - A section's BodyMD is the MD of its bottom (WellPlan's string "Depth"), so a station at
//     MD belongs to the first section, ordered by BodyMD, with MD <= BodyMD.
//   - Stabilizer / tool-joint values that are absent mean "none" and are sent as 0.
//   - Friction comes from the hole section containing the station (casing, then open hole);
//     the string section's own coefficient is the fallback.
//   - Weight, average joint length and minimum yield are required; nothing else is invented.
func buildFeatures(units []*entities.TrajectoryUnit, sections []*entities.Section, holes []*entities.Hole) (*requests.TorqueAndDragFromMLModelRequest, error) {
	var problems []string
	stations := orderedStations(units)
	secs := orderedSections(sections)
	if len(stations) < 2 {
		problems = append(problems, "в траектории нужно минимум две точки инклинометрии")
	}
	if len(secs) == 0 {
		problems = append(problems, "в колонне кейса нет секций")
	}
	if len(problems) > 0 {
		return nil, &CaseDataError{problems}
	}
	problems = append(problems, layoutProblems(secs)...)
	friction := frictionIntervals(holes)
	bottom := secs[len(secs)-1].BodyMD
	req := &requests.TorqueAndDragFromMLModelRequest{}
	used := map[*entities.Section]bool{}
	for _, u := range stations {
		if u.MD > bottom {
			break
		}
		if len(req.MD) >= maxStations {
			problems = append(problems, fmt.Sprintf("вдоль колонны больше %d точек инклинометрии", maxStations))
			break
		}
		s := sectionAt(secs, u.MD)
		if !used[s] {
			used[s] = true
			problems = append(problems, missingInputs(s)...)
		}
		cof, ok := frictionAt(friction, u.MD)
		if !ok {
			if s.FrictionCoefficient == nil {
				problems = append(problems, fmt.Sprintf("нет коэффициента трения на глубине %.2f м: добавьте секции ствола или коэффициент трения секции колонны", u.MD))
				continue
			}
			cof = *s.FrictionCoefficient
		}
		if len(problems) == 0 {
			appendStation(req, u, s, cof)
		}
	}
	if len(req.MD) < 2 && len(problems) == 0 {
		problems = append(problems, fmt.Sprintf("вдоль колонны (0–%.2f м) меньше двух точек инклинометрии", bottom))
	}
	if len(problems) > 0 {
		return nil, &CaseDataError{dedupe(problems)}
	}
	return req, nil
}

// layoutTolerance absorbs WellPlan's rounding of component lengths (reports print whole metres).
const layoutTolerance = 1.0

// layoutProblems finds overlapping or disconnected components. Zero-length components such as
// a bit printed as "0 m" at the motor's depth are allowed.
func layoutProblems(secs []*entities.Section) []string {
	var problems []string
	for i, s := range secs {
		top := s.BodyMD - s.BodyLength
		prevBottom := 0.0
		if i > 0 {
			prevBottom = secs[i-1].BodyMD
		}
		switch {
		case s.BodyMD <= 0 || s.BodyLength < 0:
			problems = append(problems, fmt.Sprintf("секция %q: некорректная глубина или длина", s.Type))
		case top < prevBottom-layoutTolerance:
			problems = append(problems, fmt.Sprintf("секция %q (%.2f–%.2f м) перекрывает секцию выше", s.Type, top, s.BodyMD))
		case top > prevBottom+layoutTolerance:
			problems = append(problems, fmt.Sprintf("разрыв между %.2f м и %.2f м над секцией %q", prevBottom, top, s.Type))
		}
	}
	return problems
}

// missingInputs lists required model inputs that a section containing stations lacks.
func missingInputs(s *entities.Section) []string {
	var problems []string
	for _, f := range []struct {
		name  string
		value *float64
	}{{"вес", s.Weight}, {"средняя длина трубы", s.AvgJointLength}, {"предел текучести", s.MinYieldStrength}} {
		if f.value == nil {
			problems = append(problems, fmt.Sprintf("секция %q (низ %.2f м): не указано — %s", s.Type, s.BodyMD, f.name))
		}
	}
	return problems
}

func orderedStations(units []*entities.TrajectoryUnit) []*entities.TrajectoryUnit {
	out := make([]*entities.TrajectoryUnit, 0, len(units))
	for _, u := range units {
		if u != nil && !math.IsNaN(u.MD) {
			out = append(out, u)
		}
	}
	sort.SliceStable(out, func(i, j int) bool { return out[i].MD < out[j].MD })
	dedup := out[:0]
	for _, u := range out {
		if len(dedup) == 0 || u.MD != dedup[len(dedup)-1].MD {
			dedup = append(dedup, u)
		}
	}
	return dedup
}

func orderedSections(sections []*entities.Section) []*entities.Section {
	out := make([]*entities.Section, 0, len(sections))
	for _, s := range sections {
		if s != nil {
			out = append(out, s)
		}
	}
	// By bottom depth, then by top: a zero-length bit sits after the motor ending at the same depth.
	sort.SliceStable(out, func(i, j int) bool {
		if out[i].BodyMD != out[j].BodyMD {
			return out[i].BodyMD < out[j].BodyMD
		}
		return out[i].BodyMD-out[i].BodyLength < out[j].BodyMD-out[j].BodyLength
	})
	return out
}

// sectionAt returns the first section whose bottom is at or below md; callers ensure md <= bottom.
func sectionAt(secs []*entities.Section, md float64) *entities.Section {
	i := sort.Search(len(secs), func(i int) bool { return secs[i].BodyMD >= md })
	return secs[i]
}

// frictionIntervals lists casing intervals first so they win where they overlap open hole.
func frictionIntervals(holes []*entities.Hole) []frictionInterval {
	var casing, open []frictionInterval
	for _, h := range holes {
		if h == nil {
			continue
		}
		for _, c := range h.Caisings {
			if c != nil && c.MDBase > c.MDTop && c.FrictionFactorCaising > 0 {
				casing = append(casing, frictionInterval{c.MDTop, c.MDBase, c.FrictionFactorCaising})
			}
		}
		if h.OpenHoleMDBase > h.OpenHoleMDTop && h.FrictionFactorOpenHole > 0 {
			open = append(open, frictionInterval{h.OpenHoleMDTop, h.OpenHoleMDBase, h.FrictionFactorOpenHole})
		}
	}
	return append(casing, open...)
}

func frictionAt(intervals []frictionInterval, md float64) (float64, bool) {
	for _, f := range intervals {
		if md >= f.top && md <= f.base {
			return f.factor, true
		}
	}
	return 0, false
}

func appendStation(r *requests.TorqueAndDragFromMLModelRequest, u *entities.TrajectoryUnit, s *entities.Section, friction float64) {
	r.MD = append(r.MD, u.MD)
	r.Incl = append(r.Incl, u.Incl)
	r.Azim = append(r.Azim, u.Azim)
	r.SubSea = append(r.SubSea, u.SubSea)
	r.TVD = append(r.TVD, u.TVD)
	r.LocalNCoord = append(r.LocalNCoord, u.LocalNCoord)
	r.LocalECoord = append(r.LocalECoord, u.LocalECoord)
	r.GlobalNCoord = append(r.GlobalNCoord, u.GlobalNCoord)
	r.GlobalECoord = append(r.GlobalECoord, u.GlobalECoord)
	r.Dogleg = append(r.Dogleg, u.Dogleg)
	r.VerticalSection = append(r.VerticalSection, u.VerticalSection)
	r.BodyOD = append(r.BodyOD, s.BodyOD)
	r.BodyID = append(r.BodyID, s.BodyID)
	r.BodyAvgJointLength = append(r.BodyAvgJointLength, *s.AvgJointLength)
	r.StabilizerLength = append(r.StabilizerLength, zeroIfNil(s.StabilizerLength))
	r.StabilizerOD = append(r.StabilizerOD, zeroIfNil(s.StabilizerOD))
	r.StabilizerID = append(r.StabilizerID, zeroIfNil(s.StabilizerID))
	r.Weight = append(r.Weight, *s.Weight)
	r.CoefficientOfFriction = append(r.CoefficientOfFriction, friction)
	r.MinimumYieldStrength = append(r.MinimumYieldStrength, *s.MinYieldStrength)
}

func zeroIfNil(p *float64) float64 {
	if p == nil {
		return 0
	}
	return *p
}

func dedupe(items []string) []string {
	seen := map[string]bool{}
	out := items[:0]
	for _, s := range items {
		if !seen[s] {
			seen[s] = true
			out = append(out, s)
		}
	}
	sort.Strings(out)
	return out
}
