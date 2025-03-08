package config

import (
	"time"

	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 全局配置变量
var GlobalCfg Config

// 配置文件结构
type Config struct {
	Server struct {
		Port  string `mapstructure:"port"`
		Debug bool   `mapstructure:"debug"`
	}
	Smtp struct {
		SenderName     string `mapstructure:"senderName"`
		SenderEmail    string `mapstructure:"senderEmail"`
		SenderPassword string `mapstructure:"senderPassword"`
		SmtpServer     string `mapstructure:"smtpServer"`
		SmtpPort       int    `mapstructure:"smtpPort"`
	}
	Secure struct {
		ParamToken string `mapstructure:"paramToken"`
	}
	Rate struct {
		MaxRequest int           `mapstructure:"maxRequest"`
		Duration   time.Duration `mapstructure:"duration"`
	}
}

// 载入配置文件
func InitConfig() error {
	zapLogger := Logger
	viper.SetConfigName("config") // 配置文件名
	viper.SetConfigType("toml")   // 设置配置文件类型
	viper.AddConfigPath("./")     // 配置文件路径
	if err := viper.ReadInConfig(); err != nil {
		zapLogger.Error("读取配置文件失败: %s", err)
		log.Fatalln("读取配置文件失败", err)
	}
	if err := viper.UnmarshalExact(&GlobalCfg); err != nil {
		zapLogger.Error("解析配置文件失败: %s", err)
		log.Fatalln("解析配置文件失败", err)
	}
	// 热重载
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(&GlobalCfg); err != nil {
			zapLogger.Error("配置文件更新失败: %s", err)
		} else {
			zapLogger.Info("配置文件更新成功")
		}
	})
	return nil
}
