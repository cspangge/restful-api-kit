package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	u "restful-api-kit/apiHelpers"
	"restful-api-kit/cache"
	"restful-api-kit/database"
	"restful-api-kit/models"
)

func Activate(c *gin.Context) {
	activeToken := c.Query("activeToken")
	if activeToken == "" {
		c.JSON(http.StatusOK, u.Resp(u.NOT_ENOUGH_PARAMTERS))
		return
	}
	redis := cache.GetCache()
	userId, err := redis.GetOne(c, activeToken)
	if err != nil {
		c.JSON(http.StatusOK, u.Resp(u.FAIL_TO_ACTIVATE, err.Error()))
		return
	}

	var user models.TblUser
	db := database.GetDB()
	err = models.TblUserMgr(db.Model(&user).Where("user_id = ?", userId)).Update("active", 1).Error
	if err != nil {
		c.JSON(http.StatusOK, u.Resp(u.FAIL_TO_ACTIVATE, err.Error()))
	}
	c.JSON(http.StatusOK, u.Resp(u.SUCCESS))
}
