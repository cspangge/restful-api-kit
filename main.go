package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"restful-api-kit/config"
	Routers "restful-api-kit/routers"
	tools "restful-api-kit/utilities"
)

func main() {
	mode := tools.GetEnv("mode")

	if mode == "development" {
		gin.SetMode(gin.DebugMode)
	} else if mode == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	config.Setup()

	r := Routers.SetupRouter()

	port := tools.GetEnv("port")

	if port == "" {
		port = "8000" //localhost
	}

	TlsHandler()

	//err := r.Run(":" + port)
	err := r.RunTLS(":" + port, "./key/server.crt", "./key/server.key")
	tools.CheckErr(err)
}


func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:8000",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}
		c.Next()
	}
}