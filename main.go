package main

import (
	"github.com/gorilla/mux"
	"github.com/rohimihsan/go-login-sys/controllers"
	"github.com/rohimihsan/go-login-sys/middleware"
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

	user := r.PathPrefix("/user").Subrouter()
	user.HandleFunc("/profile", controllers.Profile).Methods("GET")
	user.Use(middleware.MiddlewareAuth)

	r.HandleFunc("/", controllers.TestUp).Methods("GET")
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
