.PHONY: up down build test deploy clean

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
	go test ./internal/... -cover -v

deploy:
	kubectl apply -f deploy/k3s/

clean:
	docker-compose down -v
	docker system prune -af
