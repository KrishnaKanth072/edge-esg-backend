package dtos

type AnalyzeResponse struct {
	ESGScore         string        `json:"esg_score"`
	RiskAction       string        `json:"risk_action"`
	TradingSignal    TradingSignal `json:"trading_signal"`
	AuditHash        string        `json:"audit_hash"`
	ProcessingTimeMs int64         `json:"processing_time_ms"`
	MaskedData       bool          `json:"masked_data"`
}

type TradingSignal struct {
	Action      string  `json:"action"`
	Symbol      string  `json:"symbol"`
	TargetPrice string  `json:"target_price"`
	Confidence  float64 `json:"confidence"`
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
