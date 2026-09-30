package service

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/munaiplan/munaiplan-backend/internal/domain/entities"
	"github.com/munaiplan/munaiplan-backend/internal/domain/repository"
	"github.com/munaiplan/munaiplan-backend/internal/importers/wellplan"
)

// ErrAlreadyImported means the same source files were already imported into a live case.
type ErrAlreadyImported struct{ CaseID string }

func (e *ErrAlreadyImported) Error() string {
	return "these files were already imported (case " + e.CaseID + ")"
}

// ImportFile is one uploaded file.
type ImportFile struct {
	Name string
	Data []byte
}

// WellPlanUpload holds an engineering report (.docx), a survey export (.txt), or both.
type WellPlanUpload struct {
	Report *ImportFile
	Survey *ImportFile
}

// ImportPreview is what the user reviews before committing an import.
type ImportPreview struct {
	Report      *wellplan.Report       `json:"report"`
	Fingerprint string                 `json:"fingerprint"`
	Duplicate   *entities.ImportResult `json:"duplicate,omitempty"`
	Units       map[string]string      `json:"units"`
}

type importsService struct {
	repo repository.ImportsRepository
}

func NewImportsService(repo repository.ImportsRepository) *importsService {
	return &importsService{repo: repo}
}

func (s *importsService) PreviewWellPlan(ctx context.Context, organizationID string, upload WellPlanUpload) (*ImportPreview, error) {
	report, fingerprint, err := parseUpload(upload)
	if err != nil {
		return nil, err
	}
	duplicate, err := s.repo.FindByFingerprint(ctx, organizationID, fingerprint)
	if err != nil {
		return nil, err
	}
	return &ImportPreview{Report: report, Fingerprint: fingerprint, Duplicate: duplicate, Units: wellplan.Units}, nil
}

func (s *importsService) ImportWellPlan(ctx context.Context, organizationID, userID string, upload WellPlanUpload) (*entities.ImportResult, error) {
	report, fingerprint, err := parseUpload(upload)
	if err != nil {
		return nil, err
	}
	duplicate, err := s.repo.FindByFingerprint(ctx, organizationID, fingerprint)
	if err != nil {
		return nil, err
	}
	if duplicate != nil {
		return nil, &ErrAlreadyImported{CaseID: duplicate.CaseID}
	}
	return s.repo.SaveWellPlan(ctx, organizationID, userID, fingerprint, report)
}

func (s *importsService) CaseReference(ctx context.Context, organizationID, caseID string) (*wellplan.Report, error) {
	return s.repo.GetReport(ctx, organizationID, caseID)
}

// parseUpload parses and combines the files; parse failures become ValidationErrors.
func parseUpload(upload WellPlanUpload) (*wellplan.Report, string, error) {
	var report, survey *wellplan.Report
	var err error
	if upload.Report != nil {
		if report, err = wellplan.ParseReport(upload.Report.Name, upload.Report.Data); err != nil {
			return nil, "", asValidation(upload.Report.Name, err)
		}
	}
	if upload.Survey != nil {
		if survey, err = wellplan.ParseSurvey(upload.Survey.Name, upload.Survey.Data); err != nil {
			return nil, "", asValidation(upload.Survey.Name, err)
		}
	}
	combined, err := wellplan.Combine(report, survey)
	if err != nil {
		return nil, "", asValidation("", err)
	}
	return combined, fingerprintOf(combined.Source), nil
}

func asValidation(file string, err error) error {
	if errors.Is(err, wellplan.ErrInvalidFile) || errors.Is(err, wellplan.ErrNoSurvey) {
		if file != "" {
			return &ValidationError{fmt.Sprintf("%s: %v", file, err)}
		}
		return &ValidationError{err.Error()}
	}
	return err
}

// fingerprintOf identifies an import by the content of its files, independent of names and order.
func fingerprintOf(sources []wellplan.SourceFile) string {
	parts := make([]string, 0, len(sources))
	for _, f := range sources {
		parts = append(parts, f.Kind+":"+f.SHA256)
	}
	sort.Strings(parts)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(strings.Join(parts, "\n"))))
}
