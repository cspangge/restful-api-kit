package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	tools "restful-api-kit/utilities"
)

type Database struct {
	Name string
	Host string
	Port string
	Username string
	Password string
	Database string
}

var dbList = make(map[string]*sql.DB)

func GetDb(database Database) *sql.DB {
	if dbList[database.Name] == nil {
		dbList[database.Name] = InitDb(database)
	}
	return  dbList[database.Name]
}

func InitDb(database Database) *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", database.Username, database.Password, database.Host, database.Port, database.Database))
	tools.CheckErr(err)
	log.Println("Database Connected")
	return db
}

func InitDbConnection(database Database) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", database.Username, database.Password, database.Host, database.Port, database.Database))
	tools.CheckErr(err)
	log.Println("Database Connected")
	dbList[database.Name] = db
}