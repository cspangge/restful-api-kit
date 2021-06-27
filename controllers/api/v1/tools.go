package v1

import (
	"github.com/gin-gonic/gin"
	_ "log"
	"net/http"
	u "restful-api-kit/apiHelpers"
	"restful-api-kit/cache"
	"restful-api-kit/database"
	_ "restful-api-kit/utilities"
	tools "restful-api-kit/utilities"
	"time"
)

func Ping(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Pong",
	}
	u.RespondString(c.Writer, res)
}

func DbPing(c *gin.Context) {
	db := database.GetDB()
	sqlDB, _ := db.DB()
	var res map[string]interface{}
	if err := sqlDB.Ping(); err != nil {
		tools.CheckErr(err)
	}
	res = map[string]interface{}{
		"data": sqlDB.Stats(),
	}
	u.RespondString(c.Writer, res)
}

func GetRedis(c *gin.Context) {
	redis := cache.GetCache()
	res, err := redis.GetOne(c, "a")
	if err != nil {
		c.JSON(http.StatusOK, u.Resp(u.SUCCESS, err.Error()))
		return
	}
	c.JSON(http.StatusOK, u.Resp(u.SUCCESS, res))
	return
}

func SetRedis(c *gin.Context) {
	redis := cache.GetCache()
	res, err := redis.SetOne(c, "a", true, 10)
	if err != nil {
		c.JSON(http.StatusOK, u.Resp(u.FAIl_TO_GET_CACHE_RES, err))
		return
	}
	c.JSON(http.StatusOK, u.Resp(u.SUCCESS, res))
	return
}

func HSetRedis(c *gin.Context) {
	redis := cache.GetCache()
	res, err := redis.HashSet(c, "a", true)
	redis.Expire(c, "a", 5*time.Second)
	if err != nil {
		c.JSON(http.StatusOK, u.Resp(u.FAIl_TO_GET_CACHE_RES, err.Error()))
		return
	}
	c.JSON(http.StatusOK, u.Resp(u.SUCCESS, res))
	return
}

func HGetRedis(c *gin.Context) {
	redis := cache.GetCache()
	res, err := redis.HashGet(c, "a", "b")
	if err != nil {
		c.JSON(http.StatusOK, u.Resp(u.SUCCESS, err.Error()))
		return
	}
	c.JSON(http.StatusOK, u.Resp(u.SUCCESS, res))
	return
}
