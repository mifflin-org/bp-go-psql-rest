package todo

import (
	"encoding/json"
	helper "github.com/zerefwayne/go-postgres-rest-docker-boilerplate/helpers/postgres"
	"log"
	"net/http"
)


// InsertToDoHandler GET	/todos	fetches all todos
func FetchToDoAll(w http.ResponseWriter, r *http.Request) {

	if todos, err := helper.FetchAll(); err != nil {

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

		body["todos"] = todos
		body["count"] = len(todos)

		resp.Success = true
		resp.Payload = body

		responseStr, _ := json.Marshal(resp)

		_, _ = w.Write(responseStr)

	}

}