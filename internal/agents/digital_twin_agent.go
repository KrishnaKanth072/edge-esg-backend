package agents

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type DigitalTwinAgent struct {
	twins map[string]*DigitalTwin
}

type DigitalTwin struct {
	TwinID    string
	Company   string
	ESGScore  float64
	MarketCap float64
	Industry  string
	CreatedAt time.Time
}

func NewDigitalTwinAgent() *DigitalTwinAgent {
	return &DigitalTwinAgent{
		twins: make(map[string]*DigitalTwin),
	}
}

type TwinCreationRequest struct {
	CompanyName string
	ESGScore    float64
	MarketCap   float64
	Industry    string
}

type TwinCreationResponse struct {
	TwinID    string
	ModelData string
	Success   bool
}

type ScenarioRequest struct {
	TwinID       string
	ScenarioType string  // CARBON_TAX, REGULATION_CHANGE, etc.
	ImpactFactor float64 // Magnitude of change
}

type ScenarioResponse struct {
	PredictedESGChange   float64
	PredictedValueChange float64
	Recommendations      []string
}

// CreateTwin creates a digital twin model of the company
func (d *DigitalTwinAgent) CreateTwin(ctx context.Context, req *TwinCreationRequest) (*TwinCreationResponse, error) {
	// Generate unique twin ID
	twinData := fmt.Sprintf("%s:%f:%d", req.CompanyName, req.ESGScore, time.Now().Unix())
	hashBytes := sha256.Sum256([]byte(twinData))
	twinID := hex.EncodeToString(hashBytes[:8])

	// Create twin
	twin := &DigitalTwin{
		TwinID:    twinID,
		Company:   req.CompanyName,
		ESGScore:  req.ESGScore,
		MarketCap: req.MarketCap,
		Industry:  req.Industry,
		CreatedAt: time.Now(),
	}

	d.twins[twinID] = twin

	modelData := fmt.Sprintf("Digital Twin Model: Company=%s, ESG=%.1f, MarketCap=%.2f, Industry=%s",
		req.CompanyName, req.ESGScore, req.MarketCap, req.Industry)

	return &TwinCreationResponse{
		TwinID:    twinID,
		ModelData: modelData,
		Success:   true,
	}, nil
}

// SimulateScenario runs what-if scenarios on the digital twin
func (d *DigitalTwinAgent) SimulateScenario(ctx context.Context, req *ScenarioRequest) (*ScenarioResponse, error) {
	twin, exists := d.twins[req.TwinID]
	if !exists {
		return &ScenarioResponse{
			Recommendations: []string{"Twin not found - create twin first"},
		}, nil
	}

	response := &ScenarioResponse{
		Recommendations: []string{},
	}

	switch req.ScenarioType {
	case "CARBON_TAX":
		// Carbon tax impact
		response.PredictedESGChange = -req.ImpactFactor * 0.5
		response.PredictedValueChange = -req.ImpactFactor * 2.0 // 2% value loss per unit
		response.Recommendations = append(response.Recommendations,
			"Invest in carbon reduction technologies",
			"Explore carbon offset programs",
			"Transition to renewable energy sources")

	case "REGULATION_CHANGE":
		// New ESG regulations
		if twin.ESGScore < 5.0 {
			response.PredictedESGChange = -req.ImpactFactor * 1.0
			response.PredictedValueChange = -req.ImpactFactor * 3.0
			response.Recommendations = append(response.Recommendations,
				"Urgent: Improve ESG compliance",
				"Hire ESG compliance officer",
				"Implement sustainability reporting systems")
		} else {
			response.PredictedESGChange = req.ImpactFactor * 0.3
			response.PredictedValueChange = req.ImpactFactor * 1.5
			response.Recommendations = append(response.Recommendations,
				"Leverage strong ESG position for competitive advantage",
				"Market ESG leadership to investors")
		}

	case "MARKET_SHIFT":
		// Market preference shift toward ESG
		if twin.ESGScore >= 7.0 {
			response.PredictedESGChange = req.ImpactFactor * 0.2
			response.PredictedValueChange = req.ImpactFactor * 5.0
			response.Recommendations = append(response.Recommendations,
				"Capitalize on ESG premium",
				"Expand green product lines")
		} else {
			response.PredictedESGChange = -req.ImpactFactor * 0.3
			response.PredictedValueChange = -req.ImpactFactor * 4.0
			response.Recommendations = append(response.Recommendations,
				"Accelerate ESG improvements",
				"Risk of losing market share to ESG leaders")
		}

	case "SUPPLY_CHAIN_DISRUPTION":
		response.PredictedESGChange = -req.ImpactFactor * 0.4
		response.PredictedValueChange = -req.ImpactFactor * 2.5
		response.Recommendations = append(response.Recommendations,
			"Diversify supply chain",
			"Implement supply chain ESG audits",
			"Build resilience through local sourcing")

	default:
		response.PredictedESGChange = 0.0
		response.PredictedValueChange = 0.0
		response.Recommendations = append(response.Recommendations,
			"Unknown scenario type")
	}

	return response, nil
}
