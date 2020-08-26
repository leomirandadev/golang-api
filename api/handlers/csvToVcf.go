package handlers

import (
	"net/http"

	"github.com/leomirandadev/golang-api/services/csvToVcf"
	"github.com/leomirandadev/golang-api/services/fileHttpTransfer"
	"github.com/leomirandadev/golang-api/services/httpResponse"
)

func HandleCsvToVcf(w http.ResponseWriter, r *http.Request) {

	ok, pathFileCsv := fileHttpTransfer.Up(w, r)

	if ok {
		ok, output := csvToVcf.Main(pathFileCsv)

		if ok {
			httpResponse.RenderOutput(w, "Result in output", output)
		} else {
			httpResponse.RenderError(w, "ERROR_CONVERTE_ARCHIVE", http.StatusInternalServerError)
		}
	}

}
