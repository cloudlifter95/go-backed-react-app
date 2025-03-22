package handlers

import (
	"bytes"
	"encoding/json"
	"go-todo/handlers"
	"go-todo/mocks"
	"go-todo/models"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetTodos(t *testing.T) {

	// Create a mock service
	mockService := new(mocks.TodoServiceInterface)

	// Define expected mock behavior
	mockTodos := []models.Todo{
		{Title: "Test Todo", Completed: false},
		{Title: "Another Todo", Completed: true},
	}
	mockService.On("GetAllTodos").Return(mockTodos, nil)

	// Inject the mock into the handler
	handler := handlers.TodoHandler{TodoService: mockService}

	// Call the handler (fake request, fake response)
	r, _ := http.NewRequest("GET", "/todos", nil)
	rr := httptest.NewRecorder()

	// Log request details
	var requestBody bytes.Buffer
	if r.Body != nil {
		bodyBytes, _ := io.ReadAll(r.Body)
		requestBody.Write(bodyBytes)
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // Reset body for next handler
	}
	log.Printf("Request: %s %s | Body: %s", r.Method, r.URL.Path, requestBody.String())

	// Capture the start time for the duration calculation
	start := time.Now()

	// Call the handler
	handler.GetTodos(rr, r)

	// Calculate the duration
	duration := time.Since(start)

	// Log response details
	log.Printf("Response: %d | Duration: %v | Body: %s\n", rr.Code, duration, rr.Body.String())

	// Assertions
	require.Equal(t, http.StatusOK, rr.Code) // Ensure status code is OK
	var todos []models.Todo
	err := json.NewDecoder(rr.Body).Decode(&todos) // Decode the response body
	require.NoError(t, err)                        // Ensure no error occurred
	assert.Equal(t, 2, len(todos))                 // Ensure correct number of todos
	assert.Equal(t, "Test Todo", todos[0].Title)   // Ensure correct title
}

func TestCreateTodo_Success(t *testing.T) {
	// Create a mock TodoService
	mockService := new(mocks.TodoServiceInterface)

	// Define the mock behavior for CreateTodo
	mockService.On("CreateTodo", mock.AnythingOfType("*models.Todo")).Return(nil)

	// Create a TodoHandler with the mock service
	handler := handlers.TodoHandler{TodoService: mockService}

	// Define the request payload (new Todo)
	todo := &models.Todo{
		Title:     "Test Todo",
		Completed: false,
	}

	todoData, _ := json.Marshal(todo)

	// Create a new HTTP POST request
	r, err := http.NewRequest("POST", "/todos", bytes.NewBuffer(todoData))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	start := time.Now()

	// Log request details
	var requestBody bytes.Buffer
	if r.Body != nil {
		bodyBytes, _ := io.ReadAll(r.Body)
		requestBody.Write(bodyBytes)
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // Reset body for next handler
	}
	log.Printf("Request: %s %s | Body: %s", r.Method, r.URL.Path, requestBody.String())

	// Call the CreateTodo handler
	handler.CreateTodo(rr, r)

	duration := time.Since(start)

	// Log response details
	log.Printf("Response: %d | Duration: %v | Body: %s\n", rr.Code, duration, rr.Body.String())

	// Assert the response status code and body
	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, "Todo created successfully", rr.Body.String())

	// Assert that the mock service's CreateTodo method was called
	mockService.AssertExpectations(t)
}

func TestCreateTodo_InvalidInput(t *testing.T) {
	// Create a mock TodoService
	mockService := new(mocks.TodoServiceInterface)

	// Create a TodoHandler with the mock service
	handler := handlers.TodoHandler{TodoService: mockService}

	// Create an invalid request payload (e.g., empty body or invalid JSON)
	invalidData := []byte("invalid json")

	// Create a new HTTP POST request with invalid data
	r, err := http.NewRequest("POST", "/todos", bytes.NewBuffer(invalidData))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	start := time.Now()

	// Log request details
	var requestBody bytes.Buffer
	if r.Body != nil {
		bodyBytes, _ := io.ReadAll(r.Body)
		requestBody.Write(bodyBytes)
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // Reset body for next handler
	}
	log.Printf("Request: %s %s | Body: %s", r.Method, r.URL.Path, requestBody.String())

	// Call the CreateTodo handler
	handler.CreateTodo(rr, r)

	duration := time.Since(start)

	// Log response details
	log.Printf("Response: %d | Duration: %v | Body: %s\n", rr.Code, duration, rr.Body.String())

	// Assert that the response is 400 Bad Request
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "Invalid input\n", rr.Body.String())
}
