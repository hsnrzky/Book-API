# 📚 Book API

A simple and clean **RESTful Book Management API** built with **Golang (Go)** using the **Gin Web Framework** and **SQLite** database.  
This project demonstrates basic CRUD operations and clean project structure for backend API development.

---

## 📌 Table of Contents

- [Features](#-features)
- [Tech Stack](#-tech-stack)
- [Installation](#-installation)
- [Project Structure](#-project-structure)
- [API Endpoints](#-api-endpoints)
- [Request & Response Example](#-request--response-example)
- [Notes](#-notes)
- [Contribution](#-contribution)
- [License](#-license)

---

## 🚀 Features

- RESTful API architecture
- Create, Read, Update, Delete (CRUD) books
- JSON request & response
- SQLite database integration
- Clean and modular project structure
- Beginner-friendly and easy to extend

---

## 🛠️ Tech Stack

- **Language:** Go (Golang)
- **Framework:** Gin Web Framework
- **Database:** SQLite
- **API Style:** RESTful
- **ORM / Query:** Native / lightweight approach

---

## 📦 Installation

Follow the steps below to run this project locally.

### 1️⃣ Prerequisites

Make sure you have installed:

- Go (v1.18 or later)
- Git

### 2️⃣ Clone the Repository

***```bash
git clone https://github.com/hsnrzky/Book-API.git
cd Book-API```***

### 3️⃣ Install Dependencies
```bash
Copy code
go mod tidy

4️⃣ Database Setup
This project uses SQLite and comes with a default database file.

Database file: crud.db

No additional setup required

The database will be automatically used when the application starts

To reset the database, simply delete crud.db and restart the app.

5️⃣ Run the Application
bash
Copy code
go run main.go
The server will start at:

arduino
Copy code
http://localhost:8080
📁 Project Structure
python
Copy code
Book-API/
├── config/         # Database configuration
├── controllers/    # Request handlers
├── models/         # Data models
├── routes/         # API routes
├── crud.db         # SQLite database
├── go.mod
├── go.sum
└── main.go         # Application entry point
📡 API Endpoints
Base URL

arduino
Copy code
http://localhost:8080
Method	Endpoint	Description
GET	/books	Get all books
GET	/books/{id}	Get book by ID
POST	/books	Create a new book
PUT	/books/{id}	Update an existing book
DELETE	/books/{id}	Delete a book

📄 Request & Response Example
Create Book
Request

bash
Copy code
POST /books
Content-Type: application/json
json
Copy code
{
  "title": "Learn Golang",
  "author": "Hsnrzky",
  "description": "Simple REST API using Golang",
  "year": 2025,
  "isbn": "9781234567890"
}
Response

json
Copy code
{
  "id": 1,
  "title": "Learn Golang",
  "author": "Hsnrzky",
  "description": "Simple REST API using Golang",
  "year": 2025,
  "isbn": "9781234567890"
}
📝 Notes
This API does not use authentication

Designed for learning and demonstration purposes

Easy to extend with JWT, Swagger, or PostgreSQL

🤝 Contribution
Contributions are welcome!
Feel free to open an issue or submit a pull request.

📄 License
This project is open-source and available for educational purposes.