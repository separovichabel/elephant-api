package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/separovichabel/monkey-api/config"
	. "github.com/separovichabel/monkey-api/db/dao"
	userRoutes "github.com/separovichabel/monkey-api/routes"
)

var dao = UserDAO{}
var config = Config{}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", userRoutes.GetAll).Methods("GET")
	r.HandleFunc("/users/{id}", userRoutes.GetByID).Methods("GET")
	r.HandleFunc("/users", userRoutes.Create).Methods("POST")
	r.HandleFunc("/users/{id}", userRoutes.Update).Methods("PUT")
	r.HandleFunc("/users/{id}", userRoutes.Delete).Methods("DELETE")

	var port = ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}
