package v1

import (
	"github.com/gin-gonic/gin"
	u "restful-api-kit/apiHelpers"
)

func Ping(c *gin.Context) {
	u.RespondString(c.Writer, "Pong")
}