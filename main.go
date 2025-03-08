package main

import (
	"KazePush/internal/config"
	"KazePush/internal/router"
	"log"
)

func main() {
	// 初始化日志记录器
	config.InitLogger()
	// 获取全局日志记录器
	zapLogger := config.Logger
	// 初始化项目配置文件
	config.InitConfig()
	// 启动服务
	r := router.RunServer()
	zapLogger.Infof("服务启动于端口: %s", config.GlobalCfg.Server.Port)
	log.Printf("服务启动于：http://localhost:%s", config.GlobalCfg.Server.Port)
	if err := r.Run(":" + config.GlobalCfg.Server.Port); err != nil {
		zapLogger.Errorf("XXX 启动服务出错: %v", err)
		log.Fatalln("XXX 启动服务出错", err)
	}
}
