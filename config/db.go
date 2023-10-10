package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB() (*sql.DB, error) {
	dburl := "root:RphKHVPwvaPP6ZFCi5dU@tcp(containers-us-west-106.railway.app:6042)/railway"
	db, err := sql.Open("mysql", dburl)
	return db, err
}
