package main

import (
	"log"
	"net/http"

	model "./model"

	_ "github.com/go-sql-driver/mysql"

	controller "./controller"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	log.Fatal(http.ListenAndServe(":4545", mux))
}
