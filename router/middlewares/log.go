package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LoggerToFile() gin.HandlerFunc {
	log := logrus.New()
	log.SetLevel(logrus.InfoLevel)
	log.SetFormatter(&logrus.JSONFormatter{})
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		lastTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqUrl := c.Request.RequestURI
		statusCode := c.Writer.Status()
		errs := c.Errors
		IP := c.ClientIP()

		if errs != nil {
			log.WithFields(logrus.Fields{
				"code":         statusCode,
				"ip":           IP,
				"method":       reqMethod,
				"latency_time": lastTime,
				"url":          reqUrl,
			}).Error(errs)
		} else {
			log.WithFields(logrus.Fields{
				"code":         statusCode,
				"ip":           IP,
				"method":       reqMethod,
				"latency_time": lastTime,
				"url":          reqUrl,
			}).Info()
		}

	}
}
