package main

import (
	"log"
	"net/http"
	"vcfConverter/src/models/users"
	"vcfConverter/src/routes"
	"vcfConverter/src/services/csvToVcf"

	"github.com/gorilla/mux"
)

func main() {
	migrateAll()
	routerHandle := mux.NewRouter()

	// POST router converter
	routerHandle.HandleFunc("/upload", routes.HandleCsvToVcf).Methods("POST")

	// users routes
	routerHandle.HandleFunc("/users/", routes.GetAllUsers).Methods("GET")
	routerHandle.HandleFunc("/user", routes.CreateUser).Methods("POST")

	// // deliver file paths for API
	routerHandle.PathPrefix("/files/").Handler(http.StripPrefix("/files/", http.FileServer(http.Dir(csvToVcf.PathOutput)))).Methods("GET")

	// // deliver door 80 for API
	log.Fatal(http.ListenAndServe(":80", routerHandle))

}

func migrateAll() {
	users.InitialMigrationUser()
}
