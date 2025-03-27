# Contacts API - Go + Chi + SQLite (In-Memory) + Logrus

This project is a REST API built with Go, using **Chi** for routing, **SQLite (in-memory)** for data persistence, and **Logrus** for structured logging. The API allows managing contacts with create, read, update, and delete operations.

## 🚀 Features
- ✅ **GET /contacts**: Returns the list of all registered contacts with pagination and filtering support.
- ✅ **POST /contact**: Creates a new contact.
- ✅ **PUT /contact/{id}**: Updates an existing contact.
- ✅ **DELETE /contact/{id}**: Deletes a contact by ID.
- ✅ **Structured logging with Logrus**, ensuring all logs include a `request_id`.
- ✅ **Automatic request tracing with `X-Request-ID`**, generating it when missing.
- ✅ **Middleware-based request processing using Chi**.
- ✅ **Input validation using `go-playground/validator`**, ensuring valid email and name formats.
- ✅ **Standardized error responses**, returning JSON structures with `error.message` and `error.code`.
- ✅ **JWT-based authentication**, requiring a valid token for protected routes.
- ✅ **Pagination and filtering** in `GET /contacts` using query parameters (`?page=1&limit=10&name=John`).

### 📌 Contact Model
Each contact has the following properties:
- `id`: Unique identifier of the contact.
- `name`: Contact's name (validated).
- `email`: Contact's email (validated).

## 🛠️ Technologies Used
- **Go**: Primary programming language.
- **Chi**: Lightweight router for Go.
- **SQLite**: In-memory database for data persistence.
- **Logrus**: Advanced logging library for structured and JSON-based logs.
- **JSON**: Format for API request/response data.
- **go-playground/validator**: Input validation library.
- **golang-jwt/jwt**: JWT authentication library.

## 📖 Usage
### 1️⃣ Start the API
```bash
go run main.go
```

### 2️⃣ Test with `curl`
#### ✅ Authentication (`POST /login`)
```bash
curl -X POST http://localhost:8080/login \
     -H "Content-Type: application/json" \
     -d '{
          "username": "admin",
          "password": "password"
         }'
```
**Response:**
```json
{
  "token": "your-jwt-token-here"
}
```

#### ✅ List Contacts (`GET /contacts` with pagination and filtering)
```bash
curl -X GET "http://localhost:8080/contacts?page=1&limit=10&name=John" \
     -H "Authorization: Bearer your-jwt-token-here" \
     -H "X-Request-ID: 123e4567-e89b-12d3-a456-426614174000"
```

#### ✅ Create a Contact (`POST /contact`)
```bash
curl -X POST http://localhost:8080/contact \
     -H "Authorization: Bearer your-jwt-token-here" \
     -H "Content-Type: application/json" \
     -H "X-Request-ID: 123e4567-e89b-12d3-a456-426614174000" \
     -d '{
          "name": "Alice Smith",
          "email": "alice@example.com"
         }'
```

#### ✅ Update a Contact (`PUT /contact/{id}`)
```bash
curl -X PUT http://localhost:8080/contact/1 \
     -H "Authorization: Bearer your-jwt-token-here" \
     -H "Content-Type: application/json" \
     -H "X-Request-ID: 123e4567-e89b-12d3-a456-426614174000" \
     -d '{
          "name": "Alice Updated",
          "email": "alice.updated@example.com"
         }'
```

#### ✅ Delete a Contact (`DELETE /contact/{id}`)
```bash
curl -X DELETE http://localhost:8080/contact/1 \
     -H "Authorization: Bearer your-jwt-token-here" \
     -H "X-Request-ID: 123e4567-e89b-12d3-a456-426614174000"
```

### 3️⃣ Standardized Error Response
If an error occurs, responses follow this structure:
```json
{
  "error": {
    "message": "Invalid token",
    "code": 401
  }
}
```

### 4️⃣ Check Logs with `X-Request-ID`
All logs now include a `request_id` for traceability. Example log output:
```json
{
  "level": "info",
  "time": "2025-03-24T12:00:00.123Z",
  "msg": "Contact updated",
  "request_id": "123e4567-e89b-12d3-a456-426614174000",
  "name": "Alice Updated",
  "email": "alice.updated@example.com"
}
```