package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func welcomeMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome GoLang")
}

func requestsHandler() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", welcomeMessage).Methods(http.MethodGet)
	router.HandleFunc("/users", getAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/user/{name}", getUser).Methods(http.MethodGet)
	router.HandleFunc("/user/{name}/{email}", insertUser).Methods(http.MethodPost)
	router.HandleFunc("/user/{name}/{email}", updateUser).Methods(http.MethodPut)
	router.HandleFunc("/user/{name}", deleteUser).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":1234", router))
}

func main() {
	fmt.Println("Go ORM REST API")

	dbAutoMigrate()

	requestsHandler()
}
