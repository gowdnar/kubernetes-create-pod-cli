package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf(">>healthCheckHandler")
	fmt.Fprintf(w, "Welcome home!")
}
func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf(">>updateUserHandler")
	fmt.Fprintf(w, "Welcome updateUserHandler!")
}
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf(">>createUserHandler")
	fmt.Fprintf(w, "Welcome createUserHandler!")
}
func getAllUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf(">>healthCheckHandler")
	fmt.Fprintf(w, "Welcome getAllUserHandler!")
}
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf(">>getUserHandler")
	fmt.Fprintf(w, "Welcome getUserHandler!")
}

func Endpoints() {
	log.Printf(">>Controller")
	router := mux.NewRouter().StrictSlash(true)
	s := router.PathPrefix("/user-service").Subrouter()
	s.HandleFunc("/healthcheck", healthCheckHandler)
	s.HandleFunc("/getUser", getUserHandler)
	s.HandleFunc("/getAllUser", getAllUserHandler)
	s.HandleFunc("/createUser", createUserHandler)
	s.HandleFunc("/updateUser", updateUserHandler)
	err := http.ListenAndServe(":7001", router)
	if err != nil {
		s.HandleFunc("/error", genericErrorHandler)
	}
	log.Printf("<<Controller")
}
func genericErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Application Error")
}
