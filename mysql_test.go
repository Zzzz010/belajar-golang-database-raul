package golang_mysql_raul

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetDatabase()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO customer(id, name) VALUES('A002', 'Jayan')"

	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}
	fmt.Println("Data sudah dimasukkan")
}

func TestQuerySql(t *testing.T) {
	db := GetDatabase()
	defer db.Close()

	ctx := context.Background()
	script2 := "Select id, name from customer"

	rows, err := db.QueryContext(ctx, script2)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id :", id)
		fmt.Println("Name :", name)
	}
	defer rows.Close()
}
