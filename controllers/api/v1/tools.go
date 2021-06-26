package v1

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "log"
	"net/http"
	u "restful-api-kit/apiHelpers"
	"restful-api-kit/database"
	"restful-api-kit/middlewares"
	"restful-api-kit/model"
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

func Login(c *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1000,
			"message": "fail",
			"data":    fmt.Sprintf("ERORR: %v", err),
		})
		return
	}

	var count int64
	db := database.GetDB()
	model.TblUserMgr(db.Where("email = ? and pwd = ?", req.Email, req.Password)).Count(&count)

	if count == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "Failure",
			"data":    "Check your email and password, please",
		})
		return
	}

	res, _ := GenerateToken(req.Email, req.Password)
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    res,
	})
	return
}

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()                       //当前时间
	expireTime := nowTime.Add(30 * time.Second) //有效时间

	claims := middlewares.Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "tar",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(middlewares.JwtSecret)
	return token, err
}
