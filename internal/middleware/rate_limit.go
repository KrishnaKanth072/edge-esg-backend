package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/zebbank/edge-esg-backend/internal/error_codes"
)

type RateLimiter struct {
	client *redis.Client
	limit  int
	window time.Duration
}

func NewRateLimiter(redisClient *redis.Client, limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		client: redisClient,
		limit:  limit,
		window: window,
	}
}

func (rl *RateLimiter) Limit() gin.HandlerFunc {
	return func(c *gin.Context) {
		bankID := c.GetHeader("X-Bank-ID")
		if bankID == "" {
			bankID = c.Query("bank_id")
		}
		if bankID == "" {
			bankID = "default"
		}

		key := fmt.Sprintf("rate_limit:%s", bankID)
		ctx := context.Background()

		count, err := rl.client.Incr(ctx, key).Result()
		if err != nil {
			c.Next()
			return
		}

		if count == 1 {
			rl.client.Expire(ctx, key, rl.window)
		}

		if count > int64(rl.limit) {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"code":    error_codes.RateLimitExceeded,
				"message": fmt.Sprintf("Rate limit exceeded: %d requests per minute", rl.limit),
			})
			c.Abort()
			return
		}

		c.Header("X-RateLimit-Limit", fmt.Sprintf("%d", rl.limit))
		c.Header("X-RateLimit-Remaining", fmt.Sprintf("%d", rl.limit-int(count)))
		c.Next()
	}
}
