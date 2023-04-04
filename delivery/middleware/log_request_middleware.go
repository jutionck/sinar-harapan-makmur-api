package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"
)

func LogRequestMiddleware() gin.HandlerFunc {
	// Open file
	filePath := "/Users/jutioncandrakirana/Documents/GitHub/enigma/GOLANG/golang-sinar-harapan-makmur-api"
	fileName := filePath + "/LOG_REQUEST.txt"
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// Set up logrus
	logger := logrus.New()
	logger.SetOutput(file)

	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Since(startTime)

		entryLog := model.RequestLog{
			EndTime:      endTime,
			StatusCode:   c.Writer.Status(),
			ClientIP:     c.ClientIP(),
			Method:       c.Request.Method,
			RelativePath: c.Request.URL.Path,
			UserAgent:    c.Request.UserAgent(),
		}

		entryLogString := fmt.Sprintf("[LOG] : %v\n", entryLog)
		logger.Println(entryLogString)
	}
}
