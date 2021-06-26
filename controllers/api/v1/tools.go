package v1

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	_ "log"
	"net/http"
	u "restful-api-kit/apiHelpers"
	"restful-api-kit/database"
	"restful-api-kit/middlewares"
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
	var res  map[string]interface{}
	if err := sqlDB.Ping(); err != nil {
		tools.CheckErr(err)
	}
	res = map[string]interface{}{
		"data": sqlDB.Stats(),
	}
	u.RespondString(c.Writer, res)
}

func Login(c *gin.Context) {

}

func GenToken(c *gin.Context) {
	type loginReq struct {
		User string
		Password string
	}
	var login loginReq
	if err := c.ShouldBindJSON(&login); err != nil {
		log.Println(err)
	}

	res, _ := GenerateToken(login.User, login.Password)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"message": "success",
		"data" : res,
	})
}

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()  //当前时间
	expireTime := nowTime.Add(10 * time.Second)  //有效时间

	claims := middlewares.Claims{
		username,
		password,
		jwt.StandardClaims {
			ExpiresAt : expireTime.Unix(),
			Issuer : "tar",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(middlewares.JwtSecret)
	return token, err
}