package main

import (
	"log"
	"net/http"
	"vcfConverter/src/routes"
	"vcfConverter/src/services/csvToVcf"
)

func main() {

	// get router converter
	http.HandleFunc("/upload", routes.HandleCsvToVcf)

	// deliver file paths for API
	http.Handle("/files/", http.StripPrefix("/files", http.FileServer(http.Dir(csvToVcf.PathOutput))))

	// deliver door 80 for API
	log.Fatal(http.ListenAndServe(":80", nil))

}
