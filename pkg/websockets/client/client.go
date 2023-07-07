package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/azar-intelops/websockets/pkg/websockets/server/models"
	"github.com/gorilla/websocket"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}
	log.Printf("Connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Error connecting to WebSocket server:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	// Start a goroutine to read messages from the WebSocket connection
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("Error reading message from server:", err)
				return
			}

			// Convert the received message to the Message struct
			var receivedMessage models.Message
			err = json.Unmarshal(message, &receivedMessage)
			if err != nil {
				log.Println("Error unmarshaling message:", err)
				continue
			}

			// Log the received message
			log.Printf("[%s]: %s", receivedMessage.Sender, receivedMessage.Content)
		}
	}()

	// Ask user for a username if not set
	var username string
	usernameSet := false

	if usernameSet {
		log.Printf("Username: %s", username)
	} else {
		fmt.Print("Enter username: ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			username = scanner.Text()
			usernameSet = true
		}
	}

	// Start a goroutine to send messages to the WebSocket connection
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				// Read message from user input
				var content string
				// fmt.Print("Enter message: ")
				scanner := bufio.NewScanner(os.Stdin)
				if scanner.Scan() {
					content = scanner.Text()
				}

				// Create a Message struct with username and content
				message := models.Message{
					Sender:    username,
					Content:   strings.TrimSpace(content),
					Timestamp: time.Now(),
				}

				// Send the message as JSON to the WebSocket server
				jsonMessage, err := json.Marshal(message)
				if err != nil {
					log.Println("Error marshaling message:", err)
					continue
				}

				err = c.WriteMessage(websocket.TextMessage, jsonMessage)
				if err != nil {
					log.Println("Error sending message to server:", err)
					continue
				}
			}
		}
	}()

	// Wait for interrupt signal (e.g., Ctrl+C) to gracefully close the connection
	<-interrupt
	log.Println("Closing WebSocket connection...")
	err = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Println("Error closing WebSocket connection:", err)
		return
	}
	select {
	case <-done:
	case <-time.After(time.Second):
		log.Println("Timeout waiting for connection to close")
	}
}
