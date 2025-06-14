# 🍽️ Golang Restaurant Management Backend

A RESTful API backend for a restaurant management system built with **Go (Gin)** and **MongoDB**, featuring JWT auth, menu management, food items, orders, invoices, tables, and users.

---

## 🚀 Features

- ✅ User signup/login with JWT authentication  
- 🍴 Menu & Food Item CRUD  
- 🛒 Order creation & updates  
- 📦 Manage food items within orders  
- 💰 Invoice generation with due tracking  
- 🍽️ Table management  
- 🔒 Protected routes via JWT middleware  

---

## 📁 Project Structure
```
├── controllers/ # Route handlers (user, food, menu, order, etc.)
├── routes/ # Maps routes to handlers
├── models/ # Data schema definitions
├── database/ # MongoDB connection logic
├── helpers/ # JWT utilities
├── middleware/ # JWT auth middleware
├── main.go # Server entry point
├── go.mod / go.sum # Dependencies
```

---

## 🧩 Tech Stack

- **Go (Gin)** – Web framework  
- **MongoDB** – Database  
- **JWT** – Authentication  
- **validator.v10** – Input validation  

---

## 🛠️ Requirements

- Go 1.16 or later  
- MongoDB running locally or in cloud  
- Environment variables:
```
MONGODB_URI=<your-mongodb-uri>
SECRET_KEY=<your-jwt-secret>
```

---

## ⚙️ Setup

```bash
# Clone the repo
git clone https://github.com/AkhilSharma90/golang-restaurant-management-backend.git
cd golang-restaurant-management-backend

# Install dependencies
go mod tidy

# Set environment variables
export MONGODB_URI="mongodb://localhost:27017"
export SECRET_KEY="yourSecretKey"

# Run the server
go run main.go

Server will run at: http://localhost:8000
