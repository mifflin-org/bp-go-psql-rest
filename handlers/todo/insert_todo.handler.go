package todo

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	helper "github.com/zerefwayne/go-postgres-rest-docker-boilerplate/helpers/postgres"
	"github.com/zerefwayne/go-postgres-rest-docker-boilerplate/models"
)

// InsertToDoHandler POST	/todo	inserts a new ToDo
func InsertToDoHandler(w http.ResponseWriter, r *http.Request) {

	// Close the request Body after function finishes
	defer r.Body.Close()

	// Create a new insert request to read request body
	requestBody := new(insertRequest)

	// JSON decoder decodes the body into requestBody on basis of json tags in struct
	_ = json.NewDecoder(r.Body).Decode(requestBody)

	// Create a new Todo
	newToDo := new(models.Todo)
	newToDo.Content = requestBody.Content
	newToDo.CreatedAt = time.Now()

	// Use the InsertHandler to insert it into database

	// Here lies the real benefit of seperating database queries from the handler
	// In this structure, it doesn't matter which database we are using
	// Incase we need to change the database, we just have to write the helpers with same
	// function signatures and return options
	// No change has to be made in the controllers except the package with which we are importing from

	if err := helper.Insert(newToDo); err != nil {

		// Set the header type to application/json
		w.Header().Set("Content-Type", "application/json")
		// Marks that there was a 500 error
		w.WriteHeader(http.StatusInternalServerError)

		// Generate a new response body
		resp := new(response)

		resp.Success = false
		resp.Payload = err

		// Convert the body into json String
		responseStr, _ := json.Marshal(resp)

		// Writes the responseStr to the ResponseWriter w
		_, _ = w.Write(responseStr)

		log.Println(err)

	} else {

		w.Header().Set("Content-Type", "application/json")
		// Marks that there was a 200 success
		w.WriteHeader(http.StatusOK)

		resp := new(response)

		body := make(map[string]interface{})

		body["todo"] = newToDo

		resp.Success = true
		resp.Payload = body

		responseStr, _ := json.Marshal(resp)

		_, _ = w.Write(responseStr)

	}

}
