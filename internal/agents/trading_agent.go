package agents

import (
	"context"
	"fmt"
)

type TradingAgent struct {
}

func NewTradingAgent() *TradingAgent {
	return &TradingAgent{}
}

type TradingSignalRequest struct {
	CompanyName  string
	Symbol       string
	CurrentPrice float64
	ESGScore     float64
	Sentiment    float64
}

type TradingSignalResponse struct {
	Action             string // BUY/SELL/HOLD
	Symbol             string
	CurrentPrice       float64
	TargetPrice        float64
	PriceChangePercent float64
	Confidence         float64
	Reasoning          string
	HistoricalReturns  []HistoricalReturn // New field
}

type HistoricalReturn struct {
	Period     string // "1 Month", "3 Months", "6 Months", "1 Year"
	StartDate  string // "2024-01-01"
	EndDate    string // "2024-02-01"
	StartPrice float64
	EndPrice   float64
	ReturnPct  float64 // Percentage return
}

// GenerateSignal creates trading recommendations based on ESG and market data
func (t *TradingAgent) GenerateSignal(ctx context.Context, req *TradingSignalRequest) (*TradingSignalResponse, error) {
	response := &TradingSignalResponse{
		Symbol:       req.Symbol,
		CurrentPrice: req.CurrentPrice,
	}

	// Calculate composite score (0-100)
	compositeScore := (req.ESGScore * 10) + (req.Sentiment * 100)
	compositeScore = compositeScore / 2 // Average

	// Strong BUY signal
	if req.ESGScore >= 7.0 && req.Sentiment >= 0.65 {
		response.Action = "BUY"
		response.PriceChangePercent = 15.0 + (req.ESGScore-7.0)*5.0 + (req.Sentiment-0.65)*20.0
		response.Confidence = 0.85
		response.Reasoning = fmt.Sprintf("Strong ESG fundamentals (%.1f/10) combined with positive market sentiment (%.0f%%) indicate growth potential",
			req.ESGScore, req.Sentiment*100)
	} else if req.ESGScore >= 6.0 && req.Sentiment >= 0.55 {
		// Moderate BUY
		response.Action = "BUY"
		response.PriceChangePercent = 8.0 + (req.ESGScore-6.0)*3.0
		response.Confidence = 0.70
		response.Reasoning = fmt.Sprintf("Good ESG profile (%.1f/10) with favorable sentiment suggests moderate upside",
			req.ESGScore)
	} else if req.ESGScore <= 3.5 || req.Sentiment <= 0.35 {
		// SELL signal
		response.Action = "SELL"
		response.PriceChangePercent = -(12.0 + (3.5-req.ESGScore)*4.0)
		response.Confidence = 0.75
		response.Reasoning = fmt.Sprintf("Weak ESG performance (%.1f/10) and negative sentiment (%.0f%%) indicate downside risk",
			req.ESGScore, req.Sentiment*100)
	} else if req.ESGScore <= 4.5 && req.Sentiment <= 0.45 {
		// Weak SELL
		response.Action = "SELL"
		response.PriceChangePercent = -(5.0 + (4.5-req.ESGScore)*2.0)
		response.Confidence = 0.60
		response.Reasoning = "Below-average ESG metrics suggest caution"
	} else {
		// HOLD signal
		response.Action = "HOLD"
		// Small price movement for HOLD based on fundamentals
		response.PriceChangePercent = (req.ESGScore-5.0)*1.5 + (req.Sentiment-0.5)*6.0
		response.Confidence = 0.65
		response.Reasoning = fmt.Sprintf("Balanced ESG score (%.1f/10) with neutral sentiment suggests maintaining position",
			req.ESGScore)
	}

	// Calculate target price
	response.TargetPrice = req.CurrentPrice * (1 + response.PriceChangePercent/100)

	// Clamp confidence
	if response.Confidence > 1.0 {
		response.Confidence = 1.0
	}
	if response.Confidence < 0.0 {
		response.Confidence = 0.0
	}

	return response, nil
}
