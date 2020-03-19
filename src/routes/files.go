package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"vcfConverter/src/controllers/filesController"
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

func DeleteSomeFiles(w http.ResponseWriter, r *http.Request) {
	var filesId []int64
	json.NewDecoder(r.Body).Decode(&filesId)

	result := filesController.DeleteSome(filesId)

	if result {
		var outputNull interface{}
		httpResponse.RenderOutput(w, "Files deleted successfully", outputNull)
	} else {
		httpResponse.RenderError(w, "ERROR", http.StatusBadGateway)
	}

}
