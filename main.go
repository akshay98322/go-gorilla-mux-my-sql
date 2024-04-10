package main

import (
	"log"
	"net/http"

	"github.com/akshay98322/Go-gorilla-mux-my-sql/routers"
	"github.com/akshay98322/Go-gorilla-mux-my-sql/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	r := mux.NewRouter()
	// create table
	query := "CREATE TABLE users (id INT PRIMARY KEY AUTO_INCREMENT, name VARCHAR(50), email VARCHAR(100));"
	utils.CreateTable(query)

	// Define your HTTP routes using the router
	r.HandleFunc("/user", routers.CreateUserHandler).Methods("POST")
	// r.HandleFunc("/user", routers.GetAllUsersHandler).Methods("GET")
	r.HandleFunc("/user/{id}", routers.GetUserHandler).Methods("GET")
	r.HandleFunc("/user/{id}", routers.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/user/{id}", routers.DeleteUserHandler).Methods("DELETE")

	// Start the HTTP server on port 8090
	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
