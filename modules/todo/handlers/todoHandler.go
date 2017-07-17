package handlers

import (
	"../models"
	"../services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func AddTodoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		decoder := json.NewDecoder(r.Body)
		var todo models.Todo
		_ = decoder.Decode(&todo)
		res := services.AddTodo(todo)
		jData, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	default:
		// Give an error message.
	}
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		decoder := json.NewDecoder(r.Body)
		var todo models.Todo
		_ = decoder.Decode(&todo)
		res := services.UpdateTodo(todo)
		jData, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	default:
		// Give an error message.
	}
}

func RemoveTodoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "DELETE":
		vars := mux.Vars(r)
		id := vars["id"]
		res := services.RemoveTodo(id)
		jData, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	default:
		// Give an error message.
	}
}

func ViewTodoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		vars := mux.Vars(r)
		id := vars["id"]
		res := services.GetTodoById(id)
		jData, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	default:
		// Give an error message.
	}
}

func ListTodoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		res := services.GetTodoList()
		jData, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	default:
		// Give an error message.
	}
}
