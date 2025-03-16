#!/bin/bash

# Create the go-todo directory
mkdir -p go-todo

# Move into go-todo directory
cd go-todo

# Create the main files
touch main.go go.mod README.md

# Create the db directory and file
mkdir -p db
touch db/database.go

# Create the models directory and file
mkdir -p models
touch models/todo.go

# Create the handlers directory and file
mkdir -p handlers
touch handlers/todo_handler.go

# Create the routes directory and file
mkdir -p routes
touch routes/routes.go

echo "Project structure created successfully in $(pwd)/go-todo"
