package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"vcfConverter/src/models/users"
	"vcfConverter/src/services/httpResponse"

	"github.com/gorilla/mux"
)

func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idUser, _ := strconv.ParseInt(params["id"], 10, 64)
	ok, output := users.GetById(idUser)

	if ok {
		httpResponse.RenderOutput(w, "Result in output", output)
	} else {
		httpResponse.RenderError(w, "ERROR", http.StatusBadGateway)
	}

}

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

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	idUser, _ := strconv.ParseInt(params["id"], 10, 64)
	ok, output := users.Delete(idUser)

	if ok {
		httpResponse.RenderOutput(w, "Result in output", output)
	} else {
		httpResponse.RenderError(w, "ERROR", http.StatusBadGateway)
	}

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	idUser, _ := strconv.ParseInt(params["id"], 10, 64)

	var userData users.User
	json.NewDecoder(r.Body).Decode(&userData)

	ok, output := users.Update(idUser, userData)

	if ok {
		httpResponse.RenderOutput(w, "Result in output", output)
	} else {
		httpResponse.RenderError(w, "ERROR", http.StatusBadGateway)
	}

}
