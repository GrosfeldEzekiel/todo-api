package controller

import (
	"LearningGo/section4/todo-api/model"
	"LearningGo/section4/todo-api/views"
	"encoding/json"
	"fmt"
	"net/http"
)

func crud() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data := views.PostRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			fmt.Println(data)
			if err := model.CreateTodo(data.Name, data.Todo); err != nil {
				w.Write([]byte("Some Error"))
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)

		} else if r.Method == http.MethodGet {
			name := r.URL.Query().Get("name")
			if name != "" {
				data, err := model.ReadByName(name)
				if err != nil {
					w.Write([]byte(err.Error()))
				}
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(data)
			} else {
				data, err := model.ReadAll()
				if err != nil {
					w.Write([]byte(err.Error()))
				}
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(data)
			}
		} else if r.Method == http.MethodDelete {
			name := r.URL.Path[1:]
			if err := model.Delete(name); err != nil {
				w.Write([]byte("Some Error"))
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(struct {
				Message string `json:"message"`
			}{Message: "Item deleted"})
		}
	}
}