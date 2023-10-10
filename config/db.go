package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB() (*sql.DB, error) {
	connStr := "mysql -hcontainers-us-west-106.railway.app -uroot -pRphKHVPwvaPP6ZFCi5dU --port 6042 --protocol=TCP railway"
	db, err := sql.Open("mysql", connStr)
	return db, err
}
