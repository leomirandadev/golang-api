package httpResponse

import (
	"encoding/json"
	"net/http"
)

type DefaultOutput struct {
	OK      bool        `json:"ok"`
	MESSAGE string      `json:"message"`
	OUTPUT  interface{} `json:"output"`
}

func RenderError(w http.ResponseWriter, message string, serverError int) {

	w.WriteHeader(serverError)
	output := DefaultOutput{false, message, nil}
	json.NewEncoder(w).Encode(output)

}

func RenderOutput(w http.ResponseWriter, message string, result interface{}) {

	w.WriteHeader(http.StatusOK)
	output := DefaultOutput{true, message, result}
	json.NewEncoder(w).Encode(output)

}
