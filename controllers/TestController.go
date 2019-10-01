package controllers

import (
	"fmt"
	"github.com/rohimihsan/mongotest/config/db"
	"net/http"
)

func TestConn(w http.ResponseWriter, r *http.Request) {
	_, err := db.Con()

	if err != nil {
		fmt.Fprint(w, "Error connectiong to db")
	}

	fmt.Fprint(w, "Connection success")
}

func TestUp(w http.ResponseWriter, r *http.Request) {
	//random := strconv.Itoa(rand.Intn(9999))

	fmt.Fprint(w, "Server is up")
}
