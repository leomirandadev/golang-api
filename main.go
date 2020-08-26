package main

import (
	"log"
	"net/http"

	"github.com/leomirandadev/golang-api/api/handlers"
	"github.com/leomirandadev/golang-api/api/middlewares"
	"github.com/leomirandadev/golang-api/models/files"
	"github.com/leomirandadev/golang-api/models/users"
	"github.com/leomirandadev/golang-api/services/csvToVcf"

	"github.com/gorilla/mux"
)

func main() {
	migrateAll()
	routerHandle := mux.NewRouter()

	// POST router converter
	// // private api/handlers
	routerHandle.HandleFunc("/upload", middlewares.JwtVerify(handlers.HandleCsvToVcf)).Methods("POST")

	// users handlers
	routerHandle.HandleFunc("/user/login", handlers.LoginUser).Methods("POST")
	routerHandle.HandleFunc("/user", handlers.CreateUser).Methods("POST")
	// // private handlers
	routerHandle.HandleFunc("/users/", middlewares.JwtVerify(handlers.GetAllUsers)).Methods("GET")
	routerHandle.HandleFunc("/user/{id}", middlewares.JwtVerify(handlers.GetUserById)).Methods("GET")
	routerHandle.HandleFunc("/user/{id}", middlewares.JwtVerify(handlers.DeleteUser)).Methods("DELETE")
	routerHandle.HandleFunc("/user/{id}", middlewares.JwtVerify(handlers.UpdateUser)).Methods("PUT")

	// files handlers
	// // private handlers
	routerHandle.HandleFunc("/files/", middlewares.JwtVerify(handlers.GetAllFiles)).Methods("GET")
	routerHandle.HandleFunc("/file", middlewares.JwtVerify(handlers.CreateFile)).Methods("POST")

	// deliver file paths for API
	routerHandle.PathPrefix("/folder/files/").Handler(http.StripPrefix("/folder/files/", http.FileServer(http.Dir(csvToVcf.PathOutput)))).Methods("GET")

	// deliver door 80 for API
	log.Fatal(http.ListenAndServe(":80", routerHandle))

}

func migrateAll() {
	users.InitialMigration()
	files.InitialMigration()
}
