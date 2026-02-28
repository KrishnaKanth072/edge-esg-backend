# 100% FREE Deployment Guide - Oracle Cloud Always Free Tier

## What You Get (FREE Forever)
- ✅ 4 ARM CPUs (Ampere A1)
- ✅ 24GB RAM
- ✅ 200GB Block Storage
- ✅ 10TB Outbound Data Transfer/month
- ✅ Public IP Address
- ✅ No credit card charges (Always Free)

## Step 1: Create Oracle Cloud Account

1. Go to: https://www.oracle.com/cloud/free/
2. Click "Start for free"
3. Sign up (requires credit card for verification but won't charge)
4. Select "Always Free" resources only

## Step 2: Create ARM Instance

1. **Compute → Instances → Create Instance**
2. **Name**: edge-esg-backend
3. **Image**: Ubuntu 22.04 (ARM64)
4. **Shape**: VM.Standard.A1.Flex
   - OCPUs: 4
   - Memory: 24 GB
5. **Networking**: Create new VCN (default)
6. **Add SSH Key**: Upload your public key
7. **Boot Volume**: 200GB
8. Click "Create"

## Step 3: Configure Firewall

**Security Lists → Default Security List**

Add Ingress Rules:
- Port 22 (SSH) - Source: 0.0.0.0/0
- Port 80 (HTTP) - Source: 0.0.0.0/0
- Port 443 (HTTPS) - Source: 0.0.0.0/0
- Port 8000 (Gateway) - Source: 0.0.0.0/0

## Step 4: SSH into Instance

```bash
ssh ubuntu@<YOUR_PUBLIC_IP>
```

## Step 5: Install Docker (ARM64)

```bash
# Update system
sudo apt update && sudo apt upgrade -y

# Install Docker
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh

# Add user to docker group
sudo usermod -aG docker ubuntu
newgrp docker

# Install Docker Compose
sudo apt install docker-compose -y

# Verify
docker --version
docker-compose --version
```

## Step 6: Clone & Deploy

```bash
# Clone repository
git clone https://github.com/YOUR_USERNAME/edge-esg-backend.git
cd edge-esg-backend

# Start all services
make up

# Check status
docker-compose ps
```

## Step 7: FREE Domain Options

### Option 1: Free Subdomain (Easiest)
- **DuckDNS**: https://www.duckdns.org (FREE forever)
  - Get: `edge-esg.duckdns.org`
  - Setup: 2 minutes
  
- **FreeDNS**: https://freedns.afraid.org (FREE)
  - Multiple domain options
  
- **No-IP**: https://www.noip.com (FREE with monthly confirmation)

### Option 2: Free .tk/.ml/.ga Domain
- **Freenom**: https://www.freenom.com
  - FREE .tk, .ml, .ga, .cf, .gq domains
  - ⚠️ Renew every 12 months

### Option 3: GitHub Pages Custom Domain (FREE)
- Use GitHub Pages for frontend
- API on Oracle Cloud

## Step 8: Setup FREE SSL (Let's Encrypt)

```bash
# Install Certbot
sudo apt install certbot python3-certbot-nginx -y

# Install Nginx
sudo apt install nginx -y

# Get SSL certificate
sudo certbot --nginx -d edge-esg.duckdns.org

# Auto-renewal (already configured)
sudo certbot renew --dry-run
```

## Step 9: Configure Nginx Reverse Proxy

```bash
sudo nano /etc/nginx/sites-available/edge-esg
```

Paste this config (created in next file)

## Step 10: Access Your Application

- **HTTP**: http://edge-esg.duckdns.org
- **HTTPS**: https://edge-esg.duckdns.org
- **API**: https://edge-esg.duckdns.org/api/v1/analyze

## Total Cost: $0/month Forever! 🎉

## Monitoring (FREE)

- **UptimeRobot**: https://uptimerobot.com (50 monitors FREE)
- **Grafana Cloud**: https://grafana.com (FREE tier)
- **Better Stack**: https://betterstack.com (FREE tier)

## Backup Strategy (FREE)

```bash
# Backup database daily
0 2 * * * docker exec postgres pg_dump -U edge_admin edge_esg > /backup/db_$(date +\%Y\%m\%d).sql

# Keep last 7 days
find /backup -name "db_*.sql" -mtime +7 -delete
```

## Troubleshooting

### Services not starting?
```bash
# Check logs
docker-compose logs -f

# Restart services
make down
make up
```

### Out of memory?
```bash
# Check memory
free -h

# Add swap (FREE)
sudo fallocate -l 4G /swapfile
sudo chmod 600 /swapfile
sudo mkswap /swapfile
sudo swapon /swapfile
echo '/swapfile none swap sw 0 0' | sudo tee -a /etc/fstab
```

### Firewall issues?
```bash
# Check Ubuntu firewall
sudo ufw status

# Allow ports
sudo ufw allow 80
sudo ufw allow 443
sudo ufw allow 8000
```

## GitHub Actions for FREE Deployment

See `.github/workflows/oracle-free-deploy.yml` for automated deployments.

## Support

- Oracle Cloud Docs: https://docs.oracle.com/en-us/iaas/
- Community: https://community.oracle.com/
