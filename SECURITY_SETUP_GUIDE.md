# Security Setup Guide - EDGE ESG Backend

## 🔒 Production Security Checklist

### Step 1: Generate Secure Secrets

```bash
# Run the secrets generator
bash scripts/generate-secrets.sh

# This generates:
# - Main encryption key (32 bytes)
# - Bank-specific encryption keys
# - Database password
# - Redis password
```

### Step 2: Store Secrets Securely

**Option A: AWS Secrets Manager (Recommended)**
```bash
# Install AWS CLI
aws configure

# Store secrets
aws secretsmanager create-secret \
    --name edge-esg/encryption-key \
    --secret-string "YOUR_GENERATED_KEY"

aws secretsmanager create-secret \
    --name edge-esg/database-url \
    --secret-string "postgres://edge_admin:PASSWORD@host:5432/edge_esg"
```

**Option B: HashiCorp Vault**
```bash
# Store in Vault
vault kv put secret/edge-esg \
    encryption_key="YOUR_KEY" \
    database_url="postgres://..."
```

**Option C: Kubernetes Secrets**
```bash
# Create Kubernetes secret
kubectl create secret generic edge-esg-secrets \
    --from-literal=encryption-key=YOUR_KEY \
    --from-literal=database-url=postgres://... \
    -n edge-production
```

### Step 3: Update Database Passwords

```sql
-- Connect to PostgreSQL
psql -U postgres

-- Change password
ALTER USER edge_admin WITH PASSWORD 'YOUR_GENERATED_PASSWORD';

-- Verify
\du edge_admin
```

### Step 4: Update Redis Password

```bash
# Edit redis.conf
sudo nano /etc/redis/redis.conf

# Add line:
requirepass YOUR_GENERATED_PASSWORD

# Restart Redis
sudo systemctl restart redis
```

### Step 5: Enable TLS/HTTPS

**Get SSL Certificate:**

```bash
# Option 1: Let's Encrypt (Free)
sudo certbot certonly --standalone -d yourdomain.com

# Option 2: Self-signed (Development only)
openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
    -keyout /etc/ssl/private/edge-esg.key \
    -out /etc/ssl/certs/edge-esg.crt
```

**Update .env:**
```bash
TLS_ENABLED=true
TLS_CERT_FILE=/etc/letsencrypt/live/yourdomain.com/fullchain.pem
TLS_KEY_FILE=/etc/letsencrypt/live/yourdomain.com/privkey.pem
```

### Step 6: Configure Firewall

```bash
# Allow only necessary ports
sudo ufw allow 443/tcp  # HTTPS
sudo ufw allow 22/tcp   # SSH
sudo ufw deny 8000/tcp  # Block direct app access
sudo ufw enable

# Verify
sudo ufw status
```

### Step 7: Set Environment Variables

**For Docker:**
```yaml
# docker-compose.yml
services:
  gateway:
    environment:
      - ENCRYPTION_KEY=${ENCRYPTION_KEY}
      - DATABASE_URL=${DATABASE_URL}
    env_file:
      - .env.production  # Never commit this file
```

**For Kubernetes:**
```yaml
# deployment.yaml
env:
  - name: ENCRYPTION_KEY
    valueFrom:
      secretKeyRef:
        name: edge-esg-secrets
        key: encryption-key
```

### Step 8: Enable Security Features

Update your main.go to include all security middleware:

```go
import (
    "github.com/edgeesg/edge-esg-backend/internal/middleware"
)

func main() {
    router := gin.Default()
    
    // Load config (validates required env vars)
    config, err := config.Load()
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to load config")
    }
    
    // Security middleware (in order)
    router.Use(middleware.HTTPSRedirect(config.Environment))
    router.Use(middleware.SecurityHeaders())
    router.Use(middleware.RequestSizeLimit(10 * 1024 * 1024)) // 10MB
    router.Use(middleware.InputValidation())
    router.Use(middleware.SecurityLogger())
    router.Use(middleware.RateLimit())
    router.Use(middleware.CORS())
    
    // Authentication
    keycloak, _ := middleware.NewKeycloakMiddleware(
        config.KeycloakURL, 
        config.KeycloakRealm,
    )
    
    // Protected routes
    api := router.Group("/api/v1")
    api.Use(keycloak.Authenticate())
    api.Use(middleware.DataMasking())
    {
        api.POST("/analyze", handlers.Analyze)
    }
    
    // Start server with TLS
    if config.TLSEnabled {
        router.RunTLS(":"+config.ServerPort, 
            config.TLSCertFile, 
            config.TLSKeyFile)
    } else {
        router.Run(":" + config.ServerPort)
    }
}
```

