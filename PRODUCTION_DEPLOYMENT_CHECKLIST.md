# Production Deployment Checklist

## 🚀 Pre-Deployment Security Checklist

Use this checklist before deploying to production. Check off each item as you complete it.

---

### Phase 1: Secrets Management (30 minutes)

- [ ] **Generate secure secrets**
  ```bash
  bash scripts/generate-secrets.sh
  ```

- [ ] **Store secrets in secrets manager**
  - [ ] AWS Secrets Manager configured
  - [ ] OR HashiCorp Vault configured
  - [ ] OR Kubernetes Secrets created

- [ ] **Delete local .env file**
  ```bash
  rm .env
  ```

- [ ] **Verify .env not in Git**
  ```bash
  git status | grep .env
  # Should return nothing
  ```

---

### Phase 2: Database Security (20 minutes)

- [ ] **Update PostgreSQL password**
  ```sql
  ALTER USER edge_admin WITH PASSWORD 'YOUR_GENERATED_PASSWORD';
  ```

- [ ] **Enable PostgreSQL SSL**
  ```bash
  # Edit postgresql.conf
  ssl = on
  ssl_cert_file = '/path/to/cert.pem'
  ssl_key_file = '/path/to/key.pem'
  ```

- [ ] **Restrict database access**
  ```bash
  # Edit pg_hba.conf
  hostssl all all 0.0.0.0/0 md5
  ```

- [ ] **Test database connection**
  ```bash
  psql "postgresql://edge_admin:PASSWORD@host:5432/edge_esg?sslmode=require"
  ```

---

### Phase 3: Redis Security (10 minutes)

- [ ] **Update Redis password**
  ```bash
  # Edit redis.conf
  requirepass YOUR_GENERATED_PASSWORD
  ```

- [ ] **Enable Redis TLS** (optional but recommended)
  ```bash
  tls-port 6380
  tls-cert-file /path/to/cert.pem
  tls-key-file /path/to/key.pem
  ```

- [ ] **Restart Redis**
  ```bash
  sudo systemctl restart redis
  ```

- [ ] **Test Redis connection**
  ```bash
  redis-cli -a YOUR_PASSWORD ping
  # Should return PONG
  ```

---

### Phase 4: TLS/HTTPS Setup (30 minutes)

- [ ] **Get SSL certificate**
  ```bash
  # Option 1: Let's Encrypt (recommended)
  sudo certbot certonly --standalone -d yourdomain.com
  
  # Option 2: Self-signed (dev only)
  openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
    -keyout server.key -out server.crt
  ```

- [ ] **Set TLS environment variables**
  ```bash
  TLS_ENABLED=true
  TLS_CERT_FILE=/etc/letsencrypt/live/yourdomain.com/fullchain.pem
  TLS_KEY_FILE=/etc/letsencrypt/live/yourdomain.com/privkey.pem
  ```

- [ ] **Test HTTPS**
  ```bash
  curl -I https://yourdomain.com
  # Should return 200 OK
  ```

- [ ] **Verify SSL certificate**
  ```bash
  openssl s_client -connect yourdomain.com:443 -servername yourdomain.com
  ```

---

### Phase 5: Firewall Configuration (15 minutes)

- [ ] **Configure firewall rules**
  ```bash
  sudo ufw allow 443/tcp   # HTTPS
  sudo ufw allow 22/tcp    # SSH
  sudo ufw deny 8000/tcp   # Block direct app access
  sudo ufw deny 5432/tcp   # Block direct DB access
  sudo ufw deny 6379/tcp   # Block direct Redis access
  sudo ufw enable
  ```

- [ ] **Verify firewall status**
  ```bash
  sudo ufw status verbose
  ```

- [ ] **Test blocked ports**
  ```bash
  curl http://yourdomain.com:8000
  # Should timeout or be refused
  ```

---

### Phase 6: Application Configuration (20 minutes)

- [ ] **Set environment to production**
  ```bash
  ENVIRONMENT=production
  ```

- [ ] **Configure Keycloak**
  - [ ] Keycloak URL updated
  - [ ] Realm configured
  - [ ] Client ID set
  - [ ] MFA enabled

- [ ] **Update application secrets**
  - [ ] ENCRYPTION_KEY set
  - [ ] BANK_KEY_* for each bank set
  - [ ] DATABASE_URL set
  - [ ] REDIS_URL set

