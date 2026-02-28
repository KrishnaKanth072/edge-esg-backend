package types

type RiskLevel string

const (
	RiskLow      RiskLevel = "LOW"
	RiskMedium   RiskLevel = "MEDIUM"
	RiskHigh     RiskLevel = "HIGH"
	RiskCritical RiskLevel = "CRITICAL"
)

type TradeAction string

const (
	ActionBuy  TradeAction = "BUY"
	ActionSell TradeAction = "SELL"
	ActionHold TradeAction = "HOLD"
)

type RiskAction string

const (
	RiskApprove RiskAction = "APPROVE"
	RiskReject  RiskAction = "REJECT"
	RiskReview  RiskAction = "REVIEW"
)

type UserRole string

const (
	RoleTrader     UserRole = "TRADER"
	RoleCompliance UserRole = "COMPLIANCE"
	RoleRisk       UserRole = "RISK"
	RoleAdmin      UserRole = "ADMIN"
)
