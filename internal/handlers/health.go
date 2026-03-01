package handlers

import (
	"net/http"
	"time"

	"github.com/edgeesg/edge-esg-backend/internal/dtos"
	"github.com/gin-gonic/gin"
)

var startTime = time.Now()

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, dtos.HealthResponse{
		Status:  "SERVING",
		Version: "1.0.0",
		Uptime:  int64(time.Since(startTime).Seconds()),
	})
}
