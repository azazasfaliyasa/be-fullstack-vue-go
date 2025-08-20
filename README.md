# Fullstack User Management - Backend (Go)

This is the **backend service** for the Fullstack User Management application, built with **Go (Golang)**.  
It provides RESTful APIs for user authentication, authorization, and user management.

## 🚀 Features
- User registration & login with JWT authentication
- Password hashing & secure storage
- CRUD operations for users
- Middleware for authentication & authorization
- MySQL database integration

## 🛠️ Tech Stack
- **Backend**: Go (Golang)
- **Frameworks/Libraries**: net/http, gorilla/mux, GORM
- **Database**: MySQL
- **Authentication**: JWT

## 📂 Project Structure
```
be-fullstack-vue-go/
│── controllers/    # Handles request & response logic
│── database/       # Database connection setup
│── helpers/        # Utility functions (hashing, token, etc.)
│── middlewares/    # Authentication & authorization middleware
│── models/         # Database models (User, etc.)
│── routes/         # API routes
│── structs/        # Request/Response DTOs
│── main.go         # Entry point
│── go.mod
│── go.sum
```

## ⚙️ Environment Variables
Create a `.env` file in the root directory with the following:
```
APP_PORT=3000

DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASS=
DB_NAME=db_go_vue

JWT_SECRET=your-secret-key
```

## ▶️ Running the Project
```bash
# Install dependencies
go mod tidy

# Run the server
go run main.go
```

The backend will run at `http://localhost:3000`

## 📌 API Endpoints
| Method | Endpoint       | Description           |
|--------|---------------|-----------------------|
| POST   | /register     | Register a new user   |
| POST   | /login        | Login user & get token|
| GET    | /users        | Get all users         |
| GET    | /users/:id    | Get user by ID        |
| PUT    | /users/:id    | Update user           |
| DELETE | /users/:id    | Delete user           |

## 🤝 Contribution
Feel free to fork and submit pull requests.

## 📄 License
This project is licensed under the MIT License.
