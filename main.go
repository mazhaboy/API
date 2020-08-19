package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request received")
		fmt.Println(r.Method)
		w.Write([]byte("Hello World!"))
	})
	http.ListenAndServe("localhost:8181", mux)
}
