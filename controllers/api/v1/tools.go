package v1

import (
	"github.com/gin-gonic/gin"
	_ "log"
	"net/http"
	u "restful-api-kit/apiHelpers"
	"restful-api-kit/database"
	"restful-api-kit/middlewares"
	"restful-api-kit/models"
	_ "restful-api-kit/utilities"
	tools "restful-api-kit/utilities"
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

func Login(c *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, u.MakeRes(u.FAIL_TO_PARSE_PARAMETERS))
		return
	}

	var count int64
	db := database.GetDB()
	models.TblUserMgr(db.Where("email = ? and pwd = ?", req.Email, req.Password)).Count(&count)

	if count == 0 {
		c.JSON(http.StatusOK, u.MakeRes(u.FAIL_TO_LOGIN))
		return
	}

	res, err := middlewares.GenerateToken(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusOK, u.MakeRes(u.FAIL_TO_GET_TOKEN))
		return
	}

	c.JSON(http.StatusOK, u.MakeRes(u.SUCCESS, res))
	return
}
