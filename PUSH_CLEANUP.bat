@echo off
echo ========================================
echo Pushing professional repository updates
echo ========================================

cd /d "%~dp0"

echo Adding all changes...
git add -A

echo Committing changes...
git commit -m "fix: use lowercase repository names in Docker image tags for GHCR compliance

- Change IMAGE_PREFIX from uppercase to lowercase (krishnakanth072/edge)
- Fix Docker build errors: 'repository name must be lowercase'
- Remove unnecessary helper scripts and duplicate files
- Fix Slack notification errors in all workflows
- Add MIT License
- Add CONTRIBUTING.md with development guidelines
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
echo   - All builds should pass (Docker tags are lowercase)
echo   - Fully documented
echo.
pause
