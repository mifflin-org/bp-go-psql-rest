package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/zerefwayne/go-postgres-rest-docker-boilerplate/helpers/postgres"
	"github.com/zerefwayne/go-postgres-rest-docker-boilerplate/models"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/zerefwayne/go-postgres-rest-docker-boilerplate/config"
)

// DefaultHandler GET Returns a generic hello message
func DefaultHandler(w http.ResponseWriter, r *http.Request) {

	// Sets the Response Code as 200
	w.WriteHeader(http.StatusOK)

	// Fprintf writes the format string to w
	fmt.Fprintf(w, "Hello from gopher! You hit the default URL.")

	// as soon as the function ends the response w is returned

}

func testInsert() {

	newTodo := new(models.Todo)
	newTodo.Content = "A new todo!"
	newTodo.CreatedAt = time.Now()

	_ = postgres.InsertTodo(newTodo)

}

func main() {

	log.Println("server	|	initializing")

	// Database setup
	config.ConnectDB()

	// Close the database connection once the main function is finished
	defer config.DB.Close(context.Background())

	// Calls ping method
	config.PingDB()

	testInsert()

	// Creates a new Mux Router
	r := mux.NewRouter()

	// Default route : GET /
	r.HandleFunc("/", DefaultHandler).Methods("GET")

	// This is used to remove CORS that arise when request comes from the same server's another port
	handler := cors.AllowAll().Handler(r)

	// http ListenAndServe is used to listen to requests on 5000 and redirecting them to handler
	if err := http.ListenAndServe(":5000", handler); err != nil {

		// In case there's an error, server is closed and error is logged.
		log.Fatal(err)
	}

}
