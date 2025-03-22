# Go Todo API

A simple RESTful API for managing todos, implemented using **Go** with an **MVC (Model-View-Controller) architecture**.

## 🚀 Architecture: MVC + Service Layer in Go
This project follows the **MVC pattern** along with a **Service Layer**, which helps organize code into logical layers:

- **Model**: Defines the data structures and interacts with the database.
- **View (Handlers)**: Handles HTTP requests and responses.
- **Controller (Routes & Middleware)**: Manages request flow and middleware functions.
- **Service Layer**: Contains business logic, separating it from handlers for better modularity.

### **Why MVC + Service Layer?**
✅ **Separation of Concerns**: Keeps business logic, request handling, and data management independent.
✅ **Scalability**: Easily extendable by adding new models, handlers, services, and routes.
✅ **Maintainability**: Simplifies debugging and code reuse.
✅ **Testability**: Easier unit testing by isolating business logic in services.

## 📁 Folder Structure
```
go-todo/
│── main.go                # Entry point of the application
│── go.mod                 # Go module file
│── db/
│   ├── database.go        # Database connection
│── models/
│   ├── todo.go            # To-Do model
│── services/
│   ├── todo_service.go    # Business logic for todos
│── handlers/
│   ├── todo_handler.go    # HTTP Handlers
│── routes/
│   ├── routes.go          # Route definitions
│── middleware/
│   ├── logging.go         # Request logging middleware
└── README.md              # Project documentation
```

## 🛠️ Setup & Run the Application
### **1. Clone the Repository**
```sh
git clone https://github.com/your-username/go-todo.git
cd go-todo
```

### **2. Install Dependencies**
Ensure you have **Go installed** (1.18+ recommended):
```sh
go mod tidy
```

### **3. Run the Application**
```sh
go run main.go
```
Server starts at `http://localhost:8080`.

## 📡 API Endpoints
| Method | Endpoint    | Description          |
|--------|------------|----------------------|
| GET    | `/todos`   | Fetch all todos      |
| POST   | `/todos`   | Create a new todo    |

## 🧪 Testing the API with cURL
### **1. Create a Todo**
```sh
curl -X POST http://localhost:8080/todos -H "Content-Type: application/json" -d '{"title":"Buy groceries"}'
```

### **2. Get All Todos**
```sh
curl -X GET http://localhost:8080/todos
```

## 📝 Logging Middleware
This project includes a **logging middleware** that logs each HTTP request and response.

### **Example Log Output**
```
Request: POST /todos | Body: {"title":"Buy groceries"}
Response: 201 | Duration: 2.5ms | Body: {"id":1,"title":"Buy groceries","completed":false}
```

## ✅ Next Steps
- Add authentication.
- Implement database migrations.
- Expand API functionality (update, delete todos).

---
Built with ❤️ using Go! 🚀


---
# Live Dev
## Reflex
To use Reflex:
`reflex -r '\.go$' -- sh -c 'go run main.go'`
you might need to ensure the Go binary directory is in your PATH. The default location is usually $GOPATH/bin or $HOME/go/bin (depending on your Go version).
`export PATH=$PATH:$(go env GOPATH)/bin`

run reflex: `reflex -r '\.go$' -- sh -c 'go run main.go'`

## Air
`go install github.com/air-verse/air@latest`
you might need to ensure the Go binary directory is in your PATH. The default location is usually $GOPATH/bin or $HOME/go/bin (depending on your Go version).
`export PATH=$PATH:$(go env GOPATH)/bin`
run air: `air`


## Go test commands

`go test ./tests/handlers/ -run TestCreateTodo_InvalidInput -v`
`go test ./...`

