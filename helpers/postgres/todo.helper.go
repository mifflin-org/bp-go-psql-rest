package postgres

import (
	"context"
	"errors"
	"log"

	"github.com/zerefwayne/go-postgres-rest-docker-boilerplate/config"
	"github.com/zerefwayne/go-postgres-rest-docker-boilerplate/models"
)

// Insert Inserts a model Todo into the database
func Insert(todo *models.Todo) error {

	log.Printf("helper	|	inserting todo %s\n", todo.Content)

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

// FetchAll Fetches all todos and returns them
func FetchAll() ([]models.Todo, error) {

	log.Printf("helper	|	fetching all todos\n")

	// Intializing empty Todo array
	var todos []models.Todo

	query := "SELECT * FROM todos;"

	if rows, err := config.DB.Query(context.Background(), query); err != nil {
		return todos, err
	} else {

		// Don't forget to close the rows after calling Query()
		defer rows.Close()

		for rows.Next() {

			var todo models.Todo

			// Row is obtained in the order which they were created in database
			if err := rows.Scan(&todo.ID, &todo.Content, &todo.CreatedAt, &todo.Completed); err != nil {
				log.Println(err)
				return todos, err
			}

			todos = append(todos, todo)
		}

		return todos, nil
	}

}

// FetchByID Fetches todo by ID
func FetchByID(id int) (models.Todo, error) {

	log.Printf("helper	|	fetching todo by id\n")

	query := "SELECT * FROM todos WHERE id=$1;"

	var todo models.Todo

	if rows, err := config.DB.Query(context.Background(), query, id); err != nil {
		return todo, err
	} else {

		// Don't forget to close the rows after calling Query()
		defer rows.Close()

		for rows.Next() {

			// Row is obtained in the order which they were created in database
			if err := rows.Scan(&todo.ID, &todo.Content, &todo.CreatedAt, &todo.Completed); err != nil {
				log.Println(err)
				return todo, err
			}

			return todo, nil
		}

		return todo, errors.New("no record found")
	}

}

// UpdateCompletedByID Marks a TO DO Completed
func UpdateCompletedByID(id int) (models.Todo, error) {

	log.Printf("helper	|	marking todo complete by id\n")

	query := "UPDATE todos SET completed=TRUE WHERE id=$1 RETURNING id, content, created_at, completed;"

	var todo models.Todo

	rows := config.DB.QueryRow(context.Background(), query, id)

	if err := rows.Scan(&todo.ID, &todo.Content, &todo.CreatedAt, &todo.Completed); err != nil {
		return todo, err
	}

	return todo, nil

}

// DeleteByID Deleted ToDo By ID
func DeleteByID(id int) error {

	log.Printf("helper	|	deleting todo by id\n")

	query := "DELETE FROM todos WHERE id=$1;"

	if _, err := config.DB.Exec(context.Background(), query, id); err != nil {
		return err
	}

	return nil

}
