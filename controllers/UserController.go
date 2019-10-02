package controllers

import (
	"encoding/json"
	"github.com/rohimihsan/go-login-sys/models"
	"net/http"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	//param := mux.Vars(r)
	var res models.ResponseResult
	//var result user.User

	user := r.Context().Value("user")

	res.Data = user
	json.NewEncoder(w).Encode(res)
	return
	//
	////get Id
	//var id = param["id"]
	//objID, _ := primitive.ObjectIDFromHex(id)
	//
	//
	//
	////get db
	//db, _ := db.Con()
	//
	////check for mail
	//email_filter := bson.D{{"_id", objID}}
	//
	////get collection
	//err := db.Collection("users").FindOne(context.TODO(), email_filter).Decode(&result)
	//
	//if err != nil {
	//	res.Error = err.Error()
	//	//res.Data = result
	//	json.NewEncoder(w).Encode(res)
	//	return
	//}
	//
	//result.Password = ""
	//res.Data = result
	//json.NewEncoder(w).Encode(res)
	//return
}
