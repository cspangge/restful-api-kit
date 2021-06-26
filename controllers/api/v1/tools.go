package v1

import (
	"github.com/gin-gonic/gin"
	_ "log"
	u "restful-api-kit/apiHelpers"
	"restful-api-kit/database"
	_ "restful-api-kit/utilities"
	tools "restful-api-kit/utilities"
)

func Ping(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Pong",
	}
	u.RespondString(c.Writer, res)
}

func Error(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Login Failed",
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