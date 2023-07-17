package controllers

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/json"
	"log"
	"net/http"
	"poc/pkg/websocket/server/models"
	"poc/pkg/websocket/server/services"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// Define the WebSocket upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

func generateRandomID() int64 {
	// Generate 8 random bytes
	randomBytes := make([]byte, 8)
	_, err := rand.Read(randomBytes)
	if err != nil {
		// Handle the error
		panic(err)
	}

	// Convert the random bytes to int64
	randomID := int64(binary.BigEndian.Uint64(randomBytes))
	return randomID
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return nil, err
	}
	return conn, nil
}

func (userController *UserController) Create(w http.ResponseWriter, r *http.Request) {
	conn, err := HandleWebSocket(w, r)

	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}

	defer conn.Close()

	for {
		var user models.User
		err := conn.ReadJSON(&user)
		if err != nil {
			log.Println("Failed to read message:", err)
			conn.Close()
			break
		}

		user.Id = generateRandomID()

		_, err2 := userController.userService.CreateItem(user)
		if err2 != nil {
			log.Println("Unable to store the user", err2)
			w.WriteHeader(http.StatusInternalServerError)
			break
		}
		
	}
	w.WriteHeader(http.StatusCreated)
}

func (userController *UserController) List(w http.ResponseWriter, r *http.Request) {
	conn, err := HandleWebSocket(w, r)

	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}

	defer conn.Close()

	for {
		users, err := userController.userService.ListItem()
		if err != nil {
			log.Println("Failed to fetch message:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		usersJSON, err := json.Marshal(users)
		if err != nil {
			log.Println("Failed to marshal users to JSON:", err)
			break
		}

		err = conn.WriteMessage(websocket.TextMessage, usersJSON)
		if err != nil {
			log.Println("Failed to send WebSocket message:", err)
			break
		}
		time.Sleep(1 * time.Second)
	}
	w.WriteHeader(http.StatusOK)
}

func (userController *UserController) GetById(w http.ResponseWriter, r *http.Request) {
	conn, err := HandleWebSocket(w, r)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}

	defer conn.Close()

	params := mux.Vars(r)

	id, _ := strconv.ParseInt(params["id"], 10, 64)
	for {
		user, err := userController.userService.GetItem(id)
		if err != nil {
			log.Println("Failed to fetch message:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		userJSON, err := json.Marshal(user)
		if err != nil {
			log.Println("Failed to marshal users to JSON:", err)
			break
		}

		err = conn.WriteMessage(websocket.TextMessage, userJSON)
		if err != nil {
			log.Println("Failed to send WebSocket message:", err)
			break
		}
		time.Sleep(1 * time.Second)
	}
	w.WriteHeader(http.StatusOK)
}

func (userController *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	conn, err := HandleWebSocket(w, r)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}

	defer conn.Close()

	params := mux.Vars(r)

	id, _ := strconv.ParseInt(params["id"], 10, 64)
	for {
		err := userController.userService.DeleteItem(id)
		if err != nil {
			log.Println("Failed to fetch message:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		userJSON, err := json.Marshal("user deleted successfully!")
		if err != nil {
			log.Println("Failed to marshal users to JSON:", err)
			break
		}

		err = conn.WriteMessage(websocket.TextMessage, userJSON)
		if err != nil {
			log.Println("Failed to send WebSocket message:", err)
			break
		}
		time.Sleep(1 * time.Second)
	}
	w.WriteHeader(http.StatusOK)
}

func (userController *UserController) Update(w http.ResponseWriter, r *http.Request) {
	conn, err := HandleWebSocket(w, r)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}

	defer conn.Close()

	params := mux.Vars(r)

	id, _ := strconv.ParseInt(params["id"], 10, 64)
	for {
		var user models.User
		var res models.User
		err := conn.ReadJSON(&user)
		if err != nil {
			log.Println("Failed to read message:", err)
			conn.Close()
			break
		}

		res, err = userController.userService.UpdateItem(id, user)
		if err != nil {
			log.Println("Failed to fetch message:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		userJSON, err := json.Marshal(res)
		if err != nil {
			log.Println("Failed to marshal users to JSON:", err)
			break
		}

		err = conn.WriteMessage(websocket.TextMessage, userJSON)
		if err != nil {
			log.Println("Failed to send WebSocket message:", err)
			break
		}
		time.Sleep(1 * time.Second)
	}
	w.WriteHeader(http.StatusOK)
}