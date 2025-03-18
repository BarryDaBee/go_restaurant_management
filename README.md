# Restaurant API 🍽️

A Go Gin-based REST API for managing restaurant operations, including menu items, orders, invoices, tables, and user authentication (WIP)

## 📌 Features

- ✅ User Authentication (JWT)
- ✅ Menu & Food Management
- ✅ Order Processing
- ✅ Invoice Generation
- ✅ Table Reservations
- ✅ Role-Based Access Control (RBAC)
- ✅ Middleware for Secure API Access

## 🛠 Tech Stack

- **Go** (Gin Framework) 🦍
- **GORM** (ORM for PostgreSQL)
- **PostgreSQL** (Database)
- **JWT Authentication** (Token-based security)
<!-- - **Docker** (Containerization) -->

## 📂 Project Structure

```
├── controllers       # API controllers (business logic)
│   ├── food_controller.go
│   └── user_controller.go
│
├── database          # Database connection setup
│   └── database_connection.go
│
├── helpers           # Utility functions
│   └── token_helper.go
│
├── middleware        # Middleware (e.g., authentication)
│   └── auth_middleware.go
│
├── models            # Database models
│   ├── food_model.go
│   ├── order_model.go
│   └── user_model.go
│
├── routes            # API route definitions
│   ├── food_router.go
│   └── user_router.go
│
├── main.go           # Application entry point
└── go.mod            # Go module dependencies
```


## 📝 API Endpoints

### 🔐 Authentication

| Method | Endpoint    | Description               |
|--------|------------|---------------------------|
| POST   | `/register` | Register a new user      |
| POST   | `/login`   | Authenticate user & get JWT |

### 🍔 Food & Menu

| Method | Endpoint      | Description            |
|--------|--------------|------------------------|
| GET    | `/foods`     | Get all food items     |
| GET    | `/foods/:id` | Get a specific food item |
| POST   | `/foods`     | Add a new food item    |
| PUT    | `/foods/:id` | Update food item       |
| DELETE | `/foods/:id` | Delete food item       |


## 🛡️ Authentication & Middleware

All **protected routes** require a **Bearer token** (`Authorization: Bearer <token>`). The middleware ensures:

✅ Only **authenticated users** can access resources.  
