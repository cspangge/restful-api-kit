package main

import (
	_ "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	Routers "restful-api-kit/routers"
)

// use godot package to load/read the .env file and
// return the value of the key
func getEnv(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	checkErr(err)
	return os.Getenv(key)
}

func main() {
	r := Routers.SetupRouter()

	port := getEnv("port")

	if port == "" {
		port = "7000" //localhost
	}

	err := r.Run(":" + port)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
