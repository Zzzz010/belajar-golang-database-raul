package golang_mysql_raul

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetDatabase()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO customer(id, name, email, balance, rating, birth_date, married) VALUES('A003', 'Radit', 'radit@gmail.com', 10000, 3.5, '2011-03-02', false)"

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
	script2 := "Select id, name, from customer"

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

func TestQuerySqlComplex(t *testing.T) {
	db := GetDatabase()
	defer db.Close()

	ctx := context.Background()
	script2 := "Select id, name, email, balance, rating, birth_date, married, created_at from customer"

	rows, err := db.QueryContext(ctx, script2)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float32
		var married bool
		var birth_date, created_at time.Time

		err = rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &married, &created_at)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id :", id)
		fmt.Println("Name :", name)
		if email.Valid {
			fmt.Println("Email :", email.String)
		}
		fmt.Println("Balance :", balance)
		fmt.Println("Rating :", rating)
		fmt.Println("Birth Date :", birth_date)
		fmt.Println("Married :", married)
		fmt.Println("Created Date :", created_at)

		//fmt.Println("Id :", id, "Name :", name, "Email :", email, "Balance :", balance, "Rating :", rating, "Birth Date :", birth_date, "Married :", married, "Created Date :", created_at)
	}
	defer rows.Close()
}

func TestQueryInjection(t *testing.T) {
	db := GetDatabase()
	defer db.Close()

	username := "admin'; #"
	password := "salah"

	ctx := context.Background()
	sqlQuery := "Select username from user where username = '" + username +
		"' and password = '" + password + "' Limit 1"

	rows, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
	defer rows.Close()
}

func TestQueryParameterQISafe(t *testing.T) {
	db := GetDatabase()
	defer db.Close()

	username := "raul"
	password := "zzzz"

	ctx := context.Background()
	sqlQuery := "Select username from user where username = ? AND password = ? Limit 1"

	rows, err := db.QueryContext(ctx, sqlQuery, username, password)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
	defer rows.Close()
}

func TestExecSqlParameter(t *testing.T) {
	db := GetDatabase()
	defer db.Close()

	username := "raul"
	password := "zzzz"

	ctx := context.Background()
	script := "INSERT INTO user(username, password) VALUES(? , ?)"

	_, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	fmt.Println("Data user sudah dimasukkan")
}

func TestAutoIncrement(t *testing.T) {
	db := GetDatabase()
	defer db.Close()

	email := "jayan@gmail.com"
	comment := "ini email jayan"

	ctx := context.Background()
	script := "INSERT INTO comments(email, comment) VALUES(? , ?)"

	result, err := db.ExecContext(ctx, script, email, comment)
	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("Last insert Id:", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetDatabase()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO comments(email, comment) VALUES(? , ?)"

	stmt, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := "kera" + strconv.Itoa(i) + "@gmail.com"
		comment := "Ini komen ke" + strconv.Itoa(i)
		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		lastInsertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment Id:", lastInsertId)
	}
}

func TestDatabaseTransaction(t *testing.T) {
	db := GetDatabase()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	sqlQuery := "INSERT INTO comments(email, comment) VALUES (?, ?)"
	for i := 0; i < 10; i++ {
		email := "gorilla" + strconv.Itoa(i) + "gmail.com"
		comment := "Komentar ke-" + strconv.Itoa(i)
		_, err := tx.ExecContext(ctx, sqlQuery, email, comment)
		if err != nil {
			panic(err)
		}
	}

	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}
