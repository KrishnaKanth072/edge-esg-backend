@echo off
echo Fixing Go dependencies...
echo.

REM Delete old go.sum
del go.sum

REM Download and generate proper go.sum
go mod download
go mod tidy

echo.
echo Dependencies fixed!
echo.
echo Now committing and pushing...
git add go.mod go.sum
git commit -m "fix: regenerate go.sum with proper checksums"
git push origin main

echo.
echo ========================================
echo Fix pushed to GitHub!
echo ========================================
echo.
pause
