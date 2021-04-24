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
	router.HandleFunc("/login", controller.LogIn).Methods("GET")
	router.HandleFunc("/logout", controller.Logout).Methods("GET")

	router.HandleFunc("/users/tangguhkan", controller.Authenticate(controller.TangguhkanMember, 0)).Methods("GET")
	router.HandleFunc("/users/admin", controller.Authenticate(controller.GetMember, 0)).Methods("GET")
	router.HandleFunc("/films/admin", controller.Authenticate(controller.InsertFilm, 0)).Methods("POST")
	router.HandleFunc("/films/admin", controller.Authenticate(controller.GetFilmsbyAdmin, 0)).Methods("GET")
	router.HandleFunc("/films/admin", controller.Authenticate(controller.UpdateFilm, 0)).Methods("PUT")

	router.HandleFunc("/users", controller.Authenticate(controller.UpdateProfile, 1)).Methods("PUT")
	router.HandleFunc("/films", controller.Authenticate(controller.FindFilms, 1)).Methods("GET")

	router.HandleFunc("/riwayatuser", controller.Authenticate(controller.GetRiwayatUser, 1)).Methods("GET")

	router.HandleFunc("/langganan", controller.Authenticate(controller.InsertLangganan, 1)).Methods("POST")
	router.HandleFunc("/langganan", controller.Authenticate(controller.GetLangganan, 0)).Methods("GET")

	fmt.Println("Connected to port 9099")
	log.Fatal(http.ListenAndServe(":9099", router))

}
