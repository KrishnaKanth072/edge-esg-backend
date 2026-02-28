@echo off
echo Pushing final fixes to GitHub...
echo.

git add .
git commit -m "fix: remove empty gin.go file and unused variables"
git push origin main

echo.
echo ========================================
echo All fixes pushed to GitHub!
echo ========================================
echo.
echo CI/CD should pass now!
echo Check: https://github.com/KrishnaKanth072/edge-esg-backend/actions
echo.
pause
