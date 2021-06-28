package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	u "restful-api-kit/apiHelpers"
	"restful-api-kit/database"
	"restful-api-kit/middlewares"
	"restful-api-kit/models"
	"restful-api-kit/models/UserModel"
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

	userRecord, findErr := UserModel.GetUserByEmailPwd(db, req.Email, req.Password)
	if findErr != nil || userRecord.ID == 0 {
		c.JSON(http.StatusOK, u.Resp(u.FAIL_TO_LOGIN, findErr))
		return
	}

	if userRecord.Active == models.INACTIVE {
		c.JSON(http.StatusOK, u.Resp(u.ACTIVATE_FIRST))
		return
	}

	//if !models.Exist(db, "tbl_user", "email = ? and pwd = ? and active = ?", req.Email, req.Password, models.ACTIVE) {
	//	c.JSON(http.StatusOK, u.Resp(u.FAIL_TO_LOGIN))
	//	return
	//}

	res, err := middlewares.GenerateToken(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusOK, u.Resp(u.FAIL_TO_GET_TOKEN))
		return
	}

	c.JSON(http.StatusOK, u.Resp(u.SUCCESS, res))
	return
}
