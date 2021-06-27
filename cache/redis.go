package cache

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"restful-api-kit/config"
	"strconv"
)

var Cache = make(map[string]*Redis)

type Redis struct {
	*redis.Client
}

func InitRedis() {
	redisConf := config.GlobalConfigSetting.Redis
	db, _ := strconv.Atoi(redisConf.Db)
	rdb := &Redis{
		redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", redisConf.Host, redisConf.Port),
			Password: redisConf.Password,
			DB:       db,
		}),
	}
	Cache["default"] = rdb
}

func GetCache(name ...string) *Redis {
	if len(name) == 0 {
		return Cache["default"]
	}
	return Cache[name[0]]
}

func (r *Redis) GetOne(c *gin.Context, key string) (string, error) {
	val, err := r.Get(c, key).Result()
	if err != nil {
		return "", errors.New(fmt.Sprintf("Redis err: %v", err))
	}
	return val, nil
}
