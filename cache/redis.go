package cache

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"restful-api-kit/config"
	"strconv"
	"time"
)

var (
	Cache       = make(map[string]*Redis)
	SET_FAILURE = "Fail to set value in cache"
	SET_SUCCESS = "Value already set"
	GET_FAILURE = "Fail to get value from cache"
	NO_RESULT   = "No result"
)

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
		return NO_RESULT, errors.New(err.Error())
	}
	return val, nil
}

func (r *Redis) SetOne(c *gin.Context, key string, val interface{}, ttl ...time.Duration) (string, error) {
	err := r.Set(c, key, val, ttl[0]*time.Second).Err()
	if err != nil {
		return SET_FAILURE, errors.New(err.Error())
	}
	return SET_SUCCESS, nil
}

func (r *Redis) HashSet(c *gin.Context, key string, val ...interface{}) (string, error) {
	err := r.HSet(c, key, val...).Err()
	if err != nil {
		return SET_FAILURE, errors.New(err.Error())
	}
	return SET_SUCCESS, nil
}

func (r *Redis) HashGet(c *gin.Context, key, field string) (string, error) {
	cmd := r.HGet(c, key, field)
	if cmd.Err() != nil {
		return "", errors.New(cmd.Err().Error())
	}
	return cmd.Val(), nil
}
