package postgres

import (
	"context"
	"log"

	"github.com/zerefwayne/go-postgres-rest-docker-boilerplate/config"
	"github.com/zerefwayne/go-postgres-rest-docker-boilerplate/models"
)

// InsertTodo Inserts a model Todo into the database
func InsertTodo(todo *models.Todo) error {

	log.Printf("helper	|	inserting todo\n%v\n", todo)

	// Template Query to Insert Todo
	query := "INSERT INTO todos(content, created_at) VALUES($1, $2) RETURNING id;"

	var lastInsertedID int

	// Runs the Query and assigns the lastInsertedID
	if err := config.DB.QueryRow(context.Background(), query, todo.Content, todo.CreatedAt).Scan(&lastInsertedID); err != nil {
		return err
	}

	// Successfully inserted the model and assigns the lastReturnedID to todo model
	log.Printf("helper	|	successfully inserted %d\n", lastInsertedID)
	todo.ID = lastInsertedID

	return nil

}
