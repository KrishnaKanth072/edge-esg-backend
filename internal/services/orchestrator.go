package services

import (
	"context"
	"fmt"
	"time"

	"github.com/edgeesg/edge-esg-backend/internal/agents"
	"github.com/edgeesg/edge-esg-backend/internal/dtos"
)

type Orchestrator struct {
	realtimeAgents    *RealTimeAgents
	riskAgent         *agents.RiskAgent
	tradingAgent      *agents.TradingAgent
	esgScoringAgent   *agents.ESGScoringAgent
	complianceAgent   *agents.ComplianceAgent
	consensusAgent    *agents.ConsensusAgent
	blockchainAgent   *agents.BlockchainAgent
	quantumAgent      *agents.QuantumAgent
	regulationAgent   *agents.RegulationAgent
	optimizationAgent *agents.OptimizationAgent
	digitalTwinAgent  *agents.DigitalTwinAgent
}

func NewOrchestrator() *Orchestrator {
	return &Orchestrator{
		realtimeAgents:    NewRealTimeAgents(),
		riskAgent:         agents.NewRiskAgent(),
		tradingAgent:      agents.NewTradingAgent(),
		esgScoringAgent:   agents.NewESGScoringAgent(),
		complianceAgent:   agents.NewComplianceAgent(),
		consensusAgent:    agents.NewConsensusAgent(),
		blockchainAgent:   agents.NewBlockchainAgent(),
		quantumAgent:      agents.NewQuantumAgent(),
		regulationAgent:   agents.NewRegulationAgent(),
		optimizationAgent: agents.NewOptimizationAgent(),
		digitalTwinAgent:  agents.NewDigitalTwinAgent(),
	}
}

// Execute8LayerPipeline orchestrates all 10 agents with REAL-TIME data
func (o *Orchestrator) Execute8LayerPipeline(ctx context.Context, req *dtos.AnalyzeRequest) (*dtos.AnalyzeResponse, error) {
	startTime := time.Now()

	// Step 1: Get real-time market data
	sentiment, err := o.realtimeAgents.GetNewsSentiment(req.CompanyName)
	if err != nil {
		sentiment = 0.5 // Neutral default
	}

	// Step 2: ESG Scoring Agent - Calculate comprehensive ESG score
	esgReq := &agents.ESGCalculationRequest{
		CompanyName:   req.CompanyName,
		NewsSentiment: sentiment,
		Industry:      "technology", // TODO: detect industry
	}
	esgResult, err := o.esgScoringAgent.CalculateESG(ctx, esgReq)
	if err != nil {
		return nil, fmt.Errorf("ESG scoring failed: %w", err)
	}

	// Step 3: Get stock data
	stockSymbol := o.realtimeAgents.GuessStockSymbol(req.CompanyName)
	stockData, _ := o.realtimeAgents.GetStockPrice(stockSymbol)
	currentPrice := 0.0
	if stockData != nil {
		currentPrice = stockData.Price
	} else {
		// Try Alpha Vantage fallback
		currentPrice, _ = o.realtimeAgents.GetAlphaVantagePrice(stockSymbol)
	}

	// Step 4: Risk Agent - Assess financing risk
	riskReq := &agents.RiskAssessmentRequest{
		CompanyName:     req.CompanyName,
		ESGScore:        esgResult.OverallScore,
		NewsSentiment:   sentiment,
		StockVolatility: 0.15, // Default volatility
	}
	riskResult, err := o.riskAgent.AssessRisk(ctx, riskReq)
	if err != nil {
		return nil, fmt.Errorf("risk assessment failed: %w", err)
	}

	// Step 5: Trading Agent - Generate trading signal
	tradingReq := &agents.TradingSignalRequest{
		CompanyName:  req.CompanyName,
		Symbol:       stockSymbol,
		CurrentPrice: currentPrice,
		ESGScore:     esgResult.OverallScore,
		Sentiment:    sentiment,
	}
	tradingResult, err := o.tradingAgent.GenerateSignal(ctx, tradingReq)
	if err != nil {
		return nil, fmt.Errorf("trading signal failed: %w", err)
	}

	// Step 6: Compliance Agent - Check regulatory compliance
	complianceReq := &agents.ComplianceRequest{
		CompanyName: req.CompanyName,
		Industry:    "technology",
		Region:      "global",
	}
	complianceResult, err := o.complianceAgent.CheckCompliance(ctx, complianceReq)
	if err != nil {
		return nil, fmt.Errorf("compliance check failed: %w", err)
	}

	// Step 7: Consensus Agent - Aggregate all agent decisions
	consensusReq := &agents.ConsensusRequest{
		CompanyName: req.CompanyName,
		AgentDecisions: []agents.AgentDecision{
			{
				AgentName:  "RiskAgent",
				Decision:   riskResult.Action,
				Confidence: 1.0 - (riskResult.RiskScore / 100.0),
				Reasoning:  fmt.Sprintf("Risk Level: %s", riskResult.RiskLevel),
			},
			{
				AgentName:  "TradingAgent",
				Decision:   tradingResult.Action,
				Confidence: tradingResult.Confidence,
				Reasoning:  tradingResult.Reasoning,
			},
			{
				AgentName:  "ComplianceAgent",
				Decision:   fmt.Sprintf("COMPLIANT_%v", complianceResult.IsCompliant),
				Confidence: complianceResult.ComplianceScore / 100.0,
				Reasoning:  fmt.Sprintf("Compliance Score: %.0f%%", complianceResult.ComplianceScore),
			},
		},
	}
	consensusResult, err := o.consensusAgent.ReachConsensus(ctx, consensusReq)
	if err != nil {
		return nil, fmt.Errorf("consensus failed: %w", err)
	}

	// Step 8: Blockchain Agent - Record audit trail
	auditData := fmt.Sprintf("Company:%s,ESG:%.1f,Risk:%s,Trading:%s,Consensus:%s",
		req.CompanyName, esgResult.OverallScore, riskResult.Action, tradingResult.Action, consensusResult.FinalDecision)
	auditReq := &agents.AuditRecordRequest{
		CompanyName:  req.CompanyName,
		AnalysisData: auditData,
		Timestamp:    time.Now().Unix(),
	}
	auditResult, _ := o.blockchainAgent.RecordAudit(ctx, auditReq)

	// Build comprehensive response
	response := &dtos.AnalyzeResponse{
		ESGScore:   fmt.Sprintf("%.1f/10", esgResult.OverallScore),
		RiskAction: riskResult.Action,
		TradingSignal: dtos.TradingSignal{
			Action:       tradingResult.Action,
			Symbol:       tradingResult.Symbol,
			CurrentPrice: fmt.Sprintf("$%.2f", tradingResult.CurrentPrice),
			TargetPrice:  fmt.Sprintf("$%.2f", tradingResult.TargetPrice),
			PriceChange:  fmt.Sprintf("%.1f%%", tradingResult.PriceChangePercent),
			Confidence:   tradingResult.Confidence,
		},
		RiskReasons:      riskResult.Reasons,
		ProcessingTimeMs: time.Since(startTime).Milliseconds(),
		AuditHash:        auditResult.TransactionHash,
		Timestamp:        time.Now(),
	}

	return response, nil
}
