package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func LogResponseMiddleware() gin.HandlerFunc {
	// Open file
	filePath := "/Users/jutioncandrakirana/Documents/GitHub/enigma/GOLANG/golang-sinar-harapan-makmur-api"
	fileName := filePath + "/LOG_RESPONSE.txt"
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// Set up logrus
	logger := logrus.New()
	logger.SetOutput(file)

	return func(c *gin.Context) {
		c.Next()

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		responseBody := blw.body.String()
		responseLog := model.ResponseLog{
			StatusCode:   c.Writer.Status(),
			ResponseBody: responseBody,
		}

		// Set log level based on status code
		switch {
		case c.Writer.Status() >= 500:
			logger.Error(responseLog)
		case c.Writer.Status() >= 400:
			logger.Warn(responseLog)
		default:
			logger.Info(responseLog)
		}
	}
}
