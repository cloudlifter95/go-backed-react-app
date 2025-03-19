``` bash
go-todo/                # Backend (Go)
│── main.go             # Entry point of the application
│── go.mod              # Go module file
│── db/
│   ├── database.go     # Database connection
│── models/
│   ├── todo.go         # To-Do model
│── handlers/
│   ├── todo_handler.go # HTTP Handlers
│── routes/
│   ├── routes.go       # Route definitions
│── services/
│   ├── todo_service.go # Business logic (service layer)
│── middleware/
│   ├── logger.go       # Middleware for logging requests
│── README.md           # Documentation
│── .env                # Environment variables
│── Dockerfile          # Optional: Dockerfile for deployment
└── .gitignore          # Git ignore file

go-todo-frontend/       # Frontend (React + TypeScript)
│── src/
│   ├── components/
│   │   ├── TodoList.tsx    # Displays list of todos
│   │   ├── TodoForm.tsx    # Form to add new todos
│   ├── pages/
│   │   ├── Home.tsx        # Main page
│   ├── api.ts              # API calls (Axios)
│   ├── types.ts            # TypeScript types
│   ├── App.tsx             # Root component
│   ├── main.tsx            # Entry point
│── public/
│   ├── index.html          # Main HTML file
│── package.json            # Dependencies & scripts
│── tsconfig.json           # TypeScript config
│── tailwind.config.js      # TailwindCSS config
│── vite.config.ts          # Vite config
│── .gitignore              # Git ignore file
│── README.md               # Frontend documentation
└── .env                    # Environment variables (optional)
```

front end created with:
``` bash
npm create vite@latest go-todo-frontend --template react-ts
cd go-todo-frontend
npm install
npm install axios
npm install -D tailwindcss@3 postcss autoprefixer
npx tailwindcss init -p
```