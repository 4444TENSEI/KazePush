package middleware

import (
	"KazePush/internal/config"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type ipRecord struct {
	count     int
	expiresAt time.Time
}

var (
	ipRecords = make(map[string]*ipRecord)
	ipMutex   sync.RWMutex
	ticker    *time.Ticker
)

func init() {
	// 初始化自动清理过期记录的定时器
	ticker = time.NewTicker(time.Minute) // 假设每分钟检查一次
	go func() {
		for range ticker.C {
			clearExpiredIPRecords()
		}
	}()
}

// IP访问速率限制中间件
func IPRate() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		ipMutex.RLock()
		record, exists := ipRecords[ip]
		ipMutex.RUnlock()
		if exists && time.Now().Before(record.expiresAt) {
			if record.count >= config.GlobalCfg.Rate.MaxRequest {
				c.AbortWithStatusJSON(429, gin.H{
					"code":    429,
					"message": "请求过于频繁，请稍后再试",
				})
				return
			}
			record.count++
		} else {
			ipMutex.Lock()
			ipRecords[ip] = &ipRecord{
				count:     1,
				expiresAt: time.Now().Add(config.GlobalCfg.Rate.Duration),
			}
			ipMutex.Unlock()
		}
		c.Next()
	}
}

// 清理过期的IP记录
func clearExpiredIPRecords() {
	ipMutex.Lock()
	defer ipMutex.Unlock()
	for ip, record := range ipRecords {
		if time.Now().After(record.expiresAt) {
			delete(ipRecords, ip)
		}
	}
}

// 清空所有IP限制记录
func ClearIPRateCache() {
	ipMutex.Lock()
	defer ipMutex.Unlock()
	ipRecords = make(map[string]*ipRecord)
}
