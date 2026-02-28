package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zebbank/edge-esg-backend/internal/dtos"
)

var startTime = time.Now()

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, dtos.HealthResponse{
		Status:  "SERVING",
		Version: "1.0.0",
		Uptime:  int64(time.Since(startTime).Seconds()),
	})
}
