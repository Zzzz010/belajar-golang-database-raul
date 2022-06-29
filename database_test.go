package golang_mysql_raul

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/golang_mysql_raul")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
