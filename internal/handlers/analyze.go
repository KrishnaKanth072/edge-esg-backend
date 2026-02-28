package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zebbank/edge-esg-backend/internal/dtos"
	"github.com/zebbank/edge-esg-backend/internal/error_codes"
	"github.com/zebbank/edge-esg-backend/internal/services"
	"github.com/zebbank/edge-esg-backend/internal/validator"
)

type AnalyzeHandler struct {
	orchestrator *services.Orchestrator
}

func NewAnalyzeHandler(orchestrator *services.Orchestrator) *AnalyzeHandler {
	return &AnalyzeHandler{orchestrator: orchestrator}
}

func (h *AnalyzeHandler) Analyze(c *gin.Context) {
	var req dtos.AnalyzeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse{
			Code:    string(error_codes.ESGInvalidInput),
			Message: "Invalid request body",
			Details: err.Error(),
		})
		return
	}

	if err := validator.ValidateStruct(&req); err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse{
			Code:    string(error_codes.ESGInvalidInput),
			Message: "Validation failed",
			Details: err.Error(),
		})
		return
	}

	role, _ := c.Get("user_role")
	req.UserRole = role.(string)

	response, err := h.orchestrator.Execute8LayerPipeline(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResponse{
			Code:    string(error_codes.ESGProcessingFailed),
			Message: "Failed to process ESG analysis",
			Details: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
