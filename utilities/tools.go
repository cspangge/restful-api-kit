package utilities

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// GetEnv use godot package to load/read the .env file and
// return the value of the key
func GetEnv(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	CheckErr(err)
	return os.Getenv(key)
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}