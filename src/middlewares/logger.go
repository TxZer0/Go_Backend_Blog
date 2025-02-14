package middlewares

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() *zap.Logger {
	if err := os.MkdirAll("./log", 0755); err != nil {
		log.Fatalln("Error creating log directory:", err)
	}

	logFile, err := os.OpenFile("./log/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	writer := zapcore.AddSync(logFile)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		writer,
		zapcore.InfoLevel,
	)
	return zap.New(core)
}

func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		duration := time.Since(start)

		logger.Info("Request",
			zap.String("Method", ctx.Request.Method),
			zap.String("Path", ctx.Request.URL.Path),
			zap.Int("Status", ctx.Writer.Status()),
			zap.String("Client_IP", ctx.ClientIP()),
			zap.Duration("Latency", duration),
		)
	}
}