---

## 🔐 Security Best Practices

### 1. Secrets Management
- ✅ Never commit secrets to Git
- ✅ Use secrets manager (AWS/Vault/K8s)
- ✅ Rotate secrets every 90 days
- ✅ Use different secrets per environment
- ✅ Audit secret access

### 2. Database Security
- ✅ Use strong passwords (32+ characters)
- ✅ Enable SSL/TLS for connections
- ✅ Restrict network access
- ✅ Enable audit logging
- ✅ Regular backups (encrypted)

### 3. API Security
- ✅ Always use HTTPS in production
- ✅ Implement rate limiting
- ✅ Validate all inputs
- ✅ Use security headers
- ✅ Log security events

### 4. Authentication
- ✅ Use Keycloak SSO
- ✅ Enable MFA (TOTP)
- ✅ Short token expiration (15 min)
- ✅ Refresh token rotation
- ✅ Log all auth attempts

### 5. Monitoring
- ✅ Set up security alerts
- ✅ Monitor failed auth attempts
- ✅ Track unusual access patterns
- ✅ Regular security audits
- ✅ Incident response plan

---

## 🚨 Incident Response

### If Secrets Are Compromised:

1. **Immediate Actions:**
   ```bash
   # Rotate all secrets immediately
   bash scripts/generate-secrets.sh
   
   # Update all services
   kubectl rollout restart deployment -n edge-production
   
   # Revoke all active tokens
   # (via Keycloak admin console)
   ```

2. **Investigation:**
   - Check audit logs
   - Identify affected systems
   - Assess data exposure
   - Document timeline

3. **Notification:**
   - Notify security team
   - Inform affected users
   - Report to RBI (if required)

4. **Prevention:**
   - Review access controls
   - Update security policies
   - Conduct security training

---

## 📋 Security Audit Checklist

Run this checklist monthly:

```bash
# 1. Check for exposed secrets
git log -p | grep -i "password\|secret\|key"

# 2. Verify TLS is enabled
curl -I https://yourdomain.com

# 3. Test rate limiting
ab -n 1000 -c 10 https://yourdomain.com/api/v1/health

# 4. Check security headers
curl -I https://yourdomain.com | grep -i "security\|x-"

# 5. Verify authentication
curl https://yourdomain.com/api/v1/analyze
# Should return 401 Unauthorized

# 6. Check logs for suspicious activity
grep "AUTH_FAILURE" /var/log/edge-esg/security.log

# 7. Verify database encryption
psql -c "SELECT * FROM esg_scores LIMIT 1;"
# Encrypted fields should show BYTEA

# 8. Test input validation
curl -X POST https://yourdomain.com/api/v1/analyze \
    -d '{"company": "<script>alert(1)</script>"}'
# Should return 400 Bad Request
```

---

## 🎯 Compliance Requirements

### RBI Guidelines Compliance:

- ✅ Data encryption at rest (AES-256-GCM)
- ✅ Data encryption in transit (TLS 1.3)
- ✅ Multi-factor authentication
- ✅ Audit trail (immutable)
- ✅ Access controls (RBAC)
- ✅ Data masking
- ✅ Incident response plan
- ✅ Regular security audits
- ✅ Data retention policy
- ✅ Disaster recovery plan

---

## 📞 Support

For security issues:
- Email: security@yourdomain.com
- Emergency: +91-XXX-XXX-XXXX
- Bug Bounty: https://yourdomain.com/security

**Report vulnerabilities responsibly.**
