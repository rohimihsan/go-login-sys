package main

import (
	"github.com/gorilla/mux"
	"github.com/rohimihsan/mongotest/controllers"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	//test
	r.HandleFunc("/test-connection", controllers.TestConn).Methods("GET")
	r.HandleFunc("/", controllers.TestUp).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
