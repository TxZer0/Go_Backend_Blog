package middlewares

import (
	"net/http"

	"github.com/TxZer0/Go_Backend_Blog/src/dto/response"
	"github.com/gin-gonic/gin"
)

func ErrorHandleMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				ctx.JSON(http.StatusInternalServerError, response.NewInternalError())
			}
		}()
		ctx.Next()
	}
}
