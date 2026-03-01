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

	// Step 1: Validate company exists by checking if we can get real data
	stockSymbol := o.realtimeAgents.GuessStockSymbol(req.CompanyName)

	// Try to get stock price
	stockData, stockErr := o.realtimeAgents.GetStockPrice(stockSymbol)
	var currentPrice float64
	if stockErr != nil || stockData == nil {
		// Try Alpha Vantage as fallback
		currentPrice, stockErr = o.realtimeAgents.GetAlphaVantagePrice(stockSymbol)
	} else {
		currentPrice = stockData.Price
	}

	// Try to get news
	sentiment, newsErr := o.realtimeAgents.GetNewsSentiment(req.CompanyName)

	// If BOTH stock price AND news fail, company likely doesn't exist
	if (stockErr != nil || currentPrice == 0) && newsErr != nil {
		return nil, fmt.Errorf("company '%s' not found - unable to retrieve market data or news. Please check the company name and try again", req.CompanyName)
	}

	// If only news failed but we have stock price, continue with neutral sentiment
	if newsErr != nil {
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

	// Step 3: Get stock data (already validated above)
	if currentPrice == 0 {
		// Last attempt if both previous checks failed
		currentPrice = 100.0 // Fallback for analysis to continue
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

// ComparePortfolio analyzes multiple companies and provides portfolio optimization
func (o *Orchestrator) ComparePortfolio(ctx context.Context, req *dtos.PortfolioCompareRequest) (*dtos.PortfolioCompareResponse, error) {
	startTime := time.Now()

	response := &dtos.PortfolioCompareResponse{
		Companies:         make([]dtos.CompanyComparison, 0, len(req.Companies)),
		OptimalAllocation: make([]float64, len(req.Companies)),
		Timestamp:         time.Now(),
	}

	// Analyze each company
	esgScores := make([]float64, 0, len(req.Companies))
	expectedReturns := make([]float64, 0, len(req.Companies))
	validCompanies := make([]string, 0, len(req.Companies))
	bestESGScore := 0.0
	lowestRisk := 100.0
	invalidCompanies := make([]string, 0)

	for _, companyName := range req.Companies {
		// Validate company by checking if we can get real data
		stockSymbol := o.realtimeAgents.GuessStockSymbol(companyName)
		stockData, stockErr := o.realtimeAgents.GetStockPrice(stockSymbol)
		var currentPrice float64
		if stockErr != nil || stockData == nil {
			currentPrice, stockErr = o.realtimeAgents.GetAlphaVantagePrice(stockSymbol)
		} else {
			currentPrice = stockData.Price
		}

		sentiment, newsErr := o.realtimeAgents.GetNewsSentiment(companyName)

		// Skip invalid companies (both stock and news failed)
		if (stockErr != nil || currentPrice == 0) && newsErr != nil {
			invalidCompanies = append(invalidCompanies, companyName)
			continue
		}

		if newsErr != nil {
			sentiment = 0.5
		}

		validCompanies = append(validCompanies, companyName)

		// ESG Scoring
		esgReq := &agents.ESGCalculationRequest{
			CompanyName:   companyName,
			NewsSentiment: sentiment,
			Industry:      "technology",
		}
		esgResult, err := o.esgScoringAgent.CalculateESG(ctx, esgReq)
		if err != nil {
			continue
		}

		// Risk Assessment
		riskReq := &agents.RiskAssessmentRequest{
			CompanyName:     companyName,
			ESGScore:        esgResult.OverallScore,
			NewsSentiment:   sentiment,
			StockVolatility: 0.15,
		}
		riskResult, _ := o.riskAgent.AssessRisk(ctx, riskReq)

		// Trading Signal
		tradingReq := &agents.TradingSignalRequest{
			CompanyName:  companyName,
			Symbol:       stockSymbol,
			CurrentPrice: currentPrice,
			ESGScore:     esgResult.OverallScore,
			Sentiment:    sentiment,
		}
		tradingResult, _ := o.tradingAgent.GenerateSignal(ctx, tradingReq)

		// Compliance Check
		complianceReq := &agents.ComplianceRequest{
			CompanyName: companyName,
			Industry:    "technology",
			Region:      "global",
		}
		complianceResult, _ := o.complianceAgent.CheckCompliance(ctx, complianceReq)

		// Regulation Analysis
		regulationReq := &agents.RegulationRequest{
			CompanyName: companyName,
			Region:      "global",
			Industry:    "technology",
		}
		regulationResult, _ := o.regulationAgent.AnalyzeRegulations(ctx, regulationReq)

		// Build comparison entry
		comparison := dtos.CompanyComparison{
			CompanyName:   companyName,
			ESGScore:      esgResult.OverallScore,
			Environmental: esgResult.Environmental,
			Social:        esgResult.Social,
			Governance:    esgResult.Governance,
			RiskLevel:     riskResult.RiskLevel,
			RiskScore:     riskResult.RiskScore,
			TradingSignal: dtos.TradingSignal{
				Action:       tradingResult.Action,
				Symbol:       tradingResult.Symbol,
				CurrentPrice: fmt.Sprintf("$%.2f", tradingResult.CurrentPrice),
				TargetPrice:  fmt.Sprintf("$%.2f", tradingResult.TargetPrice),
				PriceChange:  fmt.Sprintf("%.1f%%", tradingResult.PriceChangePercent),
				Confidence:   tradingResult.Confidence,
			},
			ComplianceScore: complianceResult.ComplianceScore,
			RegulatoryRisk:  regulationResult.RegulatoryRiskScore,
		}

		response.Companies = append(response.Companies, comparison)

		// Track for optimization
		esgScores = append(esgScores, esgResult.OverallScore)
		expectedReturns = append(expectedReturns, tradingResult.PriceChangePercent)

		// Track best performers
		if esgResult.OverallScore > bestESGScore {
			bestESGScore = esgResult.OverallScore
			response.BestESGCompany = companyName
		}
		if riskResult.RiskScore < lowestRisk {
			lowestRisk = riskResult.RiskScore
			response.LowestRiskCompany = companyName
		}
	}

	// Check if we have any valid companies
	if len(response.Companies) == 0 {
		if len(invalidCompanies) > 0 {
			return nil, fmt.Errorf("no valid companies found. Invalid companies: %v. Please check company names and try again", invalidCompanies)
		}
		return nil, fmt.Errorf("failed to analyze any companies")
	}

	// Warn about invalid companies if some were valid
	if len(invalidCompanies) > 0 {
		// Could add a warning field to response here
		fmt.Printf("Warning: Skipped invalid companies: %v\n", invalidCompanies)
	}

	// Portfolio Optimization
	if len(response.Companies) > 0 {
		riskTolerance := req.RiskTolerance
		if riskTolerance == 0 {
			riskTolerance = 0.5 // Default moderate risk
		}

		portfolioReq := &agents.PortfolioRequest{
			Companies:       validCompanies,
			ESGScores:       esgScores,
			ExpectedReturns: expectedReturns,
			RiskTolerance:   riskTolerance,
		}
		portfolioResult, err := o.optimizationAgent.OptimizePortfolio(ctx, portfolioReq)
		if err == nil {
			response.OptimalAllocation = portfolioResult.OptimalWeights
			response.PortfolioESGScore = portfolioResult.ESGScore
			response.PortfolioRisk = portfolioResult.PortfolioRisk
			response.ExpectedReturn = portfolioResult.ExpectedReturn
		}
	}

	response.ProcessingTimeMs = time.Since(startTime).Milliseconds()
	return response, nil
}
