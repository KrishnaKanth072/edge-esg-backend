# EDGE ESG Backend

[![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat&logo=docker)](https://hub.docker.com)
[![Kubernetes](https://img.shields.io/badge/Kubernetes-Ready-326CE5?style=flat&logo=kubernetes)](https://kubernetes.io)

Enterprise-grade ESG (Environmental, Social, Governance) trading platform with RBI compliance, microservices architecture, and comprehensive security features.

## Features

### Core Capabilities
- **10 Microservices Architecture** - Gateway, Risk, Trading, Quantum, Compliance, Consensus, Blockchain, Digital Twin, Optimization, Regulation agents
- **ESG Scoring & Analysis** - Real-time ESG score calculation and trading signals
- **RBI Compliance** - Row-level security, field-level encryption, audit trails
- **Security Hardened** - AES-256-GCM encryption, Keycloak SSO, rate limiting, input validation
- **Cloud Native** - Docker containers, Kubernetes manifests, CI/CD pipelines

### Technical Stack
- **Backend:** Go 1.23+
- **API:** RESTful HTTP + gRPC
- **Database:** PostgreSQL 16 with pgcrypto
- **Cache:** Redis 7
- **Auth:** Keycloak (OIDC/OAuth2)
- **Deployment:** Docker, Kubernetes, GitHub Actions

## Quick Start

### Prerequisites
- Docker Desktop
- Go 1.23+ (for local development)
- 8GB RAM minimum

### Run Locally
```bash
# Clone repository
git clone https://github.com/KrishnaKanth072/edge-esg-backend.git
cd edge-esg-backend

# Start all services
make up

# Wait 30 seconds for services to initialize
sleep 30

# Test health endpoint
curl http://localhost:8000/health

# Test API
curl -X POST http://localhost:8000/api/v1/analyze \
  -H "Content-Type: application/json" \
  -d '{"company_name":"Suzlon Energy","user_role":"TRADER"}'
```

### Expected Response
```json
{
  "esg_score": "4.2/10",
  "risk_action": "REJECT",
  "trading_signal": {
    "action": "BUY",
    "symbol": "SUZLON.NS",
    "target_price": "₹312",
    "confidence": 0.91
  },
  "processing_time_ms": 45,
  "audit_hash": "0x...",
  "masked_data": false
}
```

## API Endpoints

### Health Check
```bash
GET /health
```

### ESG Analysis
```bash
POST /api/v1/analyze
Content-Type: application/json

{
  "company_name": "string",
  "user_role": "TRADER|COMPLIANCE|ADMIN"
}
```

## Architecture

### Microservices
1. **Gateway** (`:8000`) - HTTP API, routing, authentication
2. **Risk Agent** (`:50051`) - Risk assessment and scoring
3. **Trading Agent** (`:50052`) - Trading signal generation
4. **Quantum Agent** (`:50053`) - Quantum computing simulation
5. **Compliance Agent** (`:50054`) - Regulatory compliance checks
6. **Consensus Agent** (`:50055`) - Multi-agent consensus
7. **Blockchain Agent** (`:50056`) - Audit trail recording
8. **Digital Twin Agent** (`:50057`) - 3D modeling simulation
9. **Optimization Agent** (`:50058`) - Portfolio optimization
10. **Regulation Agent** (`:50059`) - Regulation analysis

### Data Flow
```
Client → Gateway → Orchestrator → [10 Agents] → Consensus → Response
                                        ↓
                                   Database
                                   (Encrypted)
```

## Security Features

### RBI Compliance
- **Data Isolation:** Row-level security per bank_id
- **Encryption:** AES-256-GCM field-level encryption
- **Authentication:** Keycloak SSO with MFA support
- **Audit Trail:** Immutable blockchain-backed logs
- **Rate Limiting:** 10,000 requests/minute per bank
- **Data Masking:** Role-based PII protection

### Security Middleware
- **Security Headers:** XSS, Clickjacking, CSP protection
- **Input Validation:** SQL injection & XSS prevention
- **HTTPS/TLS:** Enforced in production
- **Security Logging:** All auth attempts and data access logged
- **Request Limits:** 10MB max body size

## Development

### Build
```bash
# Build all services
make build

# Run tests
make test

# Run specific service
go run cmd/server/gateway/main.go
```

### Environment Variables
```bash
DATABASE_URL=postgres://user:pass@host:5432/edge_esg
REDIS_URL=redis://:pass@host:6379/0
KEYCLOAK_URL=http://keycloak:8080
KEYCLOAK_REALM=edgeesg
SERVER_PORT=8000
ENCRYPTION_KEY=<64-hex-characters>
ENVIRONMENT=development|production
```

### Generate Secure Secrets
```bash
bash scripts/generate-secrets.sh
```

## Deployment

### Docker Compose (Local)
```bash
make up
```

### Kubernetes
```bash
kubectl apply -f deploy/k3s/namespace.yaml
kubectl apply -f deploy/k3s/
```

### Cloud Platforms
- **Railway:** Connect GitHub repo, auto-deploy
- **Render:** Deploy from GitHub, select Docker
- **Heroku:** `heroku stack:set container && git push heroku main`
- **Oracle Cloud:** See `deploy/oracle-cloud-free/README.md`

## CI/CD

### GitHub Actions Workflows
- **PR Checks:** Lint, test, build validation
- **Security Scan:** Daily vulnerability scanning
- **Dev Deploy:** Auto-deploy to dev on push to `dev` branch
- **Staging Deploy:** Auto-deploy to staging on push to `main`
- **Production Deploy:** Manual approval required

### Branch Strategy
- `main` - Production-ready code
- `dev` - Development branch
- `feature/*` - Feature branches

## Database Schema

### Tables
- `esg_scores` - ESG analysis results (encrypted)
- `audit_trails` - Immutable audit logs
- `trade_signals` - Trading recommendations

### Migrations
Located in `internal/migrations/`, auto-run on startup.

## Monitoring

### Health Checks
- Liveness: `/health`
- Readiness: Database + Redis connectivity

### Logs
```bash
# View gateway logs
docker-compose logs -f gateway

# View all logs
docker-compose logs -f
```

## Troubleshooting

### Gateway won't start
```bash
# Restart gateway (Keycloak may need time to initialize)
docker-compose restart gateway
```

### Port already in use
```bash
# Check what's using port 8000
netstat -ano | findstr :8000

# Or change port in docker-compose.yml
ports:
  - "8001:8000"
```

### Out of memory
Increase Docker Desktop memory to 8GB:
Settings → Resources → Memory → 8GB

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for development guidelines.

## License

This project is licensed under the MIT License - see [LICENSE](LICENSE) file.

## Support

- **Issues:** [GitHub Issues](https://github.com/KrishnaKanth072/edge-esg-backend/issues)
- **Documentation:** See `/docs` folder
- **Email:** krishnakrishna99233@gmail.com

## Acknowledgments

Built with modern cloud-native technologies for enterprise-grade ESG analysis and trading.

---

**Status:** Production-ready MVP  
**Version:** 1.0.0  
**Last Updated:** February 2026
