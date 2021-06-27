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
	server := gin.Default()
	server.Use(middlewares.Cors())

	var (
		v1 = server.Group("/api/v1")
		v2 = server.Group("/api/v2")
	)

	v1.POST("/login", apiControllerV1.Login)
	v1.POST("/redis", apiControllerV1.TestRedis)

	v1.Use(middlewares.JWTAuth())
	// V1
	v1.GET("/ping", apiControllerV1.Ping)
	v1.GET("/health", apiControllerV1.DbPing)

	v2.GET("/ping", apiControllerV2.Ping)
	return server
}
