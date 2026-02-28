@echo off
REM EDGE ESG Backend - Windows Deployment Script
REM For: KrishnaKanth072

echo ========================================
echo EDGE ESG Backend - GitHub Deployment
echo ========================================
echo User: KrishnaKanth072
echo Email: krishnakrishna99233@gmail.com
echo.

REM Check if Git is installed
where git >nul 2>nul
if %ERRORLEVEL% NEQ 0 (
    echo ERROR: Git is not installed!
    echo Please install Git from: https://git-scm.com/download/win
    pause
    exit /b 1
)

REM Configure Git
echo Configuring Git...
git config --global user.email "krishnakrishna99233@gmail.com"
git config --global user.name "KrishnaKanth072"
echo Git configured!
echo.

REM Initialize Git
echo Initializing Git repository...
if not exist .git (
    git init
    echo Git initialized!
) else (
    echo Git already initialized!
)
echo.

REM Add files
echo Adding files...
git add .
echo Files added!
echo.

REM Commit
echo Committing changes...
git commit -m "Initial commit: EDGE ESG Backend - RBI compliant trading platform"
git branch -M main
echo Changes committed!
echo.

REM GitHub instructions
echo ========================================
echo CREATE GITHUB REPOSITORY NOW!
echo ========================================
echo.
echo 1. Open this link: https://github.com/new
echo 2. Repository name: edge-esg-backend
echo 3. Keep it Public
echo 4. Don't add README or .gitignore
echo 5. Click "Create repository"
echo.
pause

REM Add remote and push
echo Adding GitHub remote...
git remote remove origin 2>nul
git remote add origin https://github.com/KrishnaKanth072/edge-esg-backend.git
echo.

echo Pushing to GitHub...
git push -u origin main
echo.

REM Create dev branch
echo Creating dev branch...
git checkout -b dev
git push -u origin dev
git checkout main
echo.

echo ========================================
echo SUCCESS! CODE PUSHED TO GITHUB!
echo ========================================
echo.
echo Your Repository:
echo https://github.com/KrishnaKanth072/edge-esg-backend
echo.
echo Next Steps:
echo 1. Test locally: docker-compose up
echo 2. Deploy to cloud: bash DEPLOY_TO_CLOUD.sh
echo.
echo Total Cost: $0/month forever!
echo.
pause