- [ ] **Verify configuration**
  ```bash
  # Application should fail to start if secrets missing
  ./gateway
  ```

---

### Phase 7: Security Validation (15 minutes)

- [ ] **Run security check**
  ```bash
  bash scripts/security-check.sh
  # Should pass all checks
  ```

- [ ] **Check for hardcoded secrets**
  ```bash
  git grep -i "password.*=.*['\"]" -- "*.go"
  # Should return nothing
  ```

- [ ] **Verify security headers**
  ```bash
  curl -I https://yourdomain.com | grep -i "x-"
  # Should show security headers
  ```

- [ ] **Test rate limiting**
  ```bash
  for i in {1..100}; do curl https://yourdomain.com/health; done
  # Should eventually return 429 Too Many Requests
  ```

- [ ] **Test authentication**
  ```bash
  curl https://yourdomain.com/api/v1/analyze
  # Should return 401 Unauthorized
  ```

---

### Phase 8: Monitoring Setup (20 minutes)

- [ ] **Configure logging**
  - [ ] Log aggregation (ELK/Splunk)
  - [ ] Security event alerts
  - [ ] Failed auth alerts

- [ ] **Set up monitoring**
  - [ ] Prometheus/Grafana
  - [ ] Health check monitoring
  - [ ] Resource usage alerts

- [ ] **Configure backups**
  - [ ] Database backups (daily)
  - [ ] Backup encryption enabled
  - [ ] Backup restoration tested

---

### Phase 9: Final Checks (10 minutes)

- [ ] **Test complete flow**
  1. [ ] User can authenticate
  2. [ ] User can access authorized endpoints
  3. [ ] User cannot access unauthorized endpoints
  4. [ ] Data is encrypted in database
  5. [ ] Audit logs are created

- [ ] **Performance test**
  ```bash
  ab -n 1000 -c 10 https://yourdomain.com/health
  ```

- [ ] **Security scan**
  ```bash
  # Use external tool
  nmap -sV yourdomain.com
  ```

---

### Phase 10: Documentation (10 minutes)

- [ ] **Document deployment**
  - [ ] Server details
  - [ ] Secret locations
  - [ ] Access procedures
  - [ ] Rollback procedures

- [ ] **Create incident response plan**
  - [ ] Contact list
  - [ ] Escalation procedures
  - [ ] Recovery procedures

- [ ] **Train team**
  - [ ] Security procedures
  - [ ] Incident response
  - [ ] Access management

---

## 🎯 Deployment

Once all checks are complete:

```bash
# 1. Build Docker images
docker-compose build

# 2. Push to registry
docker-compose push

# 3. Deploy to production
kubectl apply -f deploy/k3s/
kubectl rollout status deployment/gateway -n edge-production

# 4. Verify deployment
curl https://yourdomain.com/health
```

---

## 🚨 Post-Deployment

### Immediate (First Hour):

- [ ] Monitor logs for errors
- [ ] Check all services are running
- [ ] Verify database connections
- [ ] Test critical endpoints
- [ ] Monitor resource usage

### First Day:

- [ ] Review security logs
- [ ] Check for failed auth attempts
- [ ] Monitor performance metrics
- [ ] Verify backups completed
- [ ] Test alerting system

### First Week:

- [ ] Security audit
- [ ] Performance optimization
- [ ] User feedback review
- [ ] Incident response drill
- [ ] Documentation updates

---

## 📊 Success Criteria

Deployment is successful when:

- ✅ All checklist items completed
- ✅ Security check passes
- ✅ All services healthy
- ✅ No critical errors in logs
- ✅ Authentication working
- ✅ Data encrypted
- ✅ Backups running
- ✅ Monitoring active
- ✅ Team trained

---

## 🆘 Rollback Procedure

If issues occur:

```bash
# 1. Rollback deployment
kubectl rollout undo deployment/gateway -n edge-production

# 2. Verify rollback
kubectl rollout status deployment/gateway -n edge-production

# 3. Investigate issue
kubectl logs -f deployment/gateway -n edge-production

# 4. Fix and redeploy
```

---

## 📞 Emergency Contacts

- **Security Team:** security@yourdomain.com
- **DevOps Team:** devops@yourdomain.com
- **On-Call:** +91-XXX-XXX-XXXX

---

**Checklist Version:** 1.0  
**Last Updated:** February 28, 2026  
**Next Review:** After first production deployment
