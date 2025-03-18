package main

import (
	"fmt"
	"log"
	"net/http"

	// "go-todo/db"
	"go-todo/db"
	"go-todo/middleware"
	"go-todo/routes"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize Database
	DB, err := db.InitDB()
	if err != nil {
		log.Fatal("Error creating DB")
	}
	sqlDB, _ := DB.DB()
	fmt.Println("DB connection:", sqlDB)

	// Create a new router
	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.EnableCORS)

	// Register routes
	routes.SetupRoutes(r, DB)

	// Start the server
	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
