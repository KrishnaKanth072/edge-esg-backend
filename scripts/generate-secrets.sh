#!/bin/bash

# EDGE ESG Backend - Secure Secrets Generator
# This script generates cryptographically secure secrets for production use

echo "=========================================="
echo "EDGE ESG Backend - Secrets Generator"
echo "=========================================="
echo ""
echo "⚠️  SECURITY WARNING:"
echo "   - Never commit generated secrets to Git"
echo "   - Store secrets in a secure secrets manager"
echo "   - Rotate secrets regularly"
echo ""

# Generate main encryption key
echo "Generating main encryption key..."
ENCRYPTION_KEY=$(openssl rand -hex 32)
echo "ENCRYPTION_KEY=$ENCRYPTION_KEY"
echo ""

# Generate bank-specific keys
echo "Generating bank-specific encryption keys..."
HDFC_KEY=$(openssl rand -hex 32)
ICICI_KEY=$(openssl rand -hex 32)
SBI_KEY=$(openssl rand -hex 32)

echo "BANK_KEY_hdfc-bank=$HDFC_KEY"
echo "BANK_KEY_icici-bank=$ICICI_KEY"
echo "BANK_KEY_sbi-bank=$SBI_KEY"
echo ""

# Generate database password
echo "Generating database password..."
DB_PASSWORD=$(openssl rand -base64 32 | tr -d "=+/" | cut -c1-32)
echo "DATABASE_PASSWORD=$DB_PASSWORD"
echo ""

# Generate Redis password
echo "Generating Redis password..."
REDIS_PASSWORD=$(openssl rand -base64 32 | tr -d "=+/" | cut -c1-32)
echo "REDIS_PASSWORD=$REDIS_PASSWORD"
echo ""

# Create .env file
echo "Creating .env file..."
cat > .env << EOF
# EDGE ESG Backend - Production Secrets
# Generated: $(date)
# ⚠️  DO NOT COMMIT THIS FILE TO GIT

# Database Configuration
DATABASE_URL=postgres://edge_admin:${DB_PASSWORD}@localhost:5432/edge_esg

# Redis Configuration
REDIS_URL=redis://:${REDIS_PASSWORD}@localhost:6379/0

# Keycloak SSO Configuration
KEYCLOAK_URL=https://keycloak.yourdomain.com
KEYCLOAK_REALM=edgeesg

# Server Configuration
SERVER_PORT=8000
ENVIRONMENT=production

# Encryption Keys
ENCRYPTION_KEY=${ENCRYPTION_KEY}

# Bank-Specific Encryption Keys
BANK_KEY_hdfc-bank=${HDFC_KEY}
BANK_KEY_icici-bank=${ICICI_KEY}
BANK_KEY_sbi-bank=${SBI_KEY}

# TLS/HTTPS Configuration
TLS_ENABLED=true
TLS_CERT_FILE=/etc/ssl/certs/edge-esg.crt
TLS_KEY_FILE=/etc/ssl/private/edge-esg.key
EOF

echo "=========================================="
echo "✅ Secrets generated successfully!"
echo "=========================================="
echo ""
echo "📋 Next Steps:"
echo "   1. Review the generated .env file"
echo "   2. Update database and Redis with new passwords"
echo "   3. Store secrets in your secrets manager:"
echo "      - AWS Secrets Manager"
echo "      - HashiCorp Vault"
echo "      - Azure Key Vault"
echo "   4. Delete .env file after storing in secrets manager"
echo "   5. Never commit .env to Git"
echo ""
echo "🔒 Security Checklist:"
echo "   [ ] Secrets stored in secrets manager"
echo "   [ ] .env file deleted from local system"
echo "   [ ] Database password updated"
echo "   [ ] Redis password updated"
echo "   [ ] TLS certificates installed"
echo "   [ ] Firewall rules configured"
echo "   [ ] Backup encryption enabled"
echo ""
