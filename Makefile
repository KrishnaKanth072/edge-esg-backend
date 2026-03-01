.PHONY: up down build test deploy clean setup security-check

setup:
	@echo "🔧 Setting up development environment..."
	git config core.hooksPath .githooks
	chmod +x .githooks/pre-commit
	@echo "✅ Git hooks installed"
	@echo "✅ Run 'make security-check' to verify your code"

security-check:
	@echo "🔍 Running security checks..."
	@if grep -r "password\s*=\s*['\"]" --include="*.go" --exclude-dir=vendor --exclude="*_test.go" .; then \
		echo "❌ Hardcoded passwords found!"; \
		exit 1; \
	fi
	@echo "✅ Security checks passed"

up:
	docker-compose up -d
	@echo "✅ All services running on localhost"

down:
	docker-compose down -v

build:
	@echo "🔨 Building 10 microservices..."
	docker build -t edge/gateway -f cmd/server/gateway/Dockerfile .
	docker build -t edge/risk-agent -f cmd/server/risk-agent/Dockerfile .
	docker build -t edge/trading-agent -f cmd/server/trading-agent/Dockerfile .
	docker build -t edge/quantum-agent -f cmd/server/quantum-agent/Dockerfile .
	docker build -t edge/compliance-agent -f cmd/server/compliance-agent/Dockerfile .
	docker build -t edge/consensus-agent -f cmd/server/consensus-agent/Dockerfile .
	docker build -t edge/blockchain-agent -f cmd/server/blockchain-agent/Dockerfile .
	docker build -t edge/digital-twin-agent -f cmd/server/digital-twin-agent/Dockerfile .
	docker build -t edge/optimization-agent -f cmd/server/optimization-agent/Dockerfile .
	docker build -t edge/regulation-agent -f cmd/server/regulation-agent/Dockerfile .

test:
	go mod download
	go test ./internal/... -cover -v

deploy:
	kubectl apply -f deploy/k3s/

clean:
	docker-compose down -v
	docker system prune -af
