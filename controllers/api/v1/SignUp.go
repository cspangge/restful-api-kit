package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	u "restful-api-kit/apiHelpers"
	"restful-api-kit/cache"
	d "restful-api-kit/database"
	"restful-api-kit/helpers"
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
		id            = models.GetMaxId(db, tbl, "user_id")
		redis         = cache.GetCache()
		uuid, uuidErr = helpers.GetUuid()
		user          = models.TblUser{
			UserID: id + 1,
			Name:   req.Name,
			Email:  req.Email,
			Pwd:    req.Password,
		}
	)

	if uuidErr != nil {
		c.JSON(http.StatusOK, u.Resp(u.FAIl_TO_CREATE_USER, uuidErr.Error()))
	}
	if _, setErr := redis.SetOne(c, uuid, user.UserID, 1*60); setErr != nil {
		c.JSON(http.StatusOK, u.Resp(u.FAIl_TO_CREATE_USER, setErr.Error()))
	}
	if err := db.Table(tbl).Create(&user).Error; err != nil {
		c.JSON(http.StatusOK, u.Resp(u.FAIl_TO_CREATE_USER, err.Error()))
		return
	}

	var (
		subject = "Sea cow tech invite you to activate your account"
		body    = fmt.Sprintf("Clink here: %s to activate your account",
			fmt.Sprintf("<a href='https://localhost:8000/api/v1/activate?activeToken=%s'>https://localhost:8000/api/v1/activate?activeToken=%s</a>", uuid, uuid),
		)
	)
	if err := helpers.SendEMail(subject, body, req.Email); err != nil {
		c.JSON(http.StatusOK, u.Resp(u.FAIl_TO_CREATE_USER, err.Error()))
		return
	}
	c.JSON(http.StatusOK, u.Resp(u.SUCCESS))
}
