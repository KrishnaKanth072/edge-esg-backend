@echo off
echo Fixing admin approval workflow...
echo.

git add .
git commit -m "fix: allow admin to bypass approval requirement on own PRs"
git push origin main

echo.
echo ========================================
echo Fix pushed!
echo ========================================
echo.
echo Now you can merge your own PRs without approval.
echo Others still need your approval.
echo.
pause
