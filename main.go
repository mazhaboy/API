package main

import (
	"net/http"

	controller "./controller"
)

func main() {
	mux := controller.Register()
	http.ListenAndServe(":3000", mux)
}
