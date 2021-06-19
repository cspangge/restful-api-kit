package main

import (
	_ "github.com/gin-gonic/gin"
	tools "restful-api-kit/Utilities"
	Routers "restful-api-kit/routers"
)

func main() {
	r := Routers.SetupRouter()

	port := tools.GetEnv("port")

	if port == "" {
		port = "7000" //localhost
	}

	err := r.Run(":" + port)
	tools.CheckErr(err)
}
