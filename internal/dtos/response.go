package dtos

import "time"

type AnalyzeResponse struct {
	ESGScore              string                   `json:"esg_score"`
	RiskAction            string                   `json:"risk_action"`
	RiskReasons           []string                 `json:"risk_reasons,omitempty"`
	TradingSignal         TradingSignal            `json:"trading_signal"`
	HistoricalReturns     []map[string]interface{} `json:"historical_returns,omitempty"`
	InvestmentProjections []map[string]interface{} `json:"investment_projections,omitempty"`
	AuditHash             string                   `json:"audit_hash"`
	ProcessingTimeMs      int64                    `json:"processing_time_ms"`
	MaskedData            bool                     `json:"masked_data"`
	Timestamp             time.Time                `json:"timestamp"`
}

type TradingSignal struct {
	Action       string  `json:"action"`
	Symbol       string  `json:"symbol"`
	CurrentPrice string  `json:"current_price,omitempty"`
	TargetPrice  string  `json:"target_price"`
	PriceChange  string  `json:"price_change,omitempty"`
	Confidence   float64 `json:"confidence"`
}

type HealthResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
	Uptime  int64  `json:"uptime_seconds"`
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

type CompanyComparison struct {
	CompanyName     string        `json:"company_name"`
	ESGScore        float64       `json:"esg_score"`
	Environmental   float64       `json:"environmental"`
	Social          float64       `json:"social"`
	Governance      float64       `json:"governance"`
	RiskLevel       string        `json:"risk_level"`
	RiskScore       float64       `json:"risk_score"`
	TradingSignal   TradingSignal `json:"trading_signal"`
	ComplianceScore float64       `json:"compliance_score"`
	RegulatoryRisk  float64       `json:"regulatory_risk"`
}

type PortfolioCompareResponse struct {
	Companies         []CompanyComparison `json:"companies"`
	OptimalAllocation []float64           `json:"optimal_allocation"`
	PortfolioESGScore float64             `json:"portfolio_esg_score"`
	PortfolioRisk     float64             `json:"portfolio_risk"`
	ExpectedReturn    float64             `json:"expected_return"`
	BestESGCompany    string              `json:"best_esg_company"`
	LowestRiskCompany string              `json:"lowest_risk_company"`
	ProcessingTimeMs  int64               `json:"processing_time_ms"`
	Timestamp         time.Time           `json:"timestamp"`
}
