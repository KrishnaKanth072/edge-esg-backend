package agents

import (
	"context"
	"strings"
)

type RegulationAgent struct {
}

func NewRegulationAgent() *RegulationAgent {
	return &RegulationAgent{}
}

type RegulationRequest struct {
	CompanyName string
	Region      string
	Industry    string
}

type RegulationResponse struct {
	ApplicableRegulations []string
	UpcomingChanges       []string
	RegulatoryRiskScore   float64 // 0-100
	Recommendations       []string
}

// AnalyzeRegulations assesses regulatory landscape
func (r *RegulationAgent) AnalyzeRegulations(ctx context.Context, req *RegulationRequest) (*RegulationResponse, error) {
	response := &RegulationResponse{
		ApplicableRegulations: []string{},
		UpcomingChanges:       []string{},
		Recommendations:       []string{},
		RegulatoryRiskScore:   30.0, // Base risk
	}

	regionLower := strings.ToLower(req.Region)
	industryLower := strings.ToLower(req.Industry)

	// EU Regulations (strictest)
	if strings.Contains(regionLower, "eu") || strings.Contains(regionLower, "europe") {
		response.ApplicableRegulations = append(response.ApplicableRegulations,
			"EU Taxonomy for Sustainable Activities",
			"Corporate Sustainability Reporting Directive (CSRD)",
			"EU Green Deal",
			"Carbon Border Adjustment Mechanism (CBAM)")
		response.UpcomingChanges = append(response.UpcomingChanges,
			"CSRD mandatory reporting from 2024",
			"CBAM full implementation by 2026")
		response.RegulatoryRiskScore += 25.0
		response.Recommendations = append(response.Recommendations,
			"Prepare for CSRD double materiality assessment",
			"Implement carbon accounting systems")
	}

	// US Regulations
	if strings.Contains(regionLower, "us") || strings.Contains(regionLower, "america") {
		response.ApplicableRegulations = append(response.ApplicableRegulations,
			"SEC Climate Disclosure Rules",
			"EPA Greenhouse Gas Reporting",
			"Inflation Reduction Act incentives")
		response.UpcomingChanges = append(response.UpcomingChanges,
			"SEC climate rules phased implementation 2024-2026")
		response.RegulatoryRiskScore += 15.0
	}

	// India Regulations
	if strings.Contains(regionLower, "india") {
		response.ApplicableRegulations = append(response.ApplicableRegulations,
			"SEBI BRSR (Business Responsibility and Sustainability Reporting)",
			"Companies Act 2013 - CSR provisions",
			"National Action Plan on Climate Change")
		response.RegulatoryRiskScore += 10.0
		response.Recommendations = append(response.Recommendations,
			"Ensure BRSR Core compliance for top 1000 listed entities")
	}

	// Industry-specific regulations
	if strings.Contains(industryLower, "finance") || strings.Contains(industryLower, "bank") {
		response.ApplicableRegulations = append(response.ApplicableRegulations,
			"Basel III ESG Risk Management",
			"TCFD Recommendations",
			"Green Finance Guidelines")
		response.RegulatoryRiskScore += 20.0
		response.Recommendations = append(response.Recommendations,
			"Integrate climate risk into credit assessments",
			"Develop green lending portfolio")
	}

	if strings.Contains(industryLower, "energy") || strings.Contains(industryLower, "oil") {
		response.ApplicableRegulations = append(response.ApplicableRegulations,
			"Methane Emissions Regulations",
			"Renewable Energy Mandates",
			"Carbon Pricing Mechanisms")
		response.RegulatoryRiskScore += 35.0
		response.Recommendations = append(response.Recommendations,
			"Develop transition plan to renewable energy",
			"Implement methane leak detection systems")
	}

	if strings.Contains(industryLower, "manufacturing") {
		response.ApplicableRegulations = append(response.ApplicableRegulations,
			"Industrial Emissions Directive",
			"Extended Producer Responsibility",
			"Circular Economy Regulations")
		response.RegulatoryRiskScore += 18.0
	}

	// Clamp risk score
	if response.RegulatoryRiskScore > 100 {
		response.RegulatoryRiskScore = 100
	}

	// General recommendations
	if response.RegulatoryRiskScore > 50 {
		response.Recommendations = append(response.Recommendations,
			"High regulatory risk - establish dedicated compliance team",
			"Conduct regular regulatory horizon scanning")
	}

	return response, nil
}
