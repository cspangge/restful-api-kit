package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	u "restful-api-kit/apiHelpers"
	d "restful-api-kit/database"
	"restful-api-kit/models"
)

type SignUpReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignUp(c *gin.Context) {
	var (
		req = SignUpReq{}
		tbl = d.TABLE_CATALOG[d.USER_ID]
		db  = d.GetDB()
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, u.Resp(u.FAIL_TO_PARSE_PARAMETERS))
		return
	}
	if models.Exist(db, tbl, "email = ?", req.Email) {
		c.JSON(http.StatusOK, u.Resp(u.ACCOUNT_ALREADY_EXISTS))
		return
	}

	var (
		id   = models.GetMaxId(db, tbl, "user_id")
		user = models.TblUser{
			UserID: id + 1,
			Name:   req.Name,
			Email:  req.Email,
			Pwd:    req.Password,
		}
	)

	if err := db.Table(tbl).Create(&user).Error; err != nil {
		c.JSON(http.StatusOK, u.Resp(u.FAIl_TO_CREATE_USER, err.Error()))
		return
	}
	c.JSON(http.StatusOK, u.Resp(u.SUCCESS))
}
