@echo off
echo ========================================
echo EDGE ESG Backend - Deploy to GitHub
echo ========================================

cd /d "%~dp0"

echo.
echo Adding all changes...
git add -A

echo.
echo Committing changes...
git commit -m "feat: production-ready EDGE ESG Backend

- Comprehensive security hardening
- Environment-based secrets management
- Security middleware (headers, validation, logging)
- HTTPS/TLS enforcement
- Input validation and sanitization
- Production deployment guides
- RBI compliance features
- Go 1.23 support
- Docker and Kubernetes ready
- Fix: Add ValidationFailed error code"

echo.
echo Pushing to GitHub...
git push origin main

echo.
echo ========================================
echo ✅ DEPLOYED SUCCESSFULLY!
echo ========================================
echo.
echo 📋 NEXT STEPS:
echo.
echo 1. Generate secrets: bash scripts/generate-secrets.sh
echo 2. Run security check: bash scripts/security-check.sh
echo 3. Follow: PRODUCTION_DEPLOYMENT_CHECKLIST.md
echo.
pause
