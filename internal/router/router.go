// 接口路由
package router

import (
	"KazePush/internal/api/cache"
	"KazePush/internal/api/email"
	"KazePush/internal/config"
	"KazePush/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RunServer() *gin.Engine {
	// 使用gin默认配置
	r := gin.Default()
	if config.GlobalCfg.Server.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	// 跨域中间件
	r.Use(middleware.CORS())
	// 请求/访问日志记录中间件
	r.Use(middleware.FileLog())
	// 公开可访问的静态资源目录, 不能定义为跟路由因为和其他路由会冲突, 故这里定义为非跟路由的端点
	r.Static("/_", "./public")
	// 跟路由302重定向到首页
	r.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/_")
	})
	// 在线接口文档
	r.GET("/doc", func(c *gin.Context) {
		c.Redirect(302, "https://apifox.com/apidoc/shared-2bbab0f7-9ab7-437d-8477-8a8e2899357d")
	})
	// 404响应
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": 404, "message": "资源不存在，但服务仍在运行~"})
	})
	// 发送接口组，套一层请求频率限制中间件
	sendAPI := r.Group("/send", middleware.IPRate())
	{
		// 发邮件
		sendAPI.POST("/email", email.SendEmail)
	}
	// 缓存处理接口组
	cacheAPI := r.Group("/cache", middleware.AuthParam())
	{
		// 清除缓存
		cacheAPI.GET("/clear", cache.ClearCache)
	}
	return r
}
