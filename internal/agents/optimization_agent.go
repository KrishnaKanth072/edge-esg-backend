package agents

import (
	"context"
	"math"
)

type OptimizationAgent struct {
}

func NewOptimizationAgent() *OptimizationAgent {
	return &OptimizationAgent{}
}

type PortfolioRequest struct {
	Companies       []string
	ESGScores       []float64
	ExpectedReturns []float64
	RiskTolerance   float64 // 0-1
}

type PortfolioResponse struct {
	OptimalWeights []float64
	ExpectedReturn float64
	PortfolioRisk  float64
	ESGScore       float64
}

// OptimizePortfolio performs portfolio optimization with ESG constraints
func (o *OptimizationAgent) OptimizePortfolio(ctx context.Context, req *PortfolioRequest) (*PortfolioResponse, error) {
	n := len(req.Companies)
	if n == 0 {
		return &PortfolioResponse{}, nil
	}

	response := &PortfolioResponse{
		OptimalWeights: make([]float64, n),
	}

	// Simple optimization: weight by ESG score and expected return
	totalScore := 0.0
	scores := make([]float64, n)

	for i := 0; i < n; i++ {
		// Composite score: ESG (60%) + Expected Return (40%)
		esgWeight := 0.6
		returnWeight := 0.4

		// Normalize ESG score (0-10 to 0-1)
		normalizedESG := req.ESGScores[i] / 10.0

		// Normalize expected return (assume -20% to +50% range)
		normalizedReturn := (req.ExpectedReturns[i] + 20.0) / 70.0
		if normalizedReturn < 0 {
			normalizedReturn = 0
		}
		if normalizedReturn > 1 {
			normalizedReturn = 1
		}

		// Adjust weights based on risk tolerance
		// High risk tolerance = more weight on returns
		// Low risk tolerance = more weight on ESG
		esgWeight = 0.8 - (req.RiskTolerance * 0.4)
		returnWeight = 0.2 + (req.RiskTolerance * 0.4)

		scores[i] = (normalizedESG * esgWeight) + (normalizedReturn * returnWeight)
		totalScore += scores[i]
	}

	// Calculate optimal weights
	for i := 0; i < n; i++ {
		response.OptimalWeights[i] = scores[i] / totalScore
	}

	// Calculate portfolio metrics
	response.ExpectedReturn = 0.0
	response.ESGScore = 0.0
	for i := 0; i < n; i++ {
		response.ExpectedReturn += response.OptimalWeights[i] * req.ExpectedReturns[i]
		response.ESGScore += response.OptimalWeights[i] * req.ESGScores[i]
	}

	// Calculate portfolio risk (simplified)
	variance := 0.0
	for i := 0; i < n; i++ {
		// Risk increases with weight concentration
		variance += math.Pow(response.OptimalWeights[i], 2) * math.Abs(req.ExpectedReturns[i])
	}
	response.PortfolioRisk = math.Sqrt(variance)

	// Adjust risk based on ESG (higher ESG = lower risk)
	esgRiskReduction := (response.ESGScore / 10.0) * 0.2
	response.PortfolioRisk *= (1.0 - esgRiskReduction)

	return response, nil
}
