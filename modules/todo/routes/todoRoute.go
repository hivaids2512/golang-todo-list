package routes

import (
	"../handlers"
	"../../../components/auth"
	"github.com/gorilla/mux"
	"net/http"
)

func InitTodoRoute(r *mux.Router, sub *mux.Router) {
	sub.HandleFunc("/todos/add", handlers.AddTodoHandler)
	sub.HandleFunc("/todos/update", handlers.UpdateTodoHandler)
	sub.HandleFunc("/todos/remove/{id}", handlers.RemoveTodoHandler)
	sub.HandleFunc("/todos/view/{id}", handlers.ViewTodoHandler)
	sub.HandleFunc("/todos/list/{userId}", handlers.ListTodoHandler)
	http.Handle("/api/todos/", auth.ShowMeYouToken(r))
}
