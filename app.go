package main

import (
	"log"
	"net/http"
	"vcfConverter/src/models/files"
	"vcfConverter/src/models/users"
	"vcfConverter/src/routes"
	"vcfConverter/src/services/csvToVcf"
	"vcfConverter/src/services/jwt"

	"github.com/gorilla/mux"
)

func main() {
	migrateAll()
	routerHandle := mux.NewRouter()

	// POST router converter
	// // private routes
	routerHandle.HandleFunc("/upload", jwt.Middleware(routes.HandleCsvToVcf)).Methods("POST")

	// users routes
	routerHandle.HandleFunc("/user/login", routes.LoginUser).Methods("POST")
	routerHandle.HandleFunc("/user", routes.CreateUser).Methods("POST")
	// // private routes
	routerHandle.HandleFunc("/users/", jwt.Middleware(routes.GetAllUsers)).Methods("GET")
	routerHandle.HandleFunc("/user/{id}", jwt.Middleware(routes.GetUserById)).Methods("GET")
	routerHandle.HandleFunc("/user/{id}", jwt.Middleware(routes.DeleteUser)).Methods("DELETE")
	routerHandle.HandleFunc("/user/{id}", jwt.Middleware(routes.UpdateUser)).Methods("PUT")

	// files routes
	// // private routes
	routerHandle.HandleFunc("/files/", jwt.Middleware(routes.GetAllFiles)).Methods("GET")
	routerHandle.HandleFunc("/file", jwt.Middleware(routes.CreateFile)).Methods("POST")
	routerHandle.HandleFunc("/file/{id}", jwt.Middleware(routes.DeleteFile)).Methods("DELETE")
	routerHandle.HandleFunc("/files/", jwt.Middleware(routes.DeleteSomeFiles)).Methods("DELETE")

	// deliver file paths for API
	routerHandle.PathPrefix("/folder/files/").Handler(http.StripPrefix("/folder/files/", http.FileServer(http.Dir(csvToVcf.PathOutput)))).Methods("GET")

	// deliver door 80 for API
	log.Fatal(http.ListenAndServe(":80", routerHandle))

}

func migrateAll() {
	users.InitialMigration()
	files.InitialMigration()
}
