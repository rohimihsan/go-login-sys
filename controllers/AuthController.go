package controllers

import (
	"encoding/json"
	"github.com/rohimihsan/go-login-sys/models"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
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

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Register(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	pass := r.FormValue("password")

	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")

	name := []string{firstname, lastname}

	username := strings.Join(name, ".")

	var res models.ResponseResult

	hash, err := HashPassword(pass)

	if err != nil {
		res.Error = err.Error()
		res.Result = "error when trying to hash password"

		json.NewEncoder(w).Encode(res)
		return
	}

	var user_data = bson.D{
		{"email", email},
		{"password", pass},
		{"hashed", hash},
	}

	res.Data = user_data
	json.NewEncoder(w).Encode(res)
	return
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func UserProfile() {

}
