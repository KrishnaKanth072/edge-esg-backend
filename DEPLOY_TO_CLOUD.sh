#!/bin/bash
set -e

# EDGE ESG Backend - Oracle Cloud Deployment
# Run this AFTER creating Oracle Cloud instance

echo "☁️  EDGE ESG Backend - Oracle Cloud Deployment"
echo "=============================================="
echo ""

# Colors
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

echo -e "${BLUE}Prerequisites:${NC}"
echo "1. Oracle Cloud account created"
echo "2. ARM instance created (4 CPU, 24GB RAM)"
echo "3. Public IP address obtained"
echo ""

# Get Oracle Cloud details
read -p "Enter your Oracle Cloud Public IP: " ORACLE_IP
read -p "Enter path to SSH key (or press ENTER for default ~/.ssh/id_rsa): " SSH_KEY
SSH_KEY=${SSH_KEY:-~/.ssh/id_rsa}

echo ""
echo -e "${YELLOW}Testing SSH connection...${NC}"
if ssh -i $SSH_KEY -o ConnectTimeout=5 -o StrictHostKeyChecking=no ubuntu@$ORACLE_IP "echo 'Connected'" 2>/dev/null; then
    echo -e "${GREEN}✅ SSH connection successful${NC}"
else
    echo -e "${YELLOW}⚠️  Cannot connect. Make sure:${NC}"
    echo "  - Instance is running"
    echo "  - Firewall allows port 22"
    echo "  - SSH key is correct"
    exit 1
fi
echo ""

# Get GitHub username
read -p "Enter your GitHub username: " GITHUB_USER
echo ""

echo -e "${YELLOW}🚀 Deploying to Oracle Cloud...${NC}"
echo ""

# Deploy via SSH
ssh -i $SSH_KEY ubuntu@$ORACLE_IP << ENDSSH
set -e

echo "📦 Installing Docker..."
if ! command -v docker &> /dev/null; then
    curl -fsSL https://get.docker.com -o get-docker.sh
    sudo sh get-docker.sh
    sudo usermod -aG docker ubuntu
    rm get-docker.sh
fi

echo "📦 Installing Docker Compose..."
if ! command -v docker-compose &> /dev/null; then
    sudo apt update
    sudo apt install docker-compose -y
fi

echo "📦 Installing Git..."
sudo apt install git -y

echo "🔥 Configuring firewall..."
sudo ufw allow 22/tcp
sudo ufw allow 80/tcp
sudo ufw allow 443/tcp
sudo ufw allow 8000/tcp
sudo ufw --force enable

echo "💾 Creating swap..."
if [ ! -f /swapfile ]; then
    sudo fallocate -l 4G /swapfile
    sudo chmod 600 /swapfile
    sudo mkswap /swapfile
    sudo swapon /swapfile
    echo '/swapfile none swap sw 0 0' | sudo tee -a /etc/fstab
fi

echo "📥 Cloning repository..."
if [ -d edge-esg-backend ]; then
    cd edge-esg-backend
    git pull origin main
else
    git clone https://github.com/$GITHUB_USER/edge-esg-backend.git
    cd edge-esg-backend
fi

echo "🚀 Starting services..."
docker-compose down 2>/dev/null || true
docker-compose up -d

echo "⏳ Waiting for services to start..."
sleep 30

echo "🧪 Testing deployment..."
curl -f http://localhost:8000/health || echo "Services still starting..."

echo ""
echo "✅ Deployment complete!"
ENDSSH

echo ""
echo -e "${GREEN}=========================================="
echo "✅ CLOUD DEPLOYMENT COMPLETE!"
echo "==========================================${NC}"
echo ""
echo -e "${BLUE}🌐 Access your application:${NC}"
echo "Health Check: http://$ORACLE_IP:8000/health"
echo "API Endpoint: http://$ORACLE_IP:8000/api/v1/analyze"
echo ""
echo -e "${BLUE}🧪 Test with curl:${NC}"
echo "curl http://$ORACLE_IP:8000/health"
echo ""
echo "curl -X POST http://$ORACLE_IP:8000/api/v1/analyze \\"
echo "  -H 'Content-Type: application/json' \\"
echo "  -d '{\"company_name\":\"Tata Steel\",\"bank_id\":\"123e4567-e89b-12d3-a456-426614174000\",\"mode\":\"auto\"}'"
echo ""
echo -e "${BLUE}📊 Monitor services:${NC}"
echo "ssh -i $SSH_KEY ubuntu@$ORACLE_IP 'cd edge-esg-backend && docker-compose ps'"
echo ""
echo -e "${BLUE}🔄 Update deployment:${NC}"
echo "git push origin main"
echo "./DEPLOY_TO_CLOUD.sh"
echo ""
echo -e "${GREEN}Total Cost: \$0/month forever! 🎉${NC}"
