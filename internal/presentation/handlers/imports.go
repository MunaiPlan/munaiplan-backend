package handlers

import (
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/munaiplan/munaiplan-backend/internal/application/service"
	"github.com/munaiplan/munaiplan-backend/internal/helpers"
	"github.com/munaiplan/munaiplan-backend/internal/importers/wellplan"
	"github.com/munaiplan/munaiplan-backend/pkg/values"
)

// Two files plus multipart overhead.
const maxImportBody = 2*wellplan.MaxUploadBytes + (1 << 20)

func (h *Handler) initImportRoutes(api *gin.RouterGroup) {
	imports := api.Group("/imports", h.authMiddleware.UserIdentity)
	{
		imports.POST("/wellplan/preview", h.previewWellPlan)
		imports.POST("/wellplan", h.importWellPlan)
		imports.GET("/cases/:id/reference", h.caseReference)
	}
}

// @Summary Preview a WellPlan import
// @Description Parses a WellPlan report (.docx, field "report") and/or survey export (.txt, field "survey") without saving.
// @Tags imports
// @Accept multipart/form-data
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} service.ImportPreview
// @Failure 400,413 {object} helpers.Response
// @Router /api/v1/imports/wellplan/preview [post]
func (h *Handler) previewWellPlan(c *gin.Context) {
	upload, ok := readWellPlanUpload(c)
	if !ok {
		return
	}
	preview, err := h.services.Imports.PreviewWellPlan(c.Request.Context(), c.GetString(values.OrganizationIdCtx), upload)
	if err != nil {
		h.respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, preview)
}

// @Summary Import a WellPlan case
// @Description Creates or reuses company/field/site/well/wellbore/design by name and always creates a new trajectory and case.
// @Tags imports
// @Accept multipart/form-data
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 201 {object} entities.ImportResult
// @Failure 400,409,413 {object} helpers.Response
// @Router /api/v1/imports/wellplan [post]
func (h *Handler) importWellPlan(c *gin.Context) {
	upload, ok := readWellPlanUpload(c)
	if !ok {
		return
	}
	result, err := h.services.Imports.ImportWellPlan(c.Request.Context(), c.GetString(values.OrganizationIdCtx), c.GetString(values.UserIdCtx), upload)
	if err != nil {
		h.respondError(c, err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

// @Summary WellPlan reference data for a case
// @Description Returns the stored WellPlan report (inputs and WellPlan's own results) for an imported case.
// @Tags imports
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Case ID"
// @Success 200 {object} wellplan.Report
// @Failure 404 {object} helpers.Response
// @Router /api/v1/imports/cases/{id}/reference [get]
func (h *Handler) caseReference(c *gin.Context) {
	report, err := h.services.Imports.CaseReference(c.Request.Context(), c.GetString(values.OrganizationIdCtx), c.Param("id"))
	if err != nil {
		h.respondError(c, err)
		return
	}
	if report == nil {
		helpers.NewErrorResponse(c, http.StatusNotFound, "no imported WellPlan data for this case")
		return
	}
	c.JSON(http.StatusOK, report)
}

// readWellPlanUpload reads the optional "report" (.docx) and "survey" (.txt) parts.
func readWellPlanUpload(c *gin.Context) (service.WellPlanUpload, bool) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxImportBody)
	var upload service.WellPlanUpload
	if err := c.Request.ParseMultipartForm(maxImportBody); err != nil {
		var tooLarge *http.MaxBytesError
		if errors.As(err, &tooLarge) {
			helpers.NewErrorResponse(c, http.StatusRequestEntityTooLarge, "upload is too large")
		} else {
			helpers.NewErrorResponse(c, http.StatusBadRequest, "expected a multipart form with a report and/or survey file")
		}
		return upload, false
	}
	defer c.Request.MultipartForm.RemoveAll()
	var err error
	if upload.Report, err = formFile(c.Request.MultipartForm, "report", ".docx"); err == nil {
		upload.Survey, err = formFile(c.Request.MultipartForm, "survey", ".txt")
	}
	if err != nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return upload, false
	}
	if upload.Report == nil && upload.Survey == nil {
		helpers.NewErrorResponse(c, http.StatusBadRequest, "attach a WellPlan report (.docx) and/or survey export (.txt)")
		return upload, false
	}
	return upload, true
}

func formFile(form *multipart.Form, field, ext string) (*service.ImportFile, error) {
	headers := form.File[field]
	if len(headers) == 0 {
		return nil, nil
	}
	if len(headers) > 1 {
		return nil, errors.New("only one " + field + " file is allowed")
	}
	header := headers[0]
	name := filepath.Base(header.Filename)
	if !strings.EqualFold(filepath.Ext(name), ext) {
		return nil, errors.New(field + " must be a " + ext + " file")
	}
	if header.Size > wellplan.MaxUploadBytes {
		return nil, errors.New(field + " file is too large")
	}
	f, err := header.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	data, err := io.ReadAll(io.LimitReader(f, wellplan.MaxUploadBytes+1))
	if err != nil {
		return nil, err
	}
	return &service.ImportFile{Name: name, Data: data}, nil
}
