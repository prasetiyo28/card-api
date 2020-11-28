package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/prasetiyo28/card-api/config"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()
	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + config.GetConfig().DB_HOST + ")/" + conf.DB_NAME
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		panic("connectionString error...")
	}

	err = db.Ping()
	if err != nil {
		panic("DSN invalid")
	}
}

func CreateCon() *sql.DB {
	return db
}
