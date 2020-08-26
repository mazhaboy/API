package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/mazha")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into dailycheklist (date, done) values (?, ?)",
		"25.08.2020", "UDEMY COURSES: WORKING WITH MYSQL")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(result.LastInsertId()) // id добавленного объекта
	fmt.Println(result.RowsAffected()) // количество затронутых строк
}

// create table products (id int auto_increment primary key, name v
