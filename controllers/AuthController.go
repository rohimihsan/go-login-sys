package controllers

import (
	"encoding/json"
	"github.com/rohimihsan/go-login-sys/models"
	"github.com/rohimihsan/go-login-sys/models/user"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strings"
)

func Login(w http.ResponseWriter, r *http.Request) {
	//email := r.FormValue("email")
	pass := r.FormValue("password")
	hash := r.FormValue("hash")

	var res models.ResponseResult

	match := CheckPasswordHash(pass, hash)

	res.Data = match
	json.NewEncoder(w).Encode(res)
	return
}

func Register(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	pass := r.FormValue("password")
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")

	v := validator.New()

	a := user.User{
		Email:     email,
		Password:  pass,
		Firstname: firstname,
		Lastname:  lastname,
	}

	err := v.Struct(a)

	var res models.ResponseResult

	if err != nil {
		res.Error = err.Error()
		res.Data = err.(validator.ValidationErrors)

		json.NewEncoder(w).Encode(res)
		return
	}

	name := []string{firstname, lastname}

	username := strings.Join(name, ".")

	hash, err := HashPassword(pass)

	if err != nil {
		res.Error = err.Error()
		res.Result = "error when trying to hash password"

		json.NewEncoder(w).Encode(res)
		return
	}

	var user_data = bson.D{
		{"email", email},
		{"password", hash},
		{"username", username},
	}

	res.Data = user_data
	json.NewEncoder(w).Encode(res)
	return
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func UserProfile() {

}
