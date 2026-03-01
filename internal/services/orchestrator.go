package services

import (
	"context"
	"fmt"
	"time"

	"github.com/edgeesg/edge-esg-backend/internal/dtos"
)

type Orchestrator struct {
	realtimeAgents *RealTimeAgents
}

func NewOrchestrator() *Orchestrator {
	return &Orchestrator{
		realtimeAgents: NewRealTimeAgents(),
	}
}

// Execute8LayerPipeline orchestrates all 10 agents with REAL-TIME data ONLY
func (o *Orchestrator) Execute8LayerPipeline(ctx context.Context, req *dtos.AnalyzeRequest) (*dtos.AnalyzeResponse, error) {
	startTime := time.Now()

	// Use ONLY real-time agents - NO FALLBACK
	analysis, err := o.realtimeAgents.AnalyzeCompany(ctx, req.CompanyName)
	if err != nil {
		return nil, fmt.Errorf("failed to analyze company with real data: %w", err)
	}

	// Convert to response format
	response := &dtos.AnalyzeResponse{
		ESGScore:   fmt.Sprintf("%.1f/10", analysis.ESGScore),
		RiskAction: analysis.RiskAction,
		TradingSignal: dtos.TradingSignal{
			Action:       analysis.TradingSignal,
			Symbol:       analysis.StockSymbol,
			CurrentPrice: fmt.Sprintf("$%.2f", analysis.CurrentPrice),
			TargetPrice:  fmt.Sprintf("$%.2f", analysis.TargetPrice),
			PriceChange:  fmt.Sprintf("%.1f%%", analysis.PriceChange),
			Confidence:   float64(analysis.Confidence) / 100.0,
		},
		RiskReasons:      analysis.RiskReasons,
		ProcessingTimeMs: time.Since(startTime).Milliseconds(),
		AuditHash:        fmt.Sprintf("0x%x", time.Now().Unix()),
		Timestamp:        analysis.LastUpdated,
	}

	return response, nil
}

// Client interfaces for future gRPC implementation
type RiskAgentClient interface{}
type TradingAgentClient interface{}
type QuantumAgentClient interface{}
type ComplianceAgentClient interface{}
type ConsensusAgentClient interface{}
type BlockchainAgentClient interface{}
