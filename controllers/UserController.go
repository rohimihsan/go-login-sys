package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rohimihsan/go-login-sys/models"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	var id = param["id"]

	var res models.ResponseResult

	res.Data = bson.D{{"id", id}}
	json.NewEncoder(w).Encode(res)
	return
}
