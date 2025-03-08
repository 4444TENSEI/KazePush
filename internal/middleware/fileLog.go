package middleware

import (
	"KazePush/internal/config"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 接口访问日志中间件, 记录到本地日志文件
func FileLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		zapLogger := config.Logger
		// 记录请求开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 记录请求结束时间
		endTime := time.Now()
		// 计算请求耗时
		latency := endTime.Sub(startTime)
		// 记录访问信息
		zapLogger.Infow("访问信息",
			"方法", c.Request.Method,
			"路径", c.Request.URL.Path,
			"响应", c.Writer.Status(),
			"IP", c.ClientIP(),
			"HOST", c.Request.Host,
			"延迟", latency,
			"大小", fmt.Sprintf("%d Byte", c.Request.ContentLength),
		)
		// 如果有错误，记录错误信息
		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				zapLogger.Errorw("Error occurred",
					"method", c.Request.Method,
					"path", c.Request.URL.Path,
					"error", err.Error(),
				)
			}
		}
	}
}
