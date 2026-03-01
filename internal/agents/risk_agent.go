package agents

import (
	"context"
	"fmt"
)

type RiskAgent struct {
}

func NewRiskAgent() *RiskAgent {
	return &RiskAgent{}
}

type RiskAssessmentRequest struct {
	CompanyName     string
	ESGScore        float64
	NewsSentiment   float64
	StockVolatility float64
}

type RiskAssessmentResponse struct {
	Action    string // APPROVE/REJECT/REVIEW
	Reasons   []string
	RiskScore float64 // 0-100
	RiskLevel string  // LOW/MEDIUM/HIGH/CRITICAL
}

// AssessRisk performs comprehensive risk analysis
func (r *RiskAgent) AssessRisk(ctx context.Context, req *RiskAssessmentRequest) (*RiskAssessmentResponse, error) {
	response := &RiskAssessmentResponse{
		Reasons: []string{},
	}

	// Calculate risk score (0-100, lower is better)
	// Start from 0 and add risk factors
	riskScore := 0.0

	// ESG Score Impact (0-10 scale, higher ESG = lower risk)
	// Low ESG adds risk: (10 - ESG) × 5 = 0-50 points
	esgRisk := (10.0 - req.ESGScore) * 5.0
	riskScore += esgRisk

	// News Sentiment Impact (0-1 scale, higher sentiment = lower risk)
	// Negative sentiment adds risk: (1 - sentiment) × 40 = 0-40 points
	sentimentRisk := (1.0 - req.NewsSentiment) * 40.0
	riskScore += sentimentRisk

	// Stock Volatility Impact (if provided)
	// High volatility adds risk: volatility × 30 = 0-30 points
	if req.StockVolatility > 0 {
		volatilityRisk := req.StockVolatility * 30.0
		riskScore += volatilityRisk
	}

	// Clamp to 0-100
	if riskScore < 0 {
		riskScore = 0
	}
	if riskScore > 100 {
		riskScore = 100
	}

	response.RiskScore = riskScore

	// Determine risk level and action
	if riskScore >= 70 {
		response.RiskLevel = "CRITICAL"
		response.Action = "REJECT"
		response.Reasons = append(response.Reasons, "Critical risk level detected")
	} else if riskScore >= 50 {
		response.RiskLevel = "HIGH"
		response.Action = "REJECT"
		response.Reasons = append(response.Reasons, "High risk level")
	} else if riskScore >= 30 {
		response.RiskLevel = "MEDIUM"
		response.Action = "REVIEW"
		response.Reasons = append(response.Reasons, "Moderate risk requires review")
	} else {
		response.RiskLevel = "LOW"
		response.Action = "APPROVE"
	}

	// Add specific reasons based on factors
	if req.ESGScore < 4.0 {
		response.Reasons = append(response.Reasons,
			fmt.Sprintf("Low ESG score (%.1f/10) indicates sustainability risks", req.ESGScore))
	} else if req.ESGScore >= 7.0 {
		response.Reasons = append(response.Reasons,
			fmt.Sprintf("Strong ESG score (%.1f/10)", req.ESGScore))
	} else {
		response.Reasons = append(response.Reasons,
			fmt.Sprintf("Moderate ESG score (%.1f/10)", req.ESGScore))
	}

	if req.NewsSentiment < 0.4 {
		response.Reasons = append(response.Reasons, "Negative news sentiment detected")
	} else if req.NewsSentiment >= 0.6 {
		response.Reasons = append(response.Reasons, "Positive news sentiment")
	} else {
		response.Reasons = append(response.Reasons, "Neutral news sentiment")
	}

	if req.StockVolatility > 0.5 {
		response.Reasons = append(response.Reasons, "High stock volatility indicates market uncertainty")
	}

	// Regulatory compliance check
	if req.ESGScore < 3.0 {
		response.Reasons = append(response.Reasons, "High regulatory compliance risk")
	}

	return response, nil
}
