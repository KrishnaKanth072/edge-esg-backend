@echo off
echo ========================================
echo Pushing professional repository updates
echo ========================================

cd /d "%~dp0"

echo Adding all changes...
git add -A

echo Committing changes...
git commit -m "fix: update Go version to 1.23 across all Dockerfiles and workflows

- Update all 10 Dockerfiles from Go 1.21 to Go 1.23
- Update all GitHub Actions workflows to use Go 1.23
- Fix Docker build error: go.mod requires go >= 1.23
- Change IMAGE_PREFIX to lowercase (krishnakanth072/edge)
- Fix Slack notification errors in all workflows
- Add MIT License and CONTRIBUTING.md
- Improve .gitignore with comprehensive patterns
- Add professional badges to README
- Remove empty pkg/server directory
- Repository is now production-ready and showcase-ready"

echo Pushing to GitHub...
git push origin main

echo ========================================
echo Professional repository ready!
echo ========================================
echo.
echo Your repository is now:
echo   - Clean and professional
echo   - Ready to showcase to companies
echo   - All builds should pass (Go 1.23 + lowercase tags)
echo   - Fully documented
echo.
pause
