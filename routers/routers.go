package routers

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
	_ "net/http"
	apiControllerV1 "restful-api-kit/controllers/api/v1"
	apiControllerV2 "restful-api-kit/controllers/api/v2"
	google "restful-api-kit/middlewares/oauth"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		// add header Access-Control-Allow-Origin
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})

	v1 := r.Group("/api/v1")
	v2 := r.Group("/api/v2")
	v3 := r.Group("/")

	// V1
	v1.GET("/ping", apiControllerV1.Ping)
	v1.GET("/health", apiControllerV1.DbPing)
	v1.GET("/google", google.GoogleAuth)
	//v1.GET("/callback", google.GoogleCallback)
	v3.GET("/", google.GoogleCallback)

	// V2
	v2.GET("/ping", apiControllerV2.Ping)
	return r
}