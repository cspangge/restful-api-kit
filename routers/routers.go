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
	v1.POST("/signUp", apiControllerV1.SignUp)

	{
		v1.POST("/getRedis", apiControllerV1.GetRedis)
		v1.POST("/setRedis", apiControllerV1.SetRedis)
		v1.POST("/setHRedis", apiControllerV1.HSetRedis)
		v1.POST("/getHRedis", apiControllerV1.HGetRedis)
	}

	v1.Use(middlewares.JWTAuth())
	// V1SignUp
	v1.GET("/ping", apiControllerV1.Ping)
	v1.GET("/health", apiControllerV1.DbPing)

	v2.GET("/ping", apiControllerV2.Ping)
	return server
}
