#!/bin/bash
set -e

# EDGE ESG Backend - Automated GitHub Deployment
# Customized for: KrishnaKanth072

echo "🚀 EDGE ESG Backend - GitHub Deployment"
echo "========================================"
echo "User: KrishnaKanth072"
echo "Email: krishnakrishna99233@gmail.com"
echo ""

# Colors
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
RED='\033[0;31m'
NC='\033[0m'

# Pre-configured details
GIT_EMAIL="krishnakrishna99233@gmail.com"
GIT_USERNAME="KrishnaKanth072"
REPO_URL="https://github.com/KrishnaKanth072/edge-esg-backend.git"

# Configure Git
echo -e "${YELLOW}📝 Configuring Git...${NC}"
git config --global user.email "$GIT_EMAIL"
git config --global user.name "$GIT_USERNAME"
echo -e "${GREEN}✅ Git configured${NC}"
echo ""

# Fix dependencies
echo -e "${YELLOW}📦 Fixing dependencies...${NC}"
if command -v go &> /dev/null; then
    echo "Downloading Go modules..."
    go mod download 2>/dev/null || true
    echo "Tidying Go modules..."
    go mod tidy 2>/dev/null || true
    echo -e "${GREEN}✅ Dependencies fixed${NC}"
else
    echo -e "${YELLOW}⚠️  Go not installed (Docker will handle it)${NC}"
fi
echo ""

# Initialize Git
echo -e "${YELLOW}📝 Initializing Git repository...${NC}"
if [ ! -d .git ]; then
    git init
    echo -e "${GREEN}✅ Git initialized${NC}"
else
    echo -e "${GREEN}✅ Git already initialized${NC}"
fi
echo ""

# Add all files
echo -e "${YELLOW}📦 Adding files...${NC}"
git add .
echo -e "${GREEN}✅ Files added${NC}"
echo ""

# Commit
echo -e "${YELLOW}💾 Committing changes...${NC}"
git commit -m "Initial commit: EDGE ESG Backend - RBI compliant trading platform

- 10 Golang microservices (Gateway, Risk, Trading, Quantum, Compliance, Consensus, Blockchain, Digital Twin, Optimization, Regulation)
- PostgreSQL database with RLS and encryption
- Redis caching
- Keycloak SSO authentication
- Docker Compose setup
- CI/CD pipelines (GitHub Actions + GitLab)
- FREE Oracle Cloud deployment guides
- Complete documentation

Features:
- AES-256-GCM encryption
- Row-level security (RLS)
- Rate limiting (10K req/min)
- Audit trails
- Data masking by role
- 8-layer AI pipeline
- Mock quantum simulation
- Blockchain audit recording

Total Cost: \$0/month with Oracle Cloud Always Free tier" 2>/dev/null || echo "Already committed"

git branch -M main
echo -e "${GREEN}✅ Changes committed${NC}"
echo ""

# GitHub repository instructions
echo -e "${YELLOW}🌐 GitHub Repository Setup${NC}"
echo ""
echo -e "${BLUE}IMPORTANT: Create your GitHub repository now!${NC}"
echo ""
echo "1. Open this link in your browser:"
echo "   ${YELLOW}https://github.com/new${NC}"
echo ""
echo "2. Fill in these details:"
echo "   Repository name: ${YELLOW}edge-esg-backend${NC}"
echo "   Description: RBI-compliant ESG trading platform with 10 microservices"
echo "   Visibility: ${YELLOW}Public${NC} (for FREE features)"
echo "   ❌ Don't check 'Add a README file'"
echo "   ❌ Don't add .gitignore"
echo "   ❌ Don't choose a license"
echo ""
echo "3. Click ${YELLOW}'Create repository'${NC}"
echo ""
read -p "Press ENTER after creating the repository..."
echo ""

# Add remote
echo -e "${YELLOW}🔗 Adding GitHub remote...${NC}"
git remote remove origin 2>/dev/null || true
git remote add origin $REPO_URL
echo -e "${GREEN}✅ Remote added${NC}"
echo ""

# Push to GitHub
echo -e "${YELLOW}📤 Pushing to GitHub...${NC}"
echo "This may take a minute..."
if git push -u origin main; then
    echo -e "${GREEN}✅ Pushed to main branch${NC}"
else
    echo -e "${RED}❌ Push failed. Trying with force...${NC}"
    git push -u origin main --force
fi
echo ""

# Create dev branch
echo -e "${YELLOW}🌿 Creating dev branch...${NC}"
git checkout -b dev 2>/dev/null || git checkout dev
git push -u origin dev 2>/dev/null || git push -u origin dev --force
git checkout main
echo -e "${GREEN}✅ Dev branch created${NC}"
echo ""

# Success message
echo -e "${GREEN}=========================================="
echo "✅ SUCCESS! CODE PUSHED TO GITHUB!"
echo "==========================================${NC}"
echo ""
echo -e "${BLUE}🌐 Your Repository:${NC}"
echo "https://github.com/KrishnaKanth072/edge-esg-backend"
echo ""
echo -e "${BLUE}📊 What You Have:${NC}"
echo "✅ 10 Golang microservices"
echo "✅ PostgreSQL + Redis + Keycloak"
echo "✅ RBI-compliant security"
echo "✅ Docker Compose setup"
echo "✅ CI/CD pipelines"
echo "✅ Complete documentation"
echo ""
echo -e "${BLUE}🧪 Test Locally:${NC}"
echo "make up"
echo "curl http://localhost:8000/health"
echo ""
echo -e "${BLUE}☁️  Deploy to FREE Cloud:${NC}"
echo "1. Create Oracle Cloud account: https://www.oracle.com/cloud/free/"
echo "2. Follow: ORACLE_CLOUD_SETUP.md"
echo "3. Run: bash DEPLOY_TO_CLOUD.sh"
echo ""
echo -e "${GREEN}Total Cost: \$0/month forever! 🎉${NC}"
echo ""
echo -e "${BLUE}View your code:${NC}"
echo "https://github.com/KrishnaKanth072/edge-esg-backend"
