#!/bin/bash
set -e

ENV=${1:-staging}
NAMESPACE="edge-${ENV}"
SERVICE=${2:-gateway}

echo "⏪ Rolling back ${SERVICE} in ${ENV}..."

kubectl rollout undo deployment/${SERVICE} -n $NAMESPACE

echo "⏳ Waiting for rollback to complete..."
kubectl rollout status deployment/${SERVICE} -n $NAMESPACE --timeout=5m

echo "✅ Rollback completed successfully!"
