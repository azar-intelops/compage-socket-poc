package main

import (
	"log"
	"net/http"
	"poc/pkg/websocket/server/controllers"

	"github.com/gorilla/mux"
)

func main() {

	userController := controllers.NewUserController()

	r := mux.NewRouter()
	r.HandleFunc("/ws/create", userController.Create)
	r.HandleFunc("/ws/list", userController.List)
	r.HandleFunc("/ws/user/{id}", userController.GetById)
	r.HandleFunc("/ws/delete/{id}", userController.Delete)
	r.HandleFunc("/ws/update/{id}", userController.Update)



	log.Println("Server started on http://localhost:8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
