@echo off
echo ========================================
echo Pushing professional repository updates
echo ========================================

cd /d "%~dp0"

echo Adding all changes...
git add -A

echo Committing changes...
git commit -m "fix: update Docker image registry path to match repository name

- Change IMAGE_PREFIX to krishnakanth072/edge-esg-backend
- Fix GHCR push error: 'installation not allowed to Create organization package'
- Update Go version to 1.23 across all Dockerfiles and workflows
- Fix Docker build error: go.mod requires go >= 1.23
- Use lowercase repository names for Docker compliance
- Fix Slack notification errors in all workflows
- Add MIT License and CONTRIBUTING.md
- Improve .gitignore with comprehensive patterns
- Add professional badges to README
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
echo   - All builds should pass (correct GHCR path)
echo   - Fully documented
echo.
pause
