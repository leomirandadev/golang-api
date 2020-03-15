package routes

import (
	"encoding/json"
	"net/http"
	"vcfConverter/src/models/users"
	"vcfConverter/src/services/httpResponse"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	ok, output := users.GetAll()

	if ok {
		httpResponse.RenderOutput(w, "Result in output", output)
	} else {
		httpResponse.RenderError(w, "ERROR", http.StatusBadGateway)
	}

}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser users.User
	json.NewDecoder(r.Body).Decode(&newUser)

	ok, output := users.Create(newUser)

	if ok {
		httpResponse.RenderOutput(w, "Result in output", output)
	} else {
		httpResponse.RenderError(w, "ERROR", http.StatusBadGateway)
	}

}
