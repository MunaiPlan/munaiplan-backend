package postgres

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/importers/wellplan"
	"github.com/munaiplan/munaiplan-backend/internal/infrastructure/drivers/postgres/models"
	"gorm.io/gorm"
)

const formatWellPlan = "wellplan"

type importsRepository struct {
	db *gorm.DB
}

func NewImportsRepository(db *gorm.DB) *importsRepository {
	return &importsRepository{db: db}
}

func (r *importsRepository) FindByFingerprint(ctx context.Context, organizationID, fingerprint string) (*entities.ImportResult, error) {
	var row struct {
		ID        uuid.UUID
		CaseID    uuid.UUID
		CreatedAt time.Time
	}
	err := r.db.WithContext(ctx).Raw(`
		SELECT ci.id, ci.case_id, ci.created_at FROM case_imports ci
		JOIN cases c ON c.id = ci.case_id AND c.deleted_at IS NULL
		WHERE ci.organization_id = ? AND ci.fingerprint = ?`, organizationID, fingerprint).Scan(&row).Error
	if err != nil || row.ID == uuid.Nil {
		return nil, err
	}
	return &entities.ImportResult{ImportID: row.ID.String(), CaseID: row.CaseID.String(), CreatedAt: row.CreatedAt}, nil
}

func (r *importsRepository) GetReport(ctx context.Context, organizationID, caseID string) (*wellplan.Report, error) {
	if _, err := uuid.Parse(caseID); err != nil {
		return nil, nil
	}
	var imp models.CaseImport
	err := r.db.WithContext(ctx).Where("organization_id = ? AND case_id = ?", organizationID, caseID).
		Order("created_at DESC").First(&imp).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var report wellplan.Report
	if err := json.Unmarshal([]byte(imp.Report), &report); err != nil {
		return nil, fmt.Errorf("stored import report is unreadable: %w", err)
	}
	return &report, nil
}

