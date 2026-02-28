#!/bin/bash
set -e

echo "🚀 EDGE ESG Backend - Oracle Cloud FREE Setup"
echo "=============================================="

# Colors
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

# Update system
echo -e "${YELLOW}📦 Updating system...${NC}"
sudo apt update && sudo apt upgrade -y

# Install Docker
echo -e "${YELLOW}🐳 Installing Docker...${NC}"
if ! command -v docker &> /dev/null; then
    curl -fsSL https://get.docker.com -o get-docker.sh
    sudo sh get-docker.sh
    sudo usermod -aG docker $USER
    rm get-docker.sh
    echo -e "${GREEN}✅ Docker installed${NC}"
else
    echo -e "${GREEN}✅ Docker already installed${NC}"
fi

# Install Docker Compose
echo -e "${YELLOW}🔧 Installing Docker Compose...${NC}"
if ! command -v docker-compose &> /dev/null; then
    sudo apt install docker-compose -y
    echo -e "${GREEN}✅ Docker Compose installed${NC}"
else
    echo -e "${GREEN}✅ Docker Compose already installed${NC}"
fi

# Install Nginx
echo -e "${YELLOW}🌐 Installing Nginx...${NC}"
sudo apt install nginx -y
echo -e "${GREEN}✅ Nginx installed${NC}"

# Install Certbot
echo -e "${YELLOW}🔒 Installing Certbot (Let's Encrypt)...${NC}"
sudo apt install certbot python3-certbot-nginx -y
echo -e "${GREEN}✅ Certbot installed${NC}"

# Install Git
echo -e "${YELLOW}📚 Installing Git...${NC}"
sudo apt install git -y
echo -e "${GREEN}✅ Git installed${NC}"

# Configure firewall
echo -e "${YELLOW}🔥 Configuring firewall...${NC}"
sudo ufw allow 22/tcp
sudo ufw allow 80/tcp
sudo ufw allow 443/tcp
sudo ufw allow 8000/tcp
sudo ufw --force enable
echo -e "${GREEN}✅ Firewall configured${NC}"

# Create swap (4GB)
echo -e "${YELLOW}💾 Creating swap space...${NC}"
if [ ! -f /swapfile ]; then
    sudo fallocate -l 4G /swapfile
    sudo chmod 600 /swapfile
    sudo mkswap /swapfile
    sudo swapon /swapfile
    echo '/swapfile none swap sw 0 0' | sudo tee -a /etc/fstab
    echo -e "${GREEN}✅ Swap created (4GB)${NC}"
else
    echo -e "${GREEN}✅ Swap already exists${NC}"
fi

# Create directories
echo -e "${YELLOW}📁 Creating directories...${NC}"
mkdir -p ~/edge-esg-backend
mkdir -p ~/backups
echo -e "${GREEN}✅ Directories created${NC}"

# Setup backup cron
echo -e "${YELLOW}⏰ Setting up daily backups...${NC}"
(crontab -l 2>/dev/null; echo "0 2 * * * docker exec postgres pg_dump -U edge_admin edge_esg > ~/backups/db_\$(date +\%Y\%m\%d).sql") | crontab -
(crontab -l 2>/dev/null; echo "0 3 * * * find ~/backups -name 'db_*.sql' -mtime +7 -delete") | crontab -
echo -e "${GREEN}✅ Backup cron configured${NC}"

echo ""
echo -e "${GREEN}✅ Setup complete!${NC}"
echo ""
echo "Next steps:"
echo "1. Clone your repository:"
echo "   git clone https://github.com/YOUR_USERNAME/edge-esg-backend.git"
echo "   cd edge-esg-backend"
echo ""
echo "2. Start services:"
echo "   make up"
echo ""
echo "3. Setup domain (DuckDNS):"
echo "   Visit: https://www.duckdns.org"
echo ""
echo "4. Get SSL certificate:"
echo "   sudo certbot --nginx -d your-domain.duckdns.org"
echo ""
echo "5. Configure Nginx:"
echo "   sudo cp deploy/oracle-cloud-free/nginx.conf /etc/nginx/sites-available/edge-esg"
echo "   sudo ln -s /etc/nginx/sites-available/edge-esg /etc/nginx/sites-enabled/"
echo "   sudo nginx -t"
echo "   sudo systemctl restart nginx"
echo ""
echo "🎉 Your FREE production environment is ready!"
