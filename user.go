package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// User struct
type User struct {
	gorm.Model
	Name  string
	Email string
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All users:\n")

	db = dbConnect()
	defer dbDisconnect(db)

	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User:\n")
}

func insertUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "insertUser HIT!\n")

	db = dbConnect()
	defer dbDisconnect(db)

	params := mux.Vars(r)
	name := params["name"]
	email := params["email"]

	db.Create(&User{Name: name, Email: email})
	fmt.Fprintf(w, "User \"%v\" was successfully created!\n", name)

}

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "updateUser HIT!\n")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "deleteUser HIT!\n")
}
