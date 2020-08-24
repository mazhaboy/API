package controller

import (
	"encoding/json"
	"net/http"

	views "../views"
)

func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			data := views.Response{
				Code: http.StatusOK,
				Body: "Pong",
			}
			w.WriteHeader(201)
			json.NewEncoder(w).Encode(data)

		}
	}
}
