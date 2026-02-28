package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/edgeesg/edge-esg-backend/internal/config"
	"github.com/edgeesg/edge-esg-backend/internal/handlers"
	"github.com/edgeesg/edge-esg-backend/internal/loggers"
	"github.com/edgeesg/edge-esg-backend/internal/middleware"
	"github.com/edgeesg/edge-esg-backend/internal/services"
	"github.com/edgeesg/edge-esg-backend/pkg/database"
)

func main() {
	loggers.Init()
	cfg := config.Load()

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

	// Setup Gin router
	r := gin.Default()
	r.Use(middleware.CORS())
	r.Use(middleware.DataMasking())

	// Rate limiter
	rateLimiter := middleware.NewRateLimiter(redisClient, 10000, 60)
	r.Use(rateLimiter.Limit())

	// Public routes
	r.GET("/health", handlers.HealthCheck)

	// Protected routes (Keycloak auth disabled for demo)
	api := r.Group("/api/v1")
	{
		api.POST("/analyze", analyzeHandler.Analyze)
	}

	loggers.Info("Gateway server starting", map[string]interface{}{
		"port": cfg.ServerPort,
		"db":   db != nil,
	})

	r.Run(":" + cfg.ServerPort)
}
