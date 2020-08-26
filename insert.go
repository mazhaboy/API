package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type a struct {
	id       int
	name     string
	surename string
	age      int
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/kv108")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select *from kv108")
	if err != nil {
		log.Fatal(err)
	}
	reads := []a{}
	for rows.Next() {
		data := a{}
		err := rows.Scan(&data.id, &data.name, &data.surename, &data.age)
		if err != nil {
			log.Fatal(err)
		}

		reads = append(reads, data)

	}
	// for _, row := range reads {
	// 	fmt.Println(row)
	// }
	fmt.Println(reads)

}

// create table products (id int auto_increment primary key, name v
