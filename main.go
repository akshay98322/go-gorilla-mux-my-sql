package main

import (
	"log"
	"net/http"

	"github.com/akshay98322/Go-gorilla-mux-my-sql/routers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Define your HTTP routes using the router
	r.HandleFunc("/user", routers.CreateUserHandler).Methods("POST")
	r.HandleFunc("/user/{id}", routers.GetUserHandler).Methods("GET")
	r.HandleFunc("/user/{id}", routers.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/user/{id}", routers.DeleteUserHandler).Methods("DELETE")

	// Start the HTTP server on port 8090
	log.Println("Server listening on :8090")
	log.Fatal(http.ListenAndServe(":8090", r))
}