func (r *importsRepository) SaveWellPlan(ctx context.Context, organizationID, userID, fingerprint string, rep *wellplan.Report) (*entities.ImportResult, error) {
	orgID, err := uuid.Parse(organizationID)
	if err != nil {
		return nil, fmt.Errorf("invalid organization id")
	}
	payload, err := json.Marshal(rep)
	if err != nil {
		return nil, err
	}
	res := &entities.ImportResult{Created: map[string]bool{}}
	err = r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Drop provenance of a previous import whose case has since been deleted.
		if err := tx.Exec(`DELETE FROM case_imports ci USING cases c
			WHERE ci.case_id = c.id AND c.deleted_at IS NOT NULL AND ci.organization_id = ? AND ci.fingerprint = ?`,
			orgID, fingerprint).Error; err != nil {
			return err
		}
		h := hierarchy{tx: tx, res: res}
		c := rep.Case
		company := h.company(orgID, c.Company)
		field := h.field(company, c.Field)
		site := h.site(field, c.Site)
		well := h.well(site, c.Well)
		wellbore := h.wellbore(well, c.Wellbore, c.MD)
		design := h.design(wellbore, c.Design)
		if h.err != nil {
			return h.err
		}

		trajectory := buildTrajectory(design, rep)
		if err := tx.Create(&trajectory).Error; err != nil {
			return fmt.Errorf("create trajectory: %w", err)
		}
		kase := buildCase(trajectory.ID, rep)
		if err := tx.Create(&kase).Error; err != nil {
			return fmt.Errorf("create case: %w", err)
		}
		if err := createCaseChildren(tx, kase.ID, rep); err != nil {
			return err
		}
		imp := models.CaseImport{OrganizationID: orgID, CaseID: kase.ID, Format: formatWellPlan, Fingerprint: fingerprint, Report: string(payload)}
		if uid, err := uuid.Parse(userID); err == nil {
			imp.CreatedBy = &uid
		}
		if err := tx.Create(&imp).Error; err != nil {
			return mapWriteError(err)
		}
		res.CompanyID, res.FieldID, res.SiteID = company.String(), field.String(), site.String()
		res.WellID, res.WellboreID, res.DesignID = well.String(), wellbore.String(), design.String()
		res.TrajectoryID, res.CaseID, res.ImportID, res.CreatedAt = trajectory.ID.String(), kase.ID.String(), imp.ID.String(), imp.CreatedAt
		res.Created["trajectory"], res.Created["case"] = true, true
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

// hierarchy finds each level by name under its parent or creates it. The first error sticks.
type hierarchy struct {
	tx  *gorm.DB
	res *entities.ImportResult
	err error
}

// findOrCreate returns the oldest live row matching where, or inserts build().
func findOrCreate[T any](h *hierarchy, level string, build func() *T, id func(*T) uuid.UUID, where string, args ...any) uuid.UUID {
	if h.err != nil {
		return uuid.Nil
	}
	var existing T
	err := h.tx.Where(where, args...).Order("created_at").First(&existing).Error
	if err == nil {
		h.res.Created[level] = false
		return id(&existing)
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		h.err = err
		return uuid.Nil
	}
	record := build()
	if err := h.tx.Create(record).Error; err != nil {
		h.err = fmt.Errorf("create %s: %w", level, mapWriteError(err))
		return uuid.Nil
	}
	h.res.Created[level] = true
	return id(record)
}

func (h *hierarchy) company(org uuid.UUID, name string) uuid.UUID {
	name = limit(name)
	return findOrCreate(h, "company", func() *models.Company {
		return &models.Company{OrganizationID: org, Name: name, Division: "Imported from WellPlan"}
	}, func(m *models.Company) uuid.UUID { return m.ID }, "organization_id = ? AND name = ?", org, name)
}

func (h *hierarchy) field(company uuid.UUID, name string) uuid.UUID {
	name = limit(name)
	return findOrCreate(h, "field", func() *models.Field { return &models.Field{CompanyID: company, Name: name} },
		func(m *models.Field) uuid.UUID { return m.ID }, "company_id = ? AND name = ?", company, name)
}

func (h *hierarchy) site(field uuid.UUID, name string) uuid.UUID {
	name = limit(name)
	return findOrCreate(h, "site", func() *models.Site { return &models.Site{FieldID: field, Name: name} },
		func(m *models.Site) uuid.UUID { return m.ID }, "field_id = ? AND name = ?", field, name)
}

func (h *hierarchy) well(site uuid.UUID, name string) uuid.UUID {
	name = limit(name)
	return findOrCreate(h, "well", func() *models.Well { return &models.Well{SiteID: site, Name: name, WellNumber: name} },
		func(m *models.Well) uuid.UUID { return m.ID }, "site_id = ? AND name = ?", site, name)
}

func (h *hierarchy) wellbore(well uuid.UUID, name string, md *float64) uuid.UUID {
	name = limit(name)
	return findOrCreate(h, "wellbore", func() *models.Wellbore {
		w := &models.Wellbore{WellID: well, Name: name}
		if md != nil {
			w.WellboreDepth = *md
		}
		return w
	}, func(m *models.Wellbore) uuid.UUID { return m.ID }, "well_id = ? AND name = ?", well, name)
}

func (h *hierarchy) design(wellbore uuid.UUID, name string) uuid.UUID {
	name = limit(name)
	return findOrCreate(h, "design", func() *models.Design {
		return &models.Design{WellboreID: wellbore, PlanName: name, Stage: "Imported", Version: "1", ActualDate: time.Now().UTC()}
	}, func(m *models.Design) uuid.UUID { return m.ID }, "wellbore_id = ? AND plan_name = ?", wellbore, name)
}

func buildTrajectory(design uuid.UUID, rep *wellplan.Report) models.Trajectory {
	files := make([]string, 0, len(rep.Source))
	for _, f := range rep.Source {
		files = append(files, f.Name)
	}
	t := models.Trajectory{DesignID: design, Name: limit(rep.Case.Design),
		Description: "Imported from WellPlan: " + strings.Join(files, ", ")}
	sh := rep.SurveyInfo
	header := models.TrajectoryHeader{
		Customer: firstNonBlank(sh.Customer, rep.Case.Company), Project: sh.Project, ProfileType: sh.ProfileType,
		Field: firstNonBlank(sh.Field, rep.Case.Field), YourRef: sh.YourRef, Structure: firstNonBlank(sh.Structure, rep.Case.Site),
		JobNumber: sh.JobNumber, Wellhead: firstNonBlank(sh.Wellhead, rep.Case.Well), Profile: firstNonBlank(sh.Profile, rep.Case.Design),
	}
	if sh.KellyBushingElev != nil {
		header.KellyBushingElev = *sh.KellyBushingElev
	} else if rep.Case.DatumElevation != nil {
		header.KellyBushingElev = *rep.Case.DatumElevation
	}
	t.Headers = []models.TrajectoryHeader{header}
	for _, s := range rep.Survey {
		t.Units = append(t.Units, models.TrajectoryUnit{MD: s.MD, Incl: s.Inc, Azim: s.Azi, TVD: s.TVD,
			SubSea: val(s.SubSea), LocalNCoord: val(s.NS), LocalECoord: val(s.EW), GlobalNCoord: val(s.GlobalN),
			GlobalECoord: val(s.GlobalE), Dogleg: val(s.DLS), VerticalSection: val(s.VerticalSection)})
	}
	return t
}

func buildCase(trajectory uuid.UUID, rep *wellplan.Report) models.Case {
	c := models.Case{TrajectoryID: trajectory, CaseName: limit(rep.Case.Case),
		CaseDescription: "Imported from WellPlan", IsComplete: len(rep.String) > 0 && len(rep.Survey) >= 2}
	switch {
	case rep.TorqueDrag != nil && rep.TorqueDrag.BitDepth != nil:
		c.DrillDepth = *rep.TorqueDrag.BitDepth
	case rep.Case.MD != nil:
		c.DrillDepth = *rep.Case.MD
	case len(rep.String) > 0:
		c.DrillDepth = rep.String[len(rep.String)-1].Depth
	}
	if len(rep.String) > 0 {
		c.PipeSize = rep.String[0].BodyOD
	}
	return c
}

func createCaseChildren(tx *gorm.DB, caseID uuid.UUID, rep *wellplan.Report) error {
	if len(rep.String) > 0 {
		str := models.String{CaseID: caseID, Name: "WellPlan work string", Depth: rep.String[len(rep.String)-1].Depth}
		for _, c := range rep.String {
			sec := models.Section{Type: c.Type, BodyMD: c.Depth, BodyLength: c.Length, BodyOD: c.BodyOD, BodyID: val(c.BodyID),
				AvgJointLength: c.AvgJointLength, StabilizerLength: c.JointLength, StabilizerOD: c.JointOD, StabilizerID: c.JointID,
				Weight: c.Weight, Material: strPtr(c.Material), Grade: strPtr(c.Grade), MinYieldStrength: c.MinYield}
			if class, err := strconv.Atoi(c.Class); err == nil {
				sec.Class = &class
			}
			str.Sections = append(str.Sections, sec)
		}
		if err := tx.Create(&str).Error; err != nil {
			return fmt.Errorf("create string: %w", err)
		}
	}
	if hole := buildHole(caseID, rep); hole != nil {
		if err := tx.Create(hole).Error; err != nil {
			return fmt.Errorf("create hole sections: %w", err)
		}
	}
	if rep.Fluid != nil && rep.Fluid.Density != nil {
		if err := createFluid(tx, caseID, rep.Fluid); err != nil {
			return err
		}
	}
	if g := rep.Geothermal; g != nil && g.AmbientTemperature != nil && g.TemperatureAtDepth != nil && g.Gradient != nil && g.Depth != nil {
		// The legacy table name is fracture_gradients, but its columns are the geothermal profile.
		geo := models.FractureGradient{CaseID: caseID, TemperatureAtSurface: *g.AmbientTemperature,
			TemperatureAtWellTVD: *g.TemperatureAtDepth, TemperatureGradient: *g.Gradient, WellTVD: *g.Depth}
		if err := tx.Omit("Case").Create(&geo).Error; err != nil {
			return fmt.Errorf("create geothermal profile: %w", err)
		}
	}
	return nil
}

// buildHole maps casing sections to caisings and the deepest open-hole section to the hole row.
// WellPlan reports one friction factor per section type, applied to every operation.
func buildHole(caseID uuid.UUID, rep *wellplan.Report) *models.Hole {
	if len(rep.Holes) == 0 {
		return nil
	}
	hole := &models.Hole{CaseID: caseID}
	for _, s := range rep.Holes {
		top := s.Depth - s.Length
		friction := val(s.FrictionFactor)
		if s.IsOpenHole() {
			hole.OpenHoleMDTop, hole.OpenHoleMDBase, hole.OpenHoleLength = top, s.Depth, s.Length
			hole.OpenHoleVD = tvdAt(rep.Survey, s.Depth)
			hole.EffectiveDiameter = firstValue(s.EffectiveDiameter, &s.InnerDiameter)
			hole.FrictionFactorOpenHole, hole.LinearCapacityOpenHole, hole.VolumeExcess = friction, val(s.LinearCapacity), s.VolumeExcess
			hole.TrippingInOpenHole, hole.TrippingOutOpenHole, hole.RotatingOnBottomOpenHole = friction, friction, friction
			hole.SlideDrillingOpenHole, hole.BackReamingOpenHole, hole.RotatingOffBottomOpenHole = friction, friction, friction
			continue
		}
		id := s.InnerDiameter
		desc := s.Type
		hole.Caisings = append(hole.Caisings, models.Caising{MDTop: top, MDBase: s.Depth, Length: s.Length, ShoeMD: s.ShoeDepth,
			VD: tvdAt(rep.Survey, s.Depth), InnerDiameter: &id, DriftID: val(s.Drift), EffectiveHoleDiameter: val(s.EffectiveDiameter),
			FrictionFactorCaising: friction, LinearCapacityCaising: val(s.LinearCapacity), DescriptionCaising: &desc})
		hole.TrippingInCasing, hole.TrippingOutCasing, hole.RotatingOnBottomCasing = friction, friction, friction
		hole.SlideDrillingCasing, hole.BackReamingCasing, hole.RotatingOffBottomCasing = friction, friction, friction
	}
	return hole
}

func createFluid(tx *gorm.DB, caseID uuid.UUID, f *wellplan.Fluid) error {
	baseType, err := fluidType(tx, firstNonBlank(f.Base, "Unknown"))
	if err != nil {
		return err
	}
	baseFluid, err := fluidType(tx, firstNonBlank(f.BaseFluid, f.Base, "Unknown"))
	if err != nil {
		return err
	}
	desc := strings.TrimSpace(fmt.Sprintf("%s %s; PV %s cP; YP %s lbf/100ft²", f.Type, f.RheologyModel, fmtPtr(f.PlasticViscosity), fmtPtr(f.YieldPoint)))
	fluid := models.Fluid{CaseID: caseID, Name: firstNonBlank(f.Name, "Imported fluid"), Description: desc,
		Density: *f.Density, FluidBaseTypeID: baseType, BaseFluidID: baseFluid}
	if err := tx.Omit("Case", "FluidBaseType", "BaseFluid").Create(&fluid).Error; err != nil {
		return fmt.Errorf("create fluid: %w", err)
	}
	return nil
}

// fluidType returns the shared catalogue entry for a base-fluid name, creating it once.
func fluidType(tx *gorm.DB, name string) (uuid.UUID, error) {
	var t models.FluidType
	err := tx.Where("name = ?", name).Order("created_at").First(&t).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t = models.FluidType{Name: name}
		err = tx.Create(&t).Error
	}
	return t.ID, err
}

