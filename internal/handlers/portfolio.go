package handlers

import (
	"net/http"

	"github.com/edgeesg/edge-esg-backend/internal/dtos"
	"github.com/edgeesg/edge-esg-backend/internal/error_codes"
	"github.com/edgeesg/edge-esg-backend/internal/services"
	"github.com/edgeesg/edge-esg-backend/internal/validator"
	"github.com/gin-gonic/gin"
)

type PortfolioHandler struct {
	orchestrator *services.Orchestrator
}

func NewPortfolioHandler(orchestrator *services.Orchestrator) *PortfolioHandler {
	return &PortfolioHandler{
		orchestrator: orchestrator,
	}
}

// ComparePortfolio handles portfolio comparison requests
func (h *PortfolioHandler) ComparePortfolio(c *gin.Context) {
	var req dtos.PortfolioCompareRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse{
			Code:    string(error_codes.ESGInvalidInput),
			Message: "Invalid request format",
			Details: err.Error(),
		})
		return
	}

	// Validate request
	if err := validator.ValidateStruct(&req); err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse{
			Code:    string(error_codes.ESGInvalidInput),
			Message: "Validation failed",
			Details: err.Error(),
		})
		return
	}

	// Validate company count
	if len(req.Companies) < 2 {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse{
			Code:    string(error_codes.ESGInvalidInput),
			Message: "At least 2 companies required for comparison",
		})
		return
	}

	if len(req.Companies) > 10 {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse{
			Code:    string(error_codes.ESGInvalidInput),
			Message: "Maximum 10 companies allowed for comparison",
		})
		return
	}

	// Execute portfolio comparison
	response, err := h.orchestrator.ComparePortfolio(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResponse{
			Code:    string(error_codes.ESGProcessingFailed),
			Message: "Failed to compare portfolio",
			Details: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
