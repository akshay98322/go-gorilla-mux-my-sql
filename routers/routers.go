package routers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/akshay98322/Go-gorilla-mux-my-sql/constants"
	models "github.com/akshay98322/Go-gorilla-mux-my-sql/model"
	"github.com/akshay98322/Go-gorilla-mux-my-sql/utils"
	"github.com/gorilla/mux"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(constants.DbDriver, constants.DbUser+":"+constants.DbPass+"@/"+constants.DbName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Parse JSON data from the request body
	var user models.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "User details are not valid", http.StatusNotFound)
		return
	}

	err = utils.CreateUser(db, user.Name, user.Email)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "User created successfully")
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(constants.DbDriver, constants.DbUser+":"+constants.DbPass+"@/"+constants.DbName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Get the 'id' parameter from the URL
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Convert 'id' to an integer
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
		return
	}

	// Call the GetUser function to fetch the user data from the database
	user, err := utils.GetUser(db, userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Convert the user object to JSON and send it in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(constants.DbDriver, constants.DbUser+":"+constants.DbPass+"@/"+constants.DbName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Call the GetUser function to fetch the user data from the database
	users, err := utils.GetAllUsers(db)
	if err != nil {
		http.Error(w, "Users not found", http.StatusNotFound)
		return
	}

	// Convert the user object to JSON and send it in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(constants.DbDriver, constants.DbUser+":"+constants.DbPass+"@/"+constants.DbName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Get the 'id' parameter from the URL
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Convert 'id' to an integer
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
		return
	}

	var user models.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "User details are not valid", http.StatusNotFound)
		return
	}
	// Call the GetUser function to fetch the user data from the database
	err = utils.UpdateUser(db, userID, user.Name, user.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "User updated successfully")
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(constants.DbDriver, constants.DbUser+":"+constants.DbPass+"@/"+constants.DbName)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Get the 'id' parameter from the URL
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Convert 'id' to an integer
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
		return
	}

	err = utils.DeleteUser(db, userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "User deleted successfully")

}