// tvdAt interpolates TVD linearly between survey stations; outside the survey it clamps.
func tvdAt(survey []wellplan.Station, md float64) float64 {
	if len(survey) == 0 {
		return 0
	}
	if md <= survey[0].MD {
		return survey[0].TVD
	}
	for i := 1; i < len(survey); i++ {
		a, b := survey[i-1], survey[i]
		if md <= b.MD {
			return a.TVD + (b.TVD-a.TVD)*(md-a.MD)/(b.MD-a.MD)
		}
	}
	return survey[len(survey)-1].TVD
}

func val(p *float64) float64 {
	if p == nil {
		return 0
	}
	return *p
}

func firstValue(ps ...*float64) float64 {
	for _, p := range ps {
		if p != nil {
			return *p
		}
	}
	return 0
}

func strPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func firstNonBlank(values ...string) string {
	for _, v := range values {
		if strings.Trim(v, "? .-") != "" {
			return v
		}
	}
	return ""
}

func fmtPtr(p *float64) string {
	if p == nil {
		return "n/a"
	}
	return strconv.FormatFloat(*p, 'f', -1, 64)
}

// limit keeps names within the varchar(255) columns.
func limit(name string) string {
	runes := []rune(strings.TrimSpace(name))
	if len(runes) > 255 {
		runes = runes[:255]
	}
	return string(runes)
}
