package cache

import (
	"figin/config"
	"fmt"

	"log"

	"github.com/go-redis/redis"
)

var cache *redis.Client

// GetCache 获取缓存
func GetCache() *redis.Client {
	return cache
}

// Setup 初始化
func Setup() {
	conf := config.GetConf().Cache
	if conf.CacheType == "redis" {
		// 连接redis
		client := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%v:%v", conf.Host, conf.Port),
			Password: conf.Password, // password set
			DB:       conf.DB,       // use default DB
		})
		// 心跳检测
		_, err := client.Ping().Result()
		if err != nil {
			log.Fatalf("connect redis fail: %v", err)
		} else {
			cache = client
		}
	}
}
