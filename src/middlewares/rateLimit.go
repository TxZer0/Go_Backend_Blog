package middlewares

import (
	"net/http"
	"time"

	"github.com/TxZer0/Go_Backend_Blog/src/dto/response"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter"
	"github.com/ulule/limiter/drivers/store/memory"
)

func InitRateLimit() *limiter.Limiter {
	limiterRate := limiter.Rate{
		Period: time.Second * 60,
		Limit:  60,
	}

	store := memory.NewStore()
	rateLimiter := limiter.New(store, limiterRate)
	return rateLimiter
}

func RateLimitMiddleware(rateLimiter *limiter.Limiter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ipClient := ctx.ClientIP()
		limiterCtx, err := rateLimiter.Get(ctx, ipClient)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.NewInternalError())
			return
		}
		if limiterCtx.Reached {
			ctx.JSON(http.StatusTooManyRequests, response.NewTooManyRequests())
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
