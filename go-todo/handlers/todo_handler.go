package handlers

import (
	// "encoding/json"

	"encoding/json"
	"fmt"
	"go-todo/models"
	"go-todo/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type TodoHandler struct {
	TodoService services.TodoServiceInterface
}

func NewTodoHandler(todoService *services.TodoService) *TodoHandler {
	return &TodoHandler{TodoService: todoService}
}

// GetTodos handles GET requests for fetching all todos
func (h *TodoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.TodoService.GetAllTodos()
	if err != nil {
		http.Error(w, "Error fetching todos", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	err := h.TodoService.CreateTodo(&todo)

	if err != nil {
		http.Error(w, "Error creating todo", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Todo created successfully")
}

func (h *TodoHandler) GetTodoById(w http.ResponseWriter, r *http.Request) {
	// log.Print(r)
	// fmt.Fprintf(w, "GetTodoByID\n")
	params := mux.Vars(r)
	log.Print(params)
	var todo models.Todo
	mytodo, err := h.TodoService.GetTodoById(&todo, params["id"])
	if err != nil {
		http.Error(w, "Error fetching Todo", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(mytodo)

	// return *todo, err
}

func (h *TodoHandler) UpdateTodoById(w http.ResponseWriter, r *http.Request) {
	// log.Print(r)
	fmt.Fprintf(w, "UpdateTodoById \n")
	params := mux.Vars(r)
	log.Print(params)

	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	log.Print(todo)

	err = h.TodoService.UpdateTodoById(&todo, params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) DeleteTodoById(w http.ResponseWriter, r *http.Request) {
	// log.Print(r)
	fmt.Fprintf(w, "DeleteTodoById - WIP \n")
	params := mux.Vars(r)
	log.Print(params)

	var todo models.Todo
	err := h.TodoService.DeleteTodoById(&todo, params["id"])
	if err != nil {
		http.Error(w, "Error deleting Todo", http.StatusInternalServerError)
		return
	}
	log.Print("deleted:", todo)

	json.NewEncoder(w).Encode(todo)
}
