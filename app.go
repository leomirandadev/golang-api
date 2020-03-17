package main

import (
	"log"
	"net/http"
	"vcfConverter/src/models/files"
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
	routerHandle.HandleFunc("/user/login", routes.LoginUser).Methods("POST")
	routerHandle.HandleFunc("/user/{id}", routes.GetUserById).Methods("GET")
	routerHandle.HandleFunc("/user/{id}", routes.DeleteUser).Methods("DELETE")
	routerHandle.HandleFunc("/user/{id}", routes.UpdateUser).Methods("PUT")

	// files routes
	routerHandle.HandleFunc("/files/", routes.GetAllFiles).Methods("GET")
	routerHandle.HandleFunc("/file", routes.CreateFile).Methods("POST")
	routerHandle.HandleFunc("/file/{id}", routes.DeleteFile).Methods("DELETE")

	// // deliver file paths for API
	routerHandle.PathPrefix("/folder/files/").Handler(http.StripPrefix("/folder/files/", http.FileServer(http.Dir(csvToVcf.PathOutput)))).Methods("GET")

	// // deliver door 80 for API
	log.Fatal(http.ListenAndServe(":80", routerHandle))

}

func migrateAll() {
	users.InitialMigration()
	files.InitialMigration()
}
