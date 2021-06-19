package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	u "restful-api-kit/apiHelpers"
	db "restful-api-kit/database"
	tools "restful-api-kit/utilities"
)

func Ping(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Pong",
	}
	u.RespondString(c.Writer, res)
}

func ListTables(c *gin.Context) {
	dbContext := db.GetDb(db.Database{Name: "staging"})
	if dbContext == nil {
		log.Fatalln("Database Connection Error")
	}
	results, err := dbContext.Query("SHOW TABLES")
	tools.CheckErr(err)


	var slices []string
	for results.Next() {
		var table string
		err = results.Scan(&table)
		if err != nil {
			tools.CheckErr(err)
		}
		log.Println(table)
		slices = append(slices, table)
	}
	res := map[string]interface{}{
		"data": slices,
	}
	u.RespondString(c.Writer, res)
}