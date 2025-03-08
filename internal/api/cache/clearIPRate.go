// api/cache/handler.go

package cache

import (
	"KazePush/internal/middleware"

	"github.com/gin-gonic/gin"
)

// 清理IP限制记录
func ClearCache(c *gin.Context) {
	middleware.ClearIPRateCache()
	c.JSON(200, gin.H{
		"code":    200,
		"message": "IP限制记录已成功清除",
	})
}
