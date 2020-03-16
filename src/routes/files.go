package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"vcfConverter/src/models/files"
	"vcfConverter/src/services/httpResponse"

	"github.com/gorilla/mux"
)

func GetAllFiles(w http.ResponseWriter, r *http.Request) {

	ok, output := files.GetAll()

	if ok {
		httpResponse.RenderOutput(w, "Result in output", output)
	} else {
		httpResponse.RenderError(w, "ERROR", http.StatusBadGateway)
	}

}

func CreateFile(w http.ResponseWriter, r *http.Request) {
	var newFile files.File
	json.NewDecoder(r.Body).Decode(&newFile)

	ok, output := files.Create(newFile)

	if ok {
		httpResponse.RenderOutput(w, "Result in output", output)
	} else {
		httpResponse.RenderError(w, "ERROR", http.StatusBadGateway)
	}

}

func DeleteFile(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	idFile, _ := strconv.ParseInt(params["id"], 10, 64)
	ok, output := files.Delete(idFile)

	if ok {
		httpResponse.RenderOutput(w, "Result in output", output)
	} else {
		httpResponse.RenderError(w, "ERROR", http.StatusBadGateway)
	}

}
