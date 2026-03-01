package dtos

type AnalyzeRequest struct {
	CompanyName string `json:"company_name" validate:"required,min=2,max=100"`
	BankID      string `json:"bank_id" validate:"omitempty,uuid"`
	Mode        string `json:"mode" validate:"omitempty,oneof=auto manual"`
	UserRole    string `json:"user_role,omitempty"`
}

type ComplianceAuditRequest struct {
	CompanyName string `json:"company_name" validate:"required"`
	BankID      string `json:"bank_id" validate:"required,uuid"`
	StartDate   string `json:"start_date" validate:"required"`
	EndDate     string `json:"end_date" validate:"required"`
}
