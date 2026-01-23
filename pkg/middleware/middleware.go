package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func Logger() gin.HandlerFunc {
	logger := logrus.New()
	logger.Out = os.Stdout
	logger.SetFormatter(&logrus.TextFormatter{})
	logger.SetLevel(logrus.DebugLevel)

	return func(c *gin.Context) {
		baseRoute := "/"
		if c.Request.RequestURI == baseRoute {
			c.Next()
			return
		}

		startTime := time.Now()
		c.Next()
		endTime := time.Now()

		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqURI := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		userIdentifier := "-"
		userID := "-"
		currentTime := time.Now().Format("02/Jan/2006:15:04:05 -0700")
		responseSize := c.Writer.Size()

		logEntry := fmt.Sprintf("%s %s %s [%s] \"%s %s HTTP/1.0\" %d %d %v",
			clientIP, userIdentifier, userID, currentTime, reqMethod, reqURI, statusCode, responseSize, latencyTime)

		logger.Log(logrus.DebugLevel, logEntry)
	}
}
