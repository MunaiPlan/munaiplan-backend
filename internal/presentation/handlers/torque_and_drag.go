package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/munaiplan/munaiplan-backend/pkg/values"
)

func (h *Handler) initTorqueAndDragRoutes(api *gin.RouterGroup) {
	torqueAndDrag := api.Group("/torque-and-drag", h.authMiddleware.UserIdentity)
	{
		torqueAndDrag.POST("/effective-tension", h.predictHandler(func(ctx context.Context, caseID string) (any, error) {
			return h.services.TorqueAndDrag.CalculateEffectiveTensionFromMLModel(ctx, caseID)
		}))
		torqueAndDrag.POST("/weight-on-bit", h.predictHandler(func(ctx context.Context, caseID string) (any, error) {
			return h.services.TorqueAndDrag.CalculateWeightOnBitFromMlModel(ctx, caseID)
		}))
		torqueAndDrag.POST("/surface-torque", h.predictHandler(func(ctx context.Context, caseID string) (any, error) {
			return h.services.TorqueAndDrag.CalculateSurfaceTorqueFromMlModel(ctx, caseID)
		}))
		torqueAndDrag.POST("/min-weight", h.predictHandler(func(ctx context.Context, caseID string) (any, error) {
			return h.services.TorqueAndDrag.CalculateMinWeightFromMLModel(ctx, caseID)
		}))
		// Model output beside WellPlan's own results for an imported case; 404 without reference data.
		torqueAndDrag.GET("/comparison", func(c *gin.Context) {
			organizationID := c.GetString(values.OrganizationIdCtx)
			h.predictHandler(func(ctx context.Context, caseID string) (any, error) {
				return h.services.TorqueAndDrag.CompareWithReference(ctx, organizationID, caseID)
			})(c)
		})
	}
}

// predictHandler runs one model family for ?caseId=. Case ownership is enforced by UserIdentity.
// Responses: 200 series by depth, 422 case data incomplete, 503 model unavailable, 502 invalid model output.
func (h *Handler) predictHandler(run func(ctx context.Context, caseID string) (any, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		caseID, err := h.validateQueryIDParam(c, values.CaseIdQueryParam)
		if err != nil {
			return // validateQueryIDParam already responded with 400
		}
		result, err := run(c.Request.Context(), caseID)
		if err != nil {
			h.respondError(c, err)
			return
		}
		c.JSON(http.StatusOK, result)
	}
}
