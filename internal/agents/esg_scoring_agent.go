package agents

import (
	"context"
	"strings"
)

type ESGScoringAgent struct {
}

func NewESGScoringAgent() *ESGScoringAgent {
	return &ESGScoringAgent{}
}

type ESGCalculationRequest struct {
	CompanyName   string
	NewsSentiment float64
	Industry      string
}

type ESGCalculationResponse struct {
	OverallScore  float64 // 0-10
	Environmental float64
	Social        float64
	Governance    float64
	Factors       []string
}

// CalculateESG computes comprehensive ESG scores
func (e *ESGScoringAgent) CalculateESG(ctx context.Context, req *ESGCalculationRequest) (*ESGCalculationResponse, error) {
	response := &ESGCalculationResponse{
		Factors: []string{},
	}

	// Base score
	baseScore := 5.0
	companyLower := strings.ToLower(req.CompanyName)

	// Environmental Score (0-10)
	envScore := 5.0

	// Green/Renewable companies
	if strings.Contains(companyLower, "solar") || strings.Contains(companyLower, "wind") ||
		strings.Contains(companyLower, "renewable") || strings.Contains(companyLower, "green") ||
		strings.Contains(companyLower, "clean") || strings.Contains(companyLower, "sustainable") {
		envScore += 3.5
		response.Factors = append(response.Factors, "Renewable energy sector")
		baseScore += 2.5
	}

	// Tech companies (moderate environmental impact)
	if strings.Contains(companyLower, "tech") || strings.Contains(companyLower, "software") ||
		strings.Contains(companyLower, "digital") || strings.Contains(companyLower, "cloud") {
		envScore += 1.5
		response.Factors = append(response.Factors, "Technology sector with lower emissions")
		baseScore += 1.5
	}

	// Electric vehicles
	if strings.Contains(companyLower, "tesla") || strings.Contains(companyLower, "electric") ||
		strings.Contains(companyLower, "ev ") {
		envScore += 2.5
		response.Factors = append(response.Factors, "Electric vehicle innovation")
		baseScore += 2.0
	}

	// Heavy industry (negative)
	if strings.Contains(companyLower, "steel") || strings.Contains(companyLower, "cement") ||
		strings.Contains(companyLower, "mining") || strings.Contains(companyLower, "coal") {
		envScore -= 2.5
		response.Factors = append(response.Factors, "Heavy industrial emissions")
		baseScore -= 1.5
	}

	// Oil & Gas (very negative)
	if strings.Contains(companyLower, "oil") || strings.Contains(companyLower, "petroleum") ||
		strings.Contains(companyLower, "exxon") || strings.Contains(companyLower, "chevron") ||
		strings.Contains(companyLower, "shell") || strings.Contains(companyLower, "bp") {
		envScore -= 4.0
		response.Factors = append(response.Factors, "Fossil fuel industry")
		baseScore -= 3.0
	}

	// Social Score (0-10)
	socialScore := 5.0

	// Tech giants (good social programs)
	if strings.Contains(companyLower, "apple") || strings.Contains(companyLower, "microsoft") ||
		strings.Contains(companyLower, "google") || strings.Contains(companyLower, "alphabet") {
		socialScore += 2.0
		response.Factors = append(response.Factors, "Strong employee programs")
	}

	// Tobacco/weapons (negative social impact)
	if strings.Contains(companyLower, "tobacco") || strings.Contains(companyLower, "cigarette") ||
		strings.Contains(companyLower, "defense") || strings.Contains(companyLower, "weapons") {
		socialScore -= 3.5
		response.Factors = append(response.Factors, "Controversial social impact")
		baseScore -= 3.0
	}

	// Governance Score (0-10)
	govScore := 5.0

	// Large established companies (better governance)
	if strings.Contains(companyLower, "tata") || strings.Contains(companyLower, "reliance") ||
		strings.Contains(companyLower, "infosys") || strings.Contains(companyLower, "wipro") {
		govScore += 1.5
		response.Factors = append(response.Factors, "Established corporate governance")
	}

	// Sentiment adjustment (affects all pillars)
	sentimentImpact := (req.NewsSentiment - 0.5) * 6.0
	baseScore += sentimentImpact
	envScore += sentimentImpact * 0.3
	socialScore += sentimentImpact * 0.3
	govScore += sentimentImpact * 0.4

	if req.NewsSentiment >= 0.6 {
		response.Factors = append(response.Factors, "Positive news sentiment")
	} else if req.NewsSentiment <= 0.4 {
		response.Factors = append(response.Factors, "Negative news sentiment")
	}

	// Clamp scores to 0-10
	clamp := func(score float64) float64 {
		if score < 0 {
			return 0
		}
		if score > 10 {
			return 10
		}
		return score
	}

	response.Environmental = clamp(envScore)
	response.Social = clamp(socialScore)
	response.Governance = clamp(govScore)
	response.OverallScore = clamp(baseScore)

	return response, nil
}
