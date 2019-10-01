package controllers

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rohimihsan/go-login-sys/config/db"
	"github.com/rohimihsan/go-login-sys/models"
	"github.com/rohimihsan/go-login-sys/models/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	//get Id
	var id = param["id"]
	objID, _ := primitive.ObjectIDFromHex(id)

	var res models.ResponseResult
	var result user.User

	//get db
	db, _ := db.Con()

	//check for mail
	email_filter := bson.D{{"_id", objID}}

	//get collection
	err := db.Collection("users").FindOne(context.TODO(), email_filter).Decode(&result)

	if err != nil {
		res.Error = err.Error()
		//res.Data = result
		json.NewEncoder(w).Encode(res)
		return
	}

	result.Password = ""
	res.Data = result
	json.NewEncoder(w).Encode(res)
	return
}
