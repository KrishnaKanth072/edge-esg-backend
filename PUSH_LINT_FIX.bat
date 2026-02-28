@echo off
echo Pushing lint fixes to GitHub...
echo.

git add .
git commit -m "fix: remove unused variables in orchestrator"
git push origin main

echo.
echo ========================================
echo Lint fixes pushed to GitHub!
echo ========================================
echo.
echo Check: https://github.com/KrishnaKanth072/edge-esg-backend/actions
echo.
pause
