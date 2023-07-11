# Go Websockets Chat Application

This is a simple chat application built with Go, Gorilla WebSockets, and Gorilla Mux. The application allows users to send and receive messages in real-time using websockets.

## Features

- **Real-time messaging**: Users can send and receive messages instantly through a WebSocket connection.
- **Server-side storage**: Messages are stored on the server using a DAO (Data Access Object) and can be retrieved by clients.

## Prerequisites

- Go 1.16 or higher

## Installation and setup

1. Clone the repository:

   ```bash
   git clone https://github.com/azar-intelops/compage-socket-poc.git
   cd compage-socket-poc
   ```

2. Install the dependencies:
   ```bash
   go mod tidy
   ```
3. Run the server:
   ```bash
   go run main.go
   ```
4. This application runs on `http://localhost:8080` and the client uses this path to access the chat application.
5. Run the client:
   ```bash
   cd pkg/websockets/client
   go run client.go
   ```

### This repo also consists of gitsign 
