#!/bin/bash
set -e

ENV=${1:-staging}
NAMESPACE="edge-${ENV}"

echo "🚀 Deploying EDGE ESG Backend to ${ENV}..."

# Apply namespace
kubectl apply -f deploy/k3s/namespace.yaml

# Apply secrets (ensure they're configured)
kubectl apply -f deploy/k3s/secrets.yaml

# Deploy all services
kubectl apply -f deploy/k3s/gateway-deployment.yaml -n $NAMESPACE
kubectl apply -f deploy/k3s/agents-deployment.yaml -n $NAMESPACE
kubectl apply -f deploy/k3s/hpa.yaml -n $NAMESPACE

# Wait for rollout
echo "⏳ Waiting for deployments to complete..."
kubectl rollout status deployment/gateway -n $NAMESPACE --timeout=5m
kubectl rollout status deployment/risk-agent -n $NAMESPACE --timeout=5m
kubectl rollout status deployment/trading-agent -n $NAMESPACE --timeout=5m

echo "✅ Deployment to ${ENV} completed successfully!"
echo "🔗 Gateway URL: $(kubectl get svc gateway -n $NAMESPACE -o jsonpath='{.status.loadBalancer.ingress[0].ip}')"
