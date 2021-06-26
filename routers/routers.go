package routers

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
	_ "net/http"
	apiControllerV1 "restful-api-kit/controllers/api/v1"
	apiControllerV2 "restful-api-kit/controllers/api/v2"
	"restful-api-kit/middlewares"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(
		middlewares.Cors(),
		middlewares.JWTAuth(),
	)

	var (
		v1 = r.Group("/api/v1")
		v2 = r.Group("/api/v2")
	)

	// V1
	v1.GET("/ping", apiControllerV1.Ping)
	v1.GET("/health", apiControllerV1.DbPing)

	v1.POST("/login", apiControllerV1.GenToken)

	// V2
	v2.GET("/ping", apiControllerV2.Ping)
	return r
}