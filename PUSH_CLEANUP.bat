@echo off
echo ========================================
echo Pushing cleanup and workflow fixes to GitHub
echo ========================================

cd /d "%~dp0"

echo Adding all changes...
git add -A

echo Committing changes...
git commit -m "chore: remove unnecessary files and fix Slack notification errors in workflows"

echo Pushing to GitHub...
git push origin main

echo ========================================
echo Changes pushed to GitHub!
echo ========================================
pause
