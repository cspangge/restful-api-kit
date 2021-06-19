package v2

import (
	"github.com/gin-gonic/gin"
	u "restful-api-kit/apiHelpers"
)

func Ping(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Pong",
	}
	u.RespondString(c.Writer, res)
}