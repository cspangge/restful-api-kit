package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	db "restful-api-kit/database"
	Routers "restful-api-kit/routers"
	tools "restful-api-kit/utilities"
)

func main() {
	mode := tools.GetEnv("mode")

	if mode == "development" {
		gin.SetMode(gin.DebugMode)
		// Initial Database
		db.InitDbConnection(db.Database{
			Name: tools.GetEnv("db_staging_name"),
			Host: tools.GetEnv("db_staging_host"),
			Port: tools.GetEnv("db_staging_port"),
			Username: tools.GetEnv("db_staging_user"),
			Password: tools.GetEnv("db_staging_pass"),
			Database: tools.GetEnv("db_staging_database"),
		})
	} else if mode == "production" {
		gin.SetMode(gin.ReleaseMode)
		db.InitDbConnection(db.Database{
			Name: tools.GetEnv("db_production_name"),
			Host: tools.GetEnv("db_production_host"),
			Port: tools.GetEnv("db_production_port"),
			Username: tools.GetEnv("db_production_user"),
			Password: tools.GetEnv("db_production_pass"),
			Database: tools.GetEnv("db_production_database"),
		})
	} else {
		gin.SetMode(gin.DebugMode)
		// Initial Database
		db.InitDbConnection(db.Database{
			Name: tools.GetEnv("db_staging_name"),
			Host: tools.GetEnv("db_staging_host"),
			Port: tools.GetEnv("db_staging_port"),
			Username: tools.GetEnv("db_staging_user"),
			Password: tools.GetEnv("db_staging_pass"),
			Database: tools.GetEnv("db_staging_database"),
		})
	}

	r := Routers.SetupRouter()

	port := tools.GetEnv("port")

	if port == "" {
		port = "8000" //localhost
	}

	err := r.Run(":" + port)
	tools.CheckErr(err)
}
