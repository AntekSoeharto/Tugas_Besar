package main

import (
	"fmt"
	"log"
	"net/http"

	controller "github.com/Tugas_Besar/controller"

	"github.com/gorilla/mux"
)

func main() {
	// Do Migrate
	controller.Migrate()

	router := mux.NewRouter()

	//! ------------- USER
	// Get User
	router.HandleFunc("/register", controller.Register).Methods("POST")
	router.HandleFunc("/users", controller.Authenticate(controller.GetAllMember, 0)).Methods("GET")
	router.HandleFunc("/login", controller.LogIn).Methods("GET")
	router.HandleFunc("/logout", controller.Logout).Methods("GET")

	router.HandleFunc("/users/tangguhkan", controller.Authenticate(controller.TangguhkanMember, 0)).Methods("GET")

	router.HandleFunc("/films", controller.Authenticate(controller.InsertFilm, 0)).Methods("POST")

	fmt.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
