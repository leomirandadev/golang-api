package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"vcfConverter/src/models/users"
	"vcfConverter/src/services/httpResponse"
	"vcfConverter/src/services/jwt"

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

	ok := users.Create(newUser)

	if ok {
		var outputNull interface{}
		httpResponse.RenderOutput(w, "User create successfully", outputNull)
	} else {
		httpResponse.RenderError(w, "ERROR", http.StatusBadGateway)
	}

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	idUser, _ := strconv.ParseInt(params["id"], 10, 64)
	ok := users.Delete(idUser)

	if ok {
		var outputNull interface{}
		httpResponse.RenderOutput(w, "User delete successfully", outputNull)
	} else {
		httpResponse.RenderError(w, "ERROR", http.StatusBadGateway)
	}

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	idUser, _ := strconv.ParseInt(params["id"], 10, 64)

	var userData users.User
	json.NewDecoder(r.Body).Decode(&userData)

	ok := users.Update(idUser, userData)

	if ok {
		var outputNull interface{}
		httpResponse.RenderOutput(w, "User update successfully", outputNull)
	} else {
		httpResponse.RenderError(w, "ERROR", http.StatusBadGateway)
	}

}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user users.User
	json.NewDecoder(r.Body).Decode(&user)

	ok, output := users.GetByEmailPassword(user.Email, user.Password)

	if ok {
		jwtOutput, _ := jwt.GenerateHash(output)
		httpResponse.RenderOutput(w, "Result in output", jwtOutput)
	} else {
		httpResponse.RenderError(w, "ERROR", http.StatusBadGateway)
	}

}
