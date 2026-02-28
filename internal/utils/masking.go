package utils

import (
	"fmt"
	"strings"
)

// MaskRevenue masks revenue data based on role
func MaskRevenue(revenue float64, role string) string {
	if role == "COMPLIANCE" || role == "ADMIN" {
		return fmt.Sprintf("₹%.2fCr", revenue)
	}
	return "₹XXCr"
}

// MaskEmissions masks carbon emissions based on role
func MaskEmissions(emissions float64, role string) string {
	if role == "COMPLIANCE" || role == "ADMIN" {
		return fmt.Sprintf("%.2fM tCO2e", emissions)
	}
	return "X.XM tCO2e"
}

// MaskCompanyDetails masks sensitive company info
func MaskCompanyDetails(details string, role string) string {
	if role == "COMPLIANCE" || role == "ADMIN" {
		return details
	}
	return strings.Repeat("*", len(details)/2) + details[len(details)/2:]
}
