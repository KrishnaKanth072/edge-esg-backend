package agents

import (
	"context"
	"strings"
)

type ComplianceAgent struct {
}

func NewComplianceAgent() *ComplianceAgent {
	return &ComplianceAgent{}
}

type ComplianceRequest struct {
	CompanyName string
	Industry    string
	Region      string
}

type ComplianceResponse struct {
	IsCompliant     bool
	Violations      []string
	Regulations     []string
	ComplianceScore float64 // 0-100
}

// CheckCompliance verifies regulatory compliance
func (c *ComplianceAgent) CheckCompliance(ctx context.Context, req *ComplianceRequest) (*ComplianceResponse, error) {
	response := &ComplianceResponse{
		IsCompliant:     true,
		Violations:      []string{},
		Regulations:     []string{},
		ComplianceScore: 100.0,
	}

	companyLower := strings.ToLower(req.CompanyName)
	industryLower := strings.ToLower(req.Industry)

	// Global regulations
	response.Regulations = append(response.Regulations, "ISO 14001 Environmental Management")
	response.Regulations = append(response.Regulations, "GRI Sustainability Reporting Standards")

	// Region-specific regulations
	if strings.Contains(strings.ToLower(req.Region), "eu") || strings.Contains(strings.ToLower(req.Region), "europe") {
		response.Regulations = append(response.Regulations, "EU Taxonomy Regulation")
		response.Regulations = append(response.Regulations, "CSRD (Corporate Sustainability Reporting Directive)")
		response.Regulations = append(response.Regulations, "SFDR (Sustainable Finance Disclosure Regulation)")
	}

	if strings.Contains(strings.ToLower(req.Region), "us") || strings.Contains(strings.ToLower(req.Region), "america") {
		response.Regulations = append(response.Regulations, "SEC Climate Disclosure Rules")
		response.Regulations = append(response.Regulations, "EPA Environmental Regulations")
	}

	if strings.Contains(strings.ToLower(req.Region), "india") || strings.Contains(strings.ToLower(req.Region), "in") {
		response.Regulations = append(response.Regulations, "SEBI BRSR (Business Responsibility and Sustainability Reporting)")
		response.Regulations = append(response.Regulations, "Companies Act 2013 CSR Requirements")
	}

	// Industry-specific checks
	if strings.Contains(industryLower, "oil") || strings.Contains(industryLower, "gas") ||
		strings.Contains(companyLower, "petroleum") {
		response.Regulations = append(response.Regulations, "Methane Emissions Regulations")
		response.Regulations = append(response.Regulations, "Oil Spill Prevention Requirements")
		// Oil companies face stricter scrutiny
		response.ComplianceScore -= 15.0
		response.Violations = append(response.Violations, "High-risk industry requiring enhanced monitoring")
	}

	if strings.Contains(industryLower, "finance") || strings.Contains(industryLower, "bank") {
		response.Regulations = append(response.Regulations, "Basel III ESG Risk Management")
		response.Regulations = append(response.Regulations, "Green Finance Guidelines")
	}

	if strings.Contains(industryLower, "manufacturing") || strings.Contains(industryLower, "industrial") {
		response.Regulations = append(response.Regulations, "Industrial Emissions Directive")
		response.Regulations = append(response.Regulations, "Waste Management Regulations")
	}

	// Check for known violators (simplified - in production, check against real database)
	if strings.Contains(companyLower, "tobacco") || strings.Contains(companyLower, "cigarette") {
		response.Violations = append(response.Violations, "Tobacco industry - restricted under ESG frameworks")
		response.ComplianceScore -= 40.0
		response.IsCompliant = false
	}

	if strings.Contains(companyLower, "weapons") || strings.Contains(companyLower, "defense") {
		response.Violations = append(response.Violations, "Defense industry - ethical concerns")
		response.ComplianceScore -= 25.0
	}

	// Determine overall compliance
	if response.ComplianceScore < 60 {
		response.IsCompliant = false
	}

	if response.ComplianceScore < 0 {
		response.ComplianceScore = 0
	}

	return response, nil
}
