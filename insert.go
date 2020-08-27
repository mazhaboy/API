package main

import (
	"database/sql"


	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/qwerty")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Query("update todo set name= ? where name= ? ",  "qwerty", "Adlet")
	
	
	
}

// create table products (id int auto_increment primary key, name v
