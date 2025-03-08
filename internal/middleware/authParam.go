package middleware

import (
	"KazePush/internal/config"

	"github.com/gin-gonic/gin"
)

// 检查URL中的token参数，是否匹配预设的token值
func AuthParam() gin.HandlerFunc {
	paramToken := config.GlobalCfg.Secure.ParamToken
	return func(c *gin.Context) {
		token := c.Query("token")
		if token != paramToken {
			c.JSON(403, gin.H{"code": 403, "message": "无权访问"})
			c.Abort()
			return
		}
		c.Next()
	}
}
