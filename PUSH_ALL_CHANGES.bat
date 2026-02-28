@echo off
echo ========================================
echo Pushing All Changes to GitHub
echo ========================================
echo.

echo Step 1: Adding all files...
git add .

echo Step 2: Committing changes...
git commit -m "feat: add branch protection and admin approval requirements

- Add CODEOWNERS file (requires @KrishnaKanth072 approval)
- Add branch protection setup guide
- Add admin approval workflow
- Remove empty gin.go file
- Fix unused variables in orchestrator"

echo Step 3: Pushing to GitHub...
git push origin main

echo.
echo ========================================
echo SUCCESS! All changes pushed!
echo ========================================
echo.
echo Next Steps:
echo 1. Go to: https://github.com/KrishnaKanth072/edge-esg-backend/settings/branches
echo 2. Follow the guide in: SETUP_BRANCH_PROTECTION.md
echo 3. Setup branch protection for 'main' branch
echo.
echo After setup:
echo - All builds must pass before merge
echo - You must approve all PRs
echo - No direct pushes to main allowed
echo.
pause
