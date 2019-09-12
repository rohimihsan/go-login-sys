package main

import (
	"github.com/gorilla/mux"
	"github.com/rohimihsan/go-login-sys/middleware"
	"github.com/rohimihsan/mongotest/controllers"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	//test route group
	t := r.PathPrefix("/test").Subrouter()
	t.HandleFunc("/", controllers.TestUp).Methods("GET")
	t.HandleFunc("/connection", controllers.TestConn).Methods("GET")
	//test middleware
	t.Use(middleware.MiddlewareAllowOnlyGet)

	r.HandleFunc("/lol", controllers.TestUp).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
