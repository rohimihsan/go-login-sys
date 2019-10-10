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

	profile := r.PathPrefix("/profile").Subrouter()
	profile.HandleFunc("/", controllers.Profile).Methods("GET")
	profile.Use(middleware.MiddlewareAuth)

	user := r.PathPrefix("/user").Subrouter()
	user.HandleFunc("/login", controllers.Login).Methods("POST")
	user.HandleFunc("/register", controllers.Register).Methods("POST")
	user.Use(middleware.CorsByPass)

	r.HandleFunc("/", controllers.TestUp).Methods("GET")
	r.HandleFunc("/set-cookie", SetCookie).Methods("POST")

	r.Methods("OPTIONS").HandlerFunc(PreFlight)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func SetCookie(w http.ResponseWriter, r *http.Request) {
	whitelist := "http://localhost:3000";
	w.Header().Set("Access-Control-Allow-Origin", whitelist)
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Expose-Headers", "Session-Token")

	c, err := r.Cookie("token")
	if err != nil {
		c = &http.Cookie{
			Name:  "token",
			Value: "asdmaskdljalje12ej912jemlkjkdajdoqu91cueovkjljkljkjljijlkajkasjdskaldjalskdjksadjsaldkuracj",
			HttpOnly: true,
			//Secure: false,
			Domain: "localhost",
		}
		http.SetCookie(w, c)
	}
	//io.WriteString(w, c.String())
}

func PreFlight(w http.ResponseWriter, r *http.Request) {
	// Check the origin is valid.
	//origin := r.Header.Get("Origin")

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Expose-Headers", "Session-Token")
}
