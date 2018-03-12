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
	db = dbConnect()
	defer dbDisconnect(db)

	params := mux.Vars(r)

	name := params["name"]

	var user User

	result := db.First(&user, "name = ?", name)

	if result.RowsAffected == 0 {
		//json.NewEncoder(w).Encode(&User{Name: "USER NOT FOUND", Email: "USER NOT FOUND"})
		fmt.Fprintf(w, "User with name \"%v\" cannot be found!\n", name)
	} else {
		json.NewEncoder(w).Encode(user)
	}

}

func insertUser(w http.ResponseWriter, r *http.Request) {

	db = dbConnect()
	defer dbDisconnect(db)

	params := mux.Vars(r)
	name := params["name"]
	email := params["email"]

	db.Create(&User{Name: name, Email: email})
	fmt.Fprintf(w, "User \"%v\" was successfully created!\n", name)

}

func updateUser(w http.ResponseWriter, r *http.Request) {
	db = dbConnect()
	defer dbDisconnect(db)

	params := mux.Vars(r)
	name := params["name"]
	email := params["email"]

	var userToUpdate User
	//db.Find(&userToUpdate, "name = ?", name)
	db.Model(&userToUpdate).Where("name = ?", name).Update("email", email)

	fmt.Fprintf(w, "Updated the \"%v\" user email address to: \"%v\"\n", name, email)

}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	db = dbConnect()
	defer dbDisconnect(db)

	vars := mux.Vars(r)
	name := vars["name"]

	var userToDelete User
	ret := db.Delete(&userToDelete, "name = ?", name)

	if ret.Error != nil {
		panic(fmt.Sprintf("User \"%v\" cannot be deleted!\nERROR: %v", name, ret.Error))
	} else if ret.RowsAffected == 0 {
		fmt.Fprintf(w, "User \"%v\" doesn't exists!\nNothing to delete!\nRows Affected: %v", name, ret.RowsAffected)
		panic("Nothing to delete!")
	}

	fmt.Fprintf(w, "User \"%v\" was successfully deleted!\n", name)

}
