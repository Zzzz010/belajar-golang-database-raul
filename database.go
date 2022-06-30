package golang_mysql_raul

import (
	"database/sql"
	"time"
)

func GetDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/golang_mysql_raul?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(1 * time.Hour)

	return db
}
