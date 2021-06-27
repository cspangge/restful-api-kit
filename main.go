package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"restful-api-kit/cache"
	"restful-api-kit/config"
	"restful-api-kit/helpers"
	Routers "restful-api-kit/routers"
	tools "restful-api-kit/utilities"
)

func main() {
	Initialize()
	if err := Routers.SetupRouter().RunTLS(
		fmt.Sprintf(":%d", GetPort()),
		"./key/localhost.crt",
		"./key/localhost.key"); err != nil {
		panic(err)
	}
}

func Initialize() {
	InitServer()
	config.Setup()
	helpers.InitMail()
	cache.InitRedis()
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
}

func GetPort() string {
	port := tools.GetEnv("port")
	if port == "" {
		port = "8000" //localhost
	}
	return port
}
