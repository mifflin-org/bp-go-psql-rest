package todo

import (
	"encoding/json"
	"github.com/gorilla/mux"
	helper "github.com/zerefwayne/go-postgres-rest-docker-boilerplate/helpers/postgres"
	"log"
	"net/http"
	"strconv"
)


// DeleteToDoByID	DELETE	/todo/{id}	deletes todo by id
func DeleteToDoByID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	if id, err := strconv.ParseInt(params["id"], 10, 64); err != nil {

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

		if err := helper.DeleteByID(int(id)); err != nil {

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

			resp.Success = true
			resp.Payload = body

			responseStr, _ := json.Marshal(resp)

			_, _ = w.Write(responseStr)

		}
	}

}