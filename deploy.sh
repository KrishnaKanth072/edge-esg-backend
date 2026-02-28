#!/bin/bash

echo "=========================================="
echo "EDGE ESG Backend - Deploy to GitHub"
echo "=========================================="
echo ""

# Add all changes
git add -A

# Commit
git commit -m "feat: production-ready EDGE ESG Backend

- Enterprise microservices architecture
- RBI compliance features
- Comprehensive security hardening
- Docker and Kubernetes ready
- CI/CD pipelines configured"

# Push
git push origin main

echo ""
echo "=========================================="
echo "✅ Deployed to GitHub!"
echo "=========================================="
echo ""
echo "Repository: https://github.com/KrishnaKanth072/edge-esg-backend"
echo ""
echo "Next steps:"
echo "1. Deploy to Railway: https://railway.app"
echo "2. Or deploy to Render: https://render.com"
echo "3. Or use Kubernetes: kubectl apply -f deploy/k3s/"
echo ""
