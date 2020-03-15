package routes

import (
	"net/http"
	"vcfConverter/src/services/csvToVcf"
	"vcfConverter/src/services/fileHttpTransfer"
	"vcfConverter/src/services/httpResponse"
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
