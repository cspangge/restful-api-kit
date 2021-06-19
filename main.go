package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
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

	r := Routers.SetupRouter()

	port := tools.GetEnv("port")

	if port == "" {
		port = "8000" //localhost
	}

	err := r.Run(":" + port)
	tools.CheckErr(err)
}
