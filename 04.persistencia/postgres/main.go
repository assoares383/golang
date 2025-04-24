package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	urlExample := "postgres://pg:password@localhost:8541/tests"
	db, err := pgxpool.New(context.Background(), urlExample)

	if err != nil {
		panic(err)
	}

	if err := db.Ping(context.Background()); err != nil {
		panic(err)
	}

	query := "create table usuarios (id bigserial primary key, bar varchar(255));"

	if _, err := db.Exec(context.Background(), query); err != nil {
		panic(err)
	}

	query = "insert into usuarios (bar) values ($1);"
	if _, err := db.Exec(context.Background(), query, "abcde"); err != nil {
		panic(err)
	}

	query = "select * from usuarios limit 1;"
	type Usuario struct {
		id  int64
		bar string
	}

	var res Usuario
	if err := db.QueryRow(context.Background(), query).Scan(&res.id, &res.bar); err != nil {
		panic(err)
	}

	fmt.Println("%#+v\n", res)
}
