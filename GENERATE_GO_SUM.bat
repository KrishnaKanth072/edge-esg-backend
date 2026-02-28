@echo off
echo ========================================
echo Generating proper go.sum file
echo ========================================
echo.

REM Check if Go is installed
where go >nul 2>nul
if %ERRORLEVEL% NEQ 0 (
    echo ERROR: Go is not installed!
    echo Please install Go from: https://go.dev/dl/
    pause
    exit /b 1
)

echo Step 1: Cleaning old files...
if exist go.sum del go.sum

echo Step 2: Downloading dependencies...
go mod download

echo Step 3: Tidying go.mod and generating go.sum...
go mod tidy

echo Step 4: Verifying go.sum was created...
if exist go.sum (
    echo.
    echo ========================================
    echo SUCCESS! go.sum generated
    echo ========================================
    echo.
    echo File size:
    dir go.sum | find "go.sum"
    echo.
    echo Now pushing to GitHub...
    git add go.mod go.sum
    git commit -m "fix: generate proper go.sum with all dependencies"
    git push origin main
    echo.
    echo ========================================
    echo Pushed to GitHub!
    echo ========================================
) else (
    echo.
    echo ERROR: go.sum was not created!
    echo This might mean Go modules are not properly configured.
)

echo.
pause
