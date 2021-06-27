package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	u "restful-api-kit/apiHelpers"
	"restful-api-kit/database"
	"restful-api-kit/middlewares"
	"restful-api-kit/models"
)

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var req LoginReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, u.Resp(u.FAIL_TO_PARSE_PARAMETERS))
		return
	}

	db := database.GetDB()
	if !models.Exist(db, "tbl_user", "email = ? and pwd = ? and active = ?", req.Email, req.Password, models.ACTIVE) {
		c.JSON(http.StatusOK, u.Resp(u.FAIL_TO_LOGIN))
		return
	}

	res, err := middlewares.GenerateToken(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusOK, u.Resp(u.FAIL_TO_GET_TOKEN))
		return
	}

	c.JSON(http.StatusOK, u.Resp(u.SUCCESS, res))
	return
}
