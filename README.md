![CI](https://github.com/cloudlifter95/go-backed-react-app/actions/workflows/deploy.yml/badge.svg?branch=main)

# Full-Stack Todo App (Go + React + TypeScript)

> [!Important]
> Purely demonstrative. Not intended for production use.

This project is a **full-stack web application** implementing a **Todo List** with a **Go backend** and a **React (TypeScript) frontend**. The backend follows the **MVC + Service Layer** architecture and includes **JWT authentication**, **database migration & seeding**, and **unit tests**. The frontend consumes the API and is built using **React + Vite**, with **frontend testing** implemented.

## **Architecture**
The application follows the **MVC (Model-View-Controller) + Service Layer** pattern:

> [!Warning]
> Authentication module is WIP

```
📂 project-root
├── 📂 backend      # Go backend (API & Business Logic)
│   ├── 📂 models  # Database models
│   ├── 📂 services # Business logic (DB interactions)
│   ├── 📂 handlers # API request handling (controllers)
│   ├── 📂 middleware # Request/response middleware (CORS, Logging, Auth)
│   ├── 📂 migrations # DB migrations & seeding scripts
│   ├── main.go    # Application entry point
│   ├── Dockerfile # Backend Dockerfile (multi-stage build)
│   ├── go.mod     # Go dependencies
│   ├── go.sum     # Go dependency lock file
│   ├── routes.go  # API routes
│   ├── db.go      # Database connection setup
│   ├── tests      # Backend tests (unit & integration)
│
├── 📂 frontend     # React + TypeScript frontend
│   ├── 📂 src     # Frontend source files
│   │   ├── 📂 components # Reusable UI components
│   │   ├── 📂 pages # Page-level components
│   │   ├── 📂 api # API client for backend communication
│   │   ├── 📂 tests # Frontend tests (unit & integration)
│   │   ├── App.tsx # Root component
│   │   ├── main.tsx # Entry point
│   ├── Dockerfile  # Frontend Dockerfile (multi-stage build with Nginx)
│   ├── package.json # Frontend dependencies
│   ├── vite.config.ts # Vite configuration
│
├── docker-compose.yml # Docker orchestration for backend, frontend & DB
├── README.md  # Project Documentation
```

## **Backend (Go)**
- **Web framework**: `gorilla/mux`
- **ORM**: `gorm` (SQLite)
- **Authentication**: JWT-based authentication
- **Middleware**: Logging, CORS, JWT Authentication
- **Unit & Integration Tests**: `testing` package + mocks
- **Database**: SQLite with AutoMigrations & Seeding
- **Containerized Deployment**: Docker + GitHub Actions

### **API Endpoints**
| Method | Endpoint          | Description          |
|--------|------------------|----------------------|
| POST   | `/register`       | Register new user |
| POST   | `/login`          | Authenticate user, return JWT |
| GET    | `/todos`          | Fetch all todos (auth required) |
| POST   | `/todos`          | Create a new todo (auth required) |
| PUT    | `/todos/{id}`     | Update a todo (auth required) |
| DELETE | `/todos/{id}`     | Delete a todo (auth required) |

### **Running the Backend Locally**
```sh
cd backend
export JWT_SECRET=your_secret_key  # Set an environment variable
export DATABASE_URL=sqlite.db       # Database file

# Run the server
go run main.go
```

### **Running Backend Tests**
```sh
cd backend
go test ./...
```

---
## **Frontend (React + TypeScript)**
- **Build tool**: Vite
- **State Management**: React Hooks
- **API Calls**: `fetchTodos` function
- **Testing**: Vitest + React Testing Library
- **Deployment**: Docker + Nginx

### **Running the Frontend Locally**
```sh
cd frontend
npm install
npm run dev
```

### **Running Frontend Tests**
```sh
cd frontend
npm run test
```

---
## **Docker & Deployment**
### **Docker Compose Setup**
A `docker-compose.yml` file orchestrates the backend, frontend, and database.

#### **Run the Full Stack Locally with Docker**
```sh
docker-compose up --build
```

#### **Seed the Database**
```sh
docker exec -it backend-container-name go run migrations/seed.go
```

### **GitHub Actions (CI/CD)**
- **Runs Tests for Backend & Frontend** inside Docker
- **Builds Docker images for backend & frontend**
- **Updates README with CI/CD badge**

---
## **CI/CD Workflow (GitHub Actions)**
The CI/CD pipeline automatically:
1. Runs backend tests inside a Docker container.
2. Runs frontend tests inside a Docker container.
3. Builds & deploys both images locally.
4. Seeds the database.
5. Updates README with a build status badge.

### **CI/CD Status Badge**
![CI/CD Pipeline](https://github.com/your-repo-name/badge.svg)

---
## **Testing the API with Curl**

#### **Register a User**
```sh
curl -X POST http://localhost:8080/register -d '{"username":"user1", "password":"pass"}' -H "Content-Type: application/json"
```

#### **Login to Get JWT Token**
```sh
curl -X POST http://localhost:8080/login -d '{"username":"user1", "password":"pass"}' -H "Content-Type: application/json"
```

#### **Get Todos (Authenticated Request)**
```sh
curl -X GET http://localhost:8080/todos -H "Authorization: Bearer <your_jwt_token>"
```

#### **Create a Todo**
```sh
curl -X POST http://localhost:8080/todos -d '{"title":"New Todo"}' -H "Authorization: Bearer <your_jwt_token>" -H "Content-Type: application/json"
```

---
## **Next Steps**
✅ Add user roles & permissions 🛠️
✅ Enhance frontend with a form to create todos 🎨
✅ Deploy the project to a cloud provider ☁️

---
## **Contributing**
1. Fork the repository.
2. Create a new branch.
3. Submit a pull request.

---
## **License**
This project is licensed under the MIT License.

🚀 **Happy Coding!**


front end created with:
``` bash
npm create vite@latest go-todo-frontend --template react-ts
cd go-todo-frontend
npm install
npm install axios
npm install -D tailwindcss@3 postcss autoprefixer
npx tailwindcss init -p
```