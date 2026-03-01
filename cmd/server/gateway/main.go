package main

import (
	"fmt"

	"github.com/edgeesg/edge-esg-backend/internal/config"
	"github.com/edgeesg/edge-esg-backend/internal/handlers"
	"github.com/edgeesg/edge-esg-backend/internal/loggers"
	"github.com/edgeesg/edge-esg-backend/internal/middleware"
	"github.com/edgeesg/edge-esg-backend/internal/services"
	"github.com/edgeesg/edge-esg-backend/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {
	loggers.Init()
	cfg, err := config.Load()
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	// Database connections
	db, err := database.NewGormDB(cfg.DatabaseURL)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	redisClient, err := database.NewRedisClient(cfg.RedisURL)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}

	// Initialize services
	orchestrator := services.NewOrchestrator()
	analyzeHandler := handlers.NewAnalyzeHandler(orchestrator)
	wsHub := handlers.NewWSHub()

	// Start WebSocket hub
	go wsHub.Run()

	// Setup Gin router
	r := gin.Default()
	r.Use(middleware.CORS())
	r.Use(middleware.DataMasking())

	// Rate limiter
	rateLimiter := middleware.NewRateLimiter(redisClient, 10000, 60)
	r.Use(rateLimiter.Limit())

	// Public routes
	r.GET("/health", handlers.HealthCheck)
	r.GET("/ws", wsHub.HandleWebSocket)

	// Protected routes (Keycloak auth disabled for demo)
	api := r.Group("/api/v1")
	{
		api.POST("/analyze", analyzeHandler.Analyze)
	}

	loggers.Info("Gateway server starting", map[string]interface{}{
		"port": cfg.ServerPort,
		"db":   db != nil,
	})

	if err := r.Run(":" + cfg.ServerPort); err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}
