package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"

	view "../views"

	model "../model"
)

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := view.PostRequest{}
		if r.Method == "POST" {
			json.NewDecoder(r.Body).Decode(&data)
			fmt.Println(data)
			if err := model.CreateToDo(data.Name, data.Todo); err !=  nil {
				w.Write([]byte("error"))
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		} else if r.Method == "GET" {
			name:=r.URL.Query().Get("name")
			data, err := model.ReadByName(name)
			if err != nil {
				w.Write([]byte("Some error"))
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		} else if r.Method=="DELETE"{
			name:= r.URL.Path[1:]
			if err:=model.DeleteToDo(name); err != nil {
				log.Fatal(err)
			}
			w.WriteHeader(http.StatusOK)
			
		}

	}

}
