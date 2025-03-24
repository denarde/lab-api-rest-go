# Contact API in Go with SQLite in Memory

This project is an example of a simple REST API developed in Go, using SQLite in memory for data persistence. The API allows contact management with create and read operations.

## Features

- **GET** `/contacts`: Returns the list of all registered contacts.
- **POST** `/contact`: Creates a new contact.

Each contact has the following properties:

- `id`: Unique identifier for the contact.
- `name`: Contact's name.
- `email`: Contact's email.

## Technologies Used

- **Go**: Main programming language.
- **SQLite**: In-memory database for data persistence.
- **JSON**: Format for data exchange between the API and the client.
