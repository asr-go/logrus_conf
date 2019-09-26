package middleware

import (
	"bytes"
	"io/ioutil"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GinMiddleware Gin Log 中间件
func GinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		printRequestLog(c, &startTime)

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()
		printResponseLog(c, &startTime, &endTime)
	}
}

// printRequestLog 输出Request日志
func printRequestLog(c *gin.Context, startTime *time.Time) {
	reqMethod := c.Request.Method              // 请求方式
	reqURI := c.Request.RequestURI             // 请求路由
	clientIP := c.ClientIP()                   // 请求IP
	contentType := c.GetHeader("content-type") // Content Type
	data, err := c.GetRawData()                // RawData

	if err != nil {
		logrus.Error("读取RequestBody失败！", err)
	}

	logrus.Info(">>>>>>>>>>>>>>>>>>>>>>>>>>> Request  >>>>>>>>>>>>>>>>>>>>>>>>>>>")
	logrus.Info("startTime   : ", startTime)
	logrus.Info("reqURI      : ", reqURI)
	logrus.Info("reqMethod   : ", reqMethod)
	logrus.Info("clientIP    : ", clientIP)
	logrus.Info("contentType : ", contentType)
	if len(data) > 0 && strings.Index(contentType, "json") > -1 {
		logrus.Infof("rawData     : \n%s", string(data))
	}
	logrus.Info("<<<<<<<<<<<<<<<<<<<<<<<<<<< Request  <<<<<<<<<<<<<<<<<<<<<<<<<<<")

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
}

// printResponseLog 输出Response日志
func printResponseLog(c *gin.Context, startTime *time.Time, endTime *time.Time) {
	// 执行时间
	latencyTime := endTime.Sub(*startTime)
	// 状态码
	statusCode := c.Writer.Status()
	// 返回数据

	logrus.Info(">>>>>>>>>>>>>>>>>>>>>>>>>>> Response >>>>>>>>>>>>>>>>>>>>>>>>>>>")
	logrus.Info("startTime   : ", startTime)
	logrus.Info("endTime     : ", endTime)
	logrus.Info("latencyTime : ", latencyTime)
	logrus.Info("statusCode  : ", statusCode)
	logrus.Info("<<<<<<<<<<<<<<<<<<<<<<<<<<< Response <<<<<<<<<<<<<<<<<<<<<<<<<<<")
}
