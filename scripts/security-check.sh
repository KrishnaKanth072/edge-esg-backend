#!/bin/bash

# EDGE ESG Backend - Security Validation Script
# Checks for common security issues before deployment

echo "=========================================="
echo "EDGE ESG Backend - Security Check"
echo "=========================================="
echo ""

ERRORS=0
WARNINGS=0

# Check 1: No hardcoded secrets in code
echo "🔍 Checking for hardcoded secrets..."
if git grep -i "password.*=.*['\"]" -- "*.go" "*.yaml" "*.yml" 2>/dev/null | grep -v ".example" | grep -v "CHANGE_ME"; then
    echo "❌ FAIL: Hardcoded passwords found in code"
    ERRORS=$((ERRORS + 1))
else
    echo "✅ PASS: No hardcoded passwords found"
fi

# Check 2: .env file not committed
echo ""
echo "🔍 Checking if .env is in .gitignore..."
if grep -q "^\.env$" .gitignore; then
    echo "✅ PASS: .env is in .gitignore"
else
    echo "❌ FAIL: .env not in .gitignore"
    ERRORS=$((ERRORS + 1))
fi

# Check 3: No .env in Git history
echo ""
echo "🔍 Checking Git history for .env files..."
if git log --all --full-history -- ".env" 2>/dev/null | grep -q "commit"; then
    echo "⚠️  WARNING: .env found in Git history - consider using git-filter-repo to remove"
    WARNINGS=$((WARNINGS + 1))
else
    echo "✅ PASS: No .env in Git history"
fi

# Check 4: Required environment variables documented
echo ""
echo "🔍 Checking .env.example..."
if [ -f ".env.example" ]; then
    if grep -q "ENCRYPTION_KEY" .env.example && \
       grep -q "DATABASE_URL" .env.example && \
       grep -q "REDIS_URL" .env.example; then
        echo "✅ PASS: Required variables documented"
    else
        echo "❌ FAIL: Missing required variables in .env.example"
        ERRORS=$((ERRORS + 1))
    fi
else
    echo "❌ FAIL: .env.example not found"
    ERRORS=$((ERRORS + 1))
fi

# Check 5: No default passwords in .env.example
echo ""
echo "🔍 Checking for default passwords in .env.example..."
if grep -i "password.*=.*2024\|password.*=.*secret\|password.*=.*admin" .env.example 2>/dev/null; then
    echo "⚠️  WARNING: Default-looking passwords in .env.example"
    WARNINGS=$((WARNINGS + 1))
else
    echo "✅ PASS: No obvious default passwords"
fi

# Check 6: Security middleware exists
echo ""
echo "🔍 Checking for security middleware..."
if [ -f "internal/middleware/security_headers.go" ] && \
   [ -f "internal/middleware/input_validation.go" ]; then
    echo "✅ PASS: Security middleware files exist"
else
    echo "❌ FAIL: Security middleware missing"
    ERRORS=$((ERRORS + 1))
fi

# Check 7: TLS configuration present
echo ""
echo "🔍 Checking TLS configuration..."
if grep -q "TLS_ENABLED" .env.example; then
    echo "✅ PASS: TLS configuration present"
else
    echo "⚠️  WARNING: TLS configuration not found"
    WARNINGS=$((WARNINGS + 1))
fi

# Check 8: No SQL injection vulnerabilities (basic check)
echo ""
echo "🔍 Checking for potential SQL injection..."
if git grep -i "fmt.Sprintf.*SELECT\|fmt.Sprintf.*INSERT\|fmt.Sprintf.*UPDATE" -- "*.go" 2>/dev/null; then
    echo "⚠️  WARNING: Potential SQL injection - use parameterized queries"
    WARNINGS=$((WARNINGS + 1))
else
    echo "✅ PASS: No obvious SQL injection patterns"
fi

# Check 9: CORS configuration
echo ""
echo "🔍 Checking CORS configuration..."
if [ -f "internal/middleware/cors.go" ]; then
    echo "✅ PASS: CORS middleware exists"
else
    echo "⚠️  WARNING: CORS middleware not found"
    WARNINGS=$((WARNINGS + 1))
fi

# Check 10: Rate limiting
echo ""
echo "🔍 Checking rate limiting..."
if [ -f "internal/middleware/rate_limit.go" ]; then
    echo "✅ PASS: Rate limiting middleware exists"
else
    echo "❌ FAIL: Rate limiting middleware missing"
    ERRORS=$((ERRORS + 1))
fi

# Summary
echo ""
echo "=========================================="
echo "Security Check Summary"
echo "=========================================="
echo "Errors: $ERRORS"
echo "Warnings: $WARNINGS"
echo ""

if [ $ERRORS -gt 0 ]; then
    echo "❌ SECURITY CHECK FAILED"
    echo "Fix all errors before deploying to production"
    exit 1
elif [ $WARNINGS -gt 0 ]; then
    echo "⚠️  SECURITY CHECK PASSED WITH WARNINGS"
    echo "Review warnings before deploying to production"
    exit 0
else
    echo "✅ SECURITY CHECK PASSED"
    echo "All security checks passed successfully"
    exit 0
fi
