package controllers

import (
	"context"
	"encoding/json"
	"github.com/rohimihsan/go-login-sys/config/db"
	"github.com/rohimihsan/go-login-sys/models"
	"github.com/rohimihsan/go-login-sys/models/user"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	pass := r.FormValue("password")
	//hash := r.FormValue("hash")

	var res models.ResponseResult

	//validate input
	v := validator.New()
	err := v.Struct(user.User{
		Firstname: " ", //bypass required check
		Lastname:  " ", //bypass required check
		Email:     email,
		Password:  pass,
	})

	if err != nil {
		res.Error = err.Error()
		res.Data = err.(validator.ValidationErrors)

		json.NewEncoder(w).Encode(res)
		return
	}

	//get db
	db, _ := db.Con()

	var result user.User

	//check for mail
	email_filter := bson.D{{"email", email}}

	//get collection
	err = db.Collection("users").FindOne(context.TODO(), email_filter).Decode(&result)

	if err != nil {
		res.Error = err.Error()
		//res.Data = result
		json.NewEncoder(w).Encode(res)
		return
	}

	match := CheckPasswordHash(pass, result.Password)

	if !match {
		res.Result = "Email and Password does not match"
		json.NewEncoder(w).Encode(res)
		return
	}

	//record token and login log
	token := RandString(100)

	var token_data = bson.D{
		{"user_id", result.Id},
		{"token", token},
		{"valid", true},
		{"created_at", time.Now()},
	}

	_, err = db.Collection("token_access").InsertOne(context.TODO(), token_data)

	if err != nil {
		res.Result = "Error Occurred when trying to store token"
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	result.Password = ""
	res.Result = "Login success"
	res.Data = bson.D{
		{"User", result},
		{"Token", token},
	}
	json.NewEncoder(w).Encode(res)
	return
}

func Register(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	pass := r.FormValue("password")
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")

	var res models.ResponseResult

	//prepare data
	a := user.User{
		Email:     email,
		Password:  pass,
		Firstname: firstname,
		Lastname:  lastname,
	}

	//validate input
	v := validator.New()
	err := v.Struct(a)

	if err != nil {
		res.Error = err.Error()
		res.Data = err.(validator.ValidationErrors)

		json.NewEncoder(w).Encode(res)
		return
	}

	//get db
	db, _ := db.Con()
	var result user.User

	//check for mail
	email_filter := bson.D{{"email", email}}

	//get collection
	err = db.Collection("users").FindOne(context.TODO(), email_filter).Decode(&result)

	if result.Email == email {
		res.Error = err.Error()
		res.Result = "Email already Exist"
		//res.Data = result
		json.NewEncoder(w).Encode(res)
		return
	}

	//create user name
	name := []string{firstname, lastname}
	username := strings.Join(name, ".")

	//check if username exist
	uname := UnameGenerator(username)

	//hash password
	hash, err := HashPassword(pass)

	if err != nil {
		res.Error = err.Error()
		res.Result = "error when trying to hash password"

		json.NewEncoder(w).Encode(res)
		return
	}

	var user_data = bson.D{
		{"firstname", firstname},
		{"lastname", lastname},
		{"username", uname},
		{"email", email},
		{"password", hash},
		{"created_at", time.Now()},
	}

	insertResult, err := db.Collection("users").InsertOne(context.TODO(), user_data)

	if err != nil {
		res.Result = "Error Occurred when trying to store data"
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Result = "Success creating account "
	res.Data = insertResult
	json.NewEncoder(w).Encode(res)
	return
}

//== Utillity ==
func UnameGenerator(username string) string {
	random := strconv.Itoa(rand.Intn(9999))

	name := []string{username, random}

	new := strings.Join(name, "")

	//get db
	db, _ := db.Con()

	//check if username exist
	uname_filter := bson.D{{"username", username}}

	var result user.User

	db.Collection("users").FindOne(context.TODO(), uname_filter).Decode(&result)

	for result.Username == new {
		db.Collection("users").FindOne(context.TODO(), uname_filter).Decode(&result)

		random := strconv.Itoa(rand.Intn(9999))

		name := []string{username, random}

		new = strings.Join(name, ".")
	}

	return new
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func RandString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
