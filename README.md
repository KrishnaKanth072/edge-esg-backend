# EDGE ESG Backend - RBI Compliant Trading Platform

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![CI/CD](https://github.com/KrishnaKanth072/edge-esg-backend/actions/workflows/pr-checks.yml/badge.svg)](https://github.com/KrishnaKanth072/edge-esg-backend/actions)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat&logo=docker)](https://hub.docker.com)
[![Kubernetes](https://img.shields.io/badge/Kubernetes-Ready-326CE5?style=flat&logo=kubernetes)](https://kubernetes.io)

> Enterprise-grade ESG trading platform with RBI compliance, quantum computing integration, and blockchain-backed audit trails. Deploy for FREE on Oracle Cloud.

## 🚀 Quick Start (2 Commands)

```bash
# 1. Deploy to GitHub
bash DEPLOY_MY_CODE.sh

# 2. Test locally
make up
```

Your repository: https://github.com/KrishnaKanth072/edge-esg-backend

---

# EDGE ESG Backend - RBI Compliant Trading Platform

## 🏦 RBI Compliance Features
- **Data Isolation**: Row-Level Security per bank_id
- **Encryption**: AES-256-GCM field-level encryption (pgcrypto)
- **Authentication**: Keycloak SSO with MFA (TOTP)
- **Audit Trail**: Blockchain-backed immutable logs
- **Rate Limiting**: 10K req/min per bank
- **Data Masking**: Role-based PII protection

## 🚀 Quick Start
```bash
make up      # Start all services
make test    # Run tests
make build   # Build Docker images
make deploy  # Deploy to k3s
```

## 🔐 Default Credentials
- **Keycloak**: admin/admin (http://localhost:8080)
- **PostgreSQL**: edge_admin/edge_secure_2024
- **Redis**: edge_redis_2024

## 📡 API Endpoints
- **Gateway**: http://localhost:8000
- **Health**: http://localhost:8000/health
- **Analyze**: POST http://localhost:8000/analyze
- **WebSocket**: ws://localhost:8000/ws

## 🏗 Architecture
10 Golang microservices communicating via gRPC:
1. Gateway (HTTP/gRPC)
2. Risk Agent (ESG scoring)
3. Trading Agent (BUY/SELL signals)
4. Quantum Agent (D-Wave simulation)
5. Compliance Agent (147 regulations)
6. Consensus Agent (9/10 voting)
7. Blockchain Agent (Polygon Mumbai)
8. Digital Twin Agent (3D factory)
9. Optimization Agent (Portfolio)
10. Regulation Agent (Zero-shot)

## 📊 Technology Stack
- **HTTP**: Gin v1.9.1
- **gRPC**: v1.60.1
- **Database**: PostgreSQL 16 + pgxpool
- **Cache**: Redis 7
- **Auth**: Keycloak (OIDC)
- **Logging**: zerolog
- **ORM**: GORM v1.25.7


## 🔄 CI/CD Pipeline

### GitHub Actions
- **Automated Testing**: Runs on every push/PR
- **Docker Build**: Builds all 10 microservices
- **Security Scanning**: Trivy + GoSec daily scans
- **Auto Deploy**: Staging (develop) + Production (main)

### GitLab CI/CD
- **3-Stage Pipeline**: test → build → deploy
- **Parallel Builds**: All 10 services build simultaneously
- **Manual Production**: Requires approval for prod deployment

### Deployment Commands
```bash
# Deploy to staging
./scripts/deploy.sh staging

# Deploy to production
./scripts/deploy.sh production

# Rollback gateway in production
./scripts/rollback.sh production gateway

# Scale trading agent
kubectl scale deployment/trading-agent --replicas=20 -n edge-production
```

### Required Secrets (GitHub/GitLab)
- `K3S_STAGING_URL` / `K3S_PROD_URL`
- `K3S_STAGING_TOKEN` / `K3S_PROD_TOKEN`
- `SLACK_WEBHOOK` (optional notifications)
- `GITHUB_TOKEN` (auto-provided by GitHub)

### Monitoring
- **Kubernetes HPA**: Auto-scales 3-20 replicas based on CPU/memory
- **Health Checks**: Liveness + Readiness probes
- **Resource Limits**: Memory 256Mi-512Mi, CPU 250m-500m


## 🌳 Git Workflow & Environments

### Branch Strategy
```
dev → main → release/v* (tags)
 ↓      ↓         ↓
Dev   Staging  Production
```

### Environments
| Environment | Branch | URL | Auto-Deploy | Replicas |
|-------------|--------|-----|-------------|----------|
| **Local** | any | localhost:8000 | No | 1 |
| **Dev** | dev | dev.edge.zebbank.com | ✅ Yes | 1 |
| **Staging** | main | staging.edge.zebbank.com | ✅ Yes | 2-3 |
| **Production** | v* tags | edge.zebbank.com | ⚠️ Manual | 3-20 |

### Quick Start with GitHub

1. **Push to GitHub**
```bash
git init
git add .
git commit -m "Initial commit: EDGE ESG Backend"
git branch -M main
git remote add origin https://github.com/YOUR_USERNAME/edge-esg-backend.git
git push -u origin main

# Create dev branch
git checkout -b dev
git push -u origin dev
```

2. **Configure GitHub Secrets**
Go to: Settings → Secrets and variables → Actions
- `KUBECONFIG_DEV` - Dev cluster config
- `KUBECONFIG_STAGING` - Staging cluster config
- `KUBECONFIG_PROD` - Production cluster config
- `SLACK_WEBHOOK` - Slack notifications (optional)

3. **Enable GitHub Actions**
- Go to Actions tab
- Enable workflows
- Push to `dev` branch → Auto-deploys to dev
- Push to `main` branch → Auto-deploys to staging
- Create tag `v1.0.0` → Manual approval for production

4. **Create Release**
```bash
git checkout main
git tag -a v1.0.0 -m "Release v1.0.0: Initial production release"
git push origin v1.0.0
```
Then approve deployment in GitHub Actions.

### Development Workflow
```bash
# Feature development
git checkout dev
git checkout -b feature/new-agent
# ... make changes ...
git commit -m "feat: add new ESG agent"
git push origin feature/new-agent
# Create PR to dev → Auto-deploys after merge

# Promote to staging
git checkout main
git merge dev
git push origin main
# Auto-deploys to staging

# Release to production
git tag -a v1.1.0 -m "Release v1.1.0"
git push origin v1.1.0
# Requires manual approval in GitHub Actions
```

See [BRANCHING_STRATEGY.md](BRANCHING_STRATEGY.md) for complete workflow details.


## 💰 100% FREE Deployment

Deploy to production for **$0/month** using Oracle Cloud Always Free tier:

```bash
# See complete guide
cat deploy/oracle-cloud-free/README.md

# Quick setup
./deploy/oracle-cloud-free/setup.sh
```

**What you get FREE:**
- 4 ARM CPUs + 24GB RAM
- 200GB storage
- 10TB bandwidth/month
- Public IP + SSL certificate
- Custom domain (DuckDNS)

**Total savings: $3,360/year** compared to AWS!

See [COST_BREAKDOWN.md](deploy/oracle-cloud-free/COST_BREAKDOWN.md) for details.


## 📝 Documentation

- [Contributing Guidelines](CONTRIBUTING.md)
- [Branch Protection Setup](SETUP_BRANCH_PROTECTION.md)
- [Oracle Cloud Deployment](deploy/oracle-cloud-free/README.md)

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🤝 Contributing

Contributions are welcome! Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

## 👨‍💻 Author

**Krishna Kanth** - [@KrishnaKanth072](https://github.com/KrishnaKanth072)

## ⭐ Show Your Support

Give a ⭐️ if this project helped you!

---

**Built with ❤️ for RBI-compliant ESG trading**
