# GoLang CRUD API with Gin and Gorm

- Simple CRUD (Create, Read, Update, Delete) API built with Go (Golang), using Gin and Gorm
- PostgreSQL is used as the database

## Getting Started

### Prerequisites

- Go installed on your machine
- PostgreSQL database set up
- Create a `.env` file with the required environment variables (refer to `.env.example`)

## Installation

- **Clone the repository:**

  ```bash
  git clone https://github.com/mxnuchim/golang-crud-api.git
  ```

- **Navigate to the project directory:**

  ```bash
  cd golang-crud-api
  ```

- **Clone the repository:**

  ```bash
  git clone https://github.com/mxnuchim/golang-crud-api.git
  ```

- **Install dependencies:**

  ```bash
  go mod tidy
  ```

- **Migrate the database:**

  ```bash
  go run migrate/migrate.go
  ```

- **Run the server:**

  ```bash
  go run main.go
  ```

- **Using CompileDaemon:**
  you may also use CompileDaemon to run the server and watch for file changes
  ```bash
  CompileDaemon -command="./golang-crud-api"
  ```

## API Endpoints

- **Create a Post**

  ```http
  POST /posts
  ```

- **Update a Post**

  ```http
  PATCH /posts/:id
  ```

- **Get all Posts**

  ```http
  GET /posts
  ```

- **Get a single Post**

  ```http
  GET /posts/:id
  ```

- **Delete a Post**

  ```http
  GET /posts/:id
  ```
# golang-crud-api
