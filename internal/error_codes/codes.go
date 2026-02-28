package error_codes

type ErrorCode string

const (
	// ESG Errors
	ESGInvalidInput     ErrorCode = "ESG_INVALID_INPUT"
	ESGProcessingFailed ErrorCode = "ESG_PROCESSING_FAILED"
	ESGNotFound         ErrorCode = "ESG_NOT_FOUND"

	// Auth Errors
	AuthUnauthorized ErrorCode = "AUTH_UNAUTHORIZED"
	AuthForbidden    ErrorCode = "AUTH_FORBIDDEN"
	AuthInvalidToken ErrorCode = "AUTH_INVALID_TOKEN"
	AuthExpiredToken ErrorCode = "AUTH_EXPIRED_TOKEN"

	// Validation Errors
	ValidationFailed ErrorCode = "VALIDATION_FAILED"

	// Database Errors
	DBConnectionFailed    ErrorCode = "DB_CONNECTION_FAILED"
	DBQueryFailed         ErrorCode = "DB_QUERY_FAILED"
	DBConstraintViolation ErrorCode = "DB_CONSTRAINT_VIOLATION"

	// Rate Limit Errors
	RateLimitExceeded ErrorCode = "RATE_LIMIT_EXCEEDED"

	// Compliance Errors
	ComplianceViolation   ErrorCode = "COMPLIANCE_VIOLATION"
	ComplianceAuditFailed ErrorCode = "COMPLIANCE_AUDIT_FAILED"
)

func (e ErrorCode) String() string {
	return string(e)
}
