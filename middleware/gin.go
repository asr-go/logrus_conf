package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GinMiddleware Gin Log 中间件
func GinMiddleware(debug bool) gin.HandlerFunc {
	if debug {
		return func(c *gin.Context) {
			// 开始时间
			startTime := time.Now()
			printRequestLog(c)

			// 处理请求
			c.Next()

			// 结束时间
			endTime := time.Now()
			printResponseLog(c, &startTime, &endTime)
		}
	}
	return func(c *gin.Context) {
		c.Next()
	}
}

// printRequestLog 输出Request日志
func printRequestLog(c *gin.Context) {
	method := c.Request.Method         // 请求方式
	requestURI := c.Request.RequestURI // 请求路由

	logrus.Debug(">>>>>>>>>>>>>>>>>>>>>>>>>>> Request  >>>>>>>>>>>>>>>>>>>>>>>>>>>")
	logrus.Debug("request_uri  : ", requestURI)
	logrus.Debug("method       : ", method)
	logrus.Debug("<<<<<<<<<<<<<<<<<<<<<<<<<<< Request  <<<<<<<<<<<<<<<<<<<<<<<<<<<")
}

// printResponseLog 输出Response日志
func printResponseLog(c *gin.Context, startTime *time.Time, endTime *time.Time) {
	// 执行时间
	latencyTime := endTime.Sub(*startTime)

	logrus.Debug(">>>>>>>>>>>>>>>>>>>>>>>>>>> Response >>>>>>>>>>>>>>>>>>>>>>>>>>>")
	logrus.Debug("latencyTime  : ", latencyTime)
	logrus.Debug("<<<<<<<<<<<<<<<<<<<<<<<<<<< Response <<<<<<<<<<<<<<<<<<<<<<<<<<<")
}
