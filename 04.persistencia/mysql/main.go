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

	query := "create table fields1 (id bigint auto_increment primary key, bar varchar(255));"

	if _, err := db.Exec(query); err != nil {
		panic(err)
	}

	query = "insert into fields1 (item) values ('data value');"
	if _, err := db.Exec(query, "abcde"); err != nil {
		panic(err)
	}

	fmt.Println("Connected to MySQL database")
}
