@echo off
echo Pushing CI/CD fix to GitHub...
echo.

git add .
git commit -m "fix: remove go.work file and fix CI/CD pipeline"
git push origin main

echo.
echo ========================================
echo Fix pushed to GitHub!
echo ========================================
echo.
echo GitHub Actions will now run successfully.
echo Check: https://github.com/KrishnaKanth072/edge-esg-backend/actions
echo.
pause
