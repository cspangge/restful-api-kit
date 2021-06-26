package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"restful-api-kit/config"
	Routers "restful-api-kit/routers"
	tools "restful-api-kit/utilities"
)

func main() {
	InitServer()
	r := Routers.SetupRouter()
	port := tools.GetEnv("port")
	if port == "" {
		port = "8000" //localhost
	}
	err := r.RunTLS(":"+port, "./key/localhost.crt", "./key/localhost.key")
	tools.CheckErr(err)
}

func InitServer() {
	mode := tools.GetEnv("mode")
	switch mode {
	case "production":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
	config.Setup()
}
