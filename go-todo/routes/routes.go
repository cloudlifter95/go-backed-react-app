package routes

import (
	"fmt"
	"go-todo/handlers"
	"go-todo/services"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetupRoutes(r *mux.Router, db *gorm.DB) {
	todoService := services.NewTodoService(db)

	todoHandler := handlers.NewTodoHandler(todoService)

	r.HandleFunc("/todos", todoHandler.GetTodos).Methods("GET")
	r.HandleFunc("/todos", todoHandler.CreateTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", todoHandler.GetTodoById).Methods("GET")
	r.HandleFunc("/todos/{id}", todoHandler.UpdateTodoById).Methods("PUT")
	r.HandleFunc("/todos/{id}", todoHandler.DeleteTodoById).Methods("DELETE")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "yops") }).Methods("GET")

}
