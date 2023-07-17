# WebSocket User Management Server

This is a Go application that provides a WebSocket-based API for managing user data. It allows creating, listing, retrieving, updating, and deleting user information using WebSocket communication.

## Prerequisites

- Go 1.16 or higher
- Gorilla Mux package (github.com/gorilla/mux)

## Installation

1. Clone the repository or download the source code files.
2. Install the Gorilla Mux package by running the following command in your terminal:
   ```sh
   go get -u github.com/gorilla/mux
   ```

## Usage

1. Start the server by running the following command:

   ```sh
   go run main.go
   ```

   The server will listen on http://localhost:8080.

2. WebSocket Endpoints:
   - Create User: POST /ws/create
     - Sends a user object to create a new user.
     - Example request payload:
     ```json
     {
       "username": "john",
       "password": "password123"
     }
     ```
   - List Users: GET /ws/list
     - Retrieves a list of all users.
   - Get User by ID: GET /ws/user/{id}
     - Retrieves a user by their ID.
   - Delete User by ID: GET /ws/delete/{id}
     - Deletes a user by their ID.
   - Update User by ID: POST /ws/update/{id}
     - Sends a user object to update an existing user by their ID.
     - Example request payload:
     ```json
     {
       "username": "john_doe",
       "password": "new_password"
     }
     ```
   - Note: Replace {id} in the URLs with the actual user ID.

## Code Structure

The codebase is organized into multiple packages to maintain separation of concerns:

- `main.go:` The entry point of the application. It sets up the Gorilla Mux router and defines the WebSocket endpoints.

- `controllers/user_controller.go`: Contains the implementation of the WebSocket handlers for user-related operations. It handles creating, listing, retrieving, updating, and deleting users.

- `daos/user_dao.go`: Provides a data access object (DAO) for managing user data. It handles the CRUD operations (create, retrieve, update, delete) on the users map.

- `models/user.go`: Defines the User model structure.
