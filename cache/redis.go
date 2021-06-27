package cache

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"restful-api-kit/config"
	"strconv"
)

var Cache = make(map[string]*redis.Client)

func InitRedis() {
	redisConf := config.GlobalConfigSetting.Redis
	db, _ := strconv.Atoi(redisConf.Db)
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisConf.Host, redisConf.Port),
		Password: redisConf.Password,
		DB:       db,
	})
	Cache["default"] = rdb
}

func GetCache(name ...string) *redis.Client {
	if len(name) == 0 {
		return Cache["default"]
	}
	return Cache[name[0]]
}
