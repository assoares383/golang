package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "user:password@/tests")

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		panic(err)
	}

	query := "create table usuarios (id bigint auto_increment primary key, bar varchar(255));"

	if _, err := db.Exec(query); err != nil {
		panic(err)
	}

	query = "insert into usuarios (bar) values (?);"
	if _, err := db.Exec(query, "abcde"); err != nil {
		panic(err)
	}

	query = "select * from usuarios limit 1;"
	type Usuario struct {
		id  int64
		bar string
	}

	var res Usuario
	if err := db.QueryRow(query).Scan(&res.id, &res.bar); err != nil {
		panic(err)
	}

	fmt.Println("%#+v\n", res)
}
