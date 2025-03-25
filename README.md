# Contacts API - Go + Chi + SQLite (In-Memory) + Logrus

This project is a simple REST API built with Go, using **Chi** for routing, **SQLite (in-memory)** for data persistence, and **Logrus** for structured logging. The API allows managing contacts with create and read operations.

## 🚀 Features
- ✅ **GET /contacts**: Returns the list of all registered contacts.
- ✅ **POST /contact**: Creates a new contact.
- ✅ **Structured logging with Logrus**, ensuring all logs include a `request_id`.
- ✅ **Automatic request tracing with `X-Request-ID`**, generating it when missing.
- ✅ **Middleware-based request processing using Chi**.

### 📌 Contact Model
Each contact has the following properties:
- `id`: Unique identifier of the contact.
- `name`: Contact's name.
- `email`: Contact's email.

## 🛠️ Technologies Used
- **Go**: Primary programming language.
- **Chi**: Lightweight router for Go.
- **SQLite**: In-memory database for data persistence.
- **Logrus**: Advanced logging library for structured and JSON-based logs.
- **JSON**: Format for API request/response data.

## 📖 Usage
### 1️⃣ Start the API
```bash
go run main.go
```

### 2️⃣ Test with `curl`
#### ✅ List Contacts (`GET /contacts`)
```bash
curl -X GET http://localhost:8080/contacts \
     -H "X-Request-ID: 123e4567-e89b-12d3-a456-426614174000"
```

#### ✅ Create a Contact (`POST /contact`)
```bash
curl -X POST http://localhost:8080/contact \
     -H "Content-Type: application/json" \
     -H "X-Request-ID: 123e4567-e89b-12d3-a456-426614174000" \
     -d '{
          "name": "Alice Smith",
          "email": "alice@example.com"
         }'
```

### 3️⃣ Check Logs with `X-Request-ID`
All logs now include a `request_id` for traceability. Example log output:
```json
{
  "level": "info",
  "time": "2025-03-24T12:00:00.123Z",
  "msg": "Contact created",
  "request_id": "123e4567-e89b-12d3-a456-426614174000",
  "name": "Alice Smith",
  "email": "alice@example.com"
}
```

