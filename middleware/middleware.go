package middleware

import (
	"context"
	"github.com/rohimihsan/go-login-sys/config/db"
	"github.com/rohimihsan/go-login-sys/models"
	"github.com/rohimihsan/go-login-sys/models/user"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strings"
)

func MiddlewareAllowOnlyGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.Write([]byte("Only GET is allowed"))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func CorsByPass(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		whitelist := "http://localhost:3000";
		w.Header().Set("Access-Control-Allow-Origin", whitelist)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Expose-Headers", "Session-Token")

		next.ServeHTTP(w, r)
	})
}

func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//get token
		reqToken := r.Header.Get("Authorization")

		if reqToken == "" {
			w.WriteHeader(http.StatusUnauthorized)

			w.Write([]byte("Unauthorized"))
			return
		}

		splitToken := strings.Split(reqToken, "Bearer")
		if len(splitToken) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid Token Format"))
			return
		}

		reqToken = strings.TrimSpace(splitToken[1])
		//reqToken = splitToken[1]

		//get db
		db, _ := db.Con()

		var result models.Token_access

		//check for token
		token_filter := bson.D{{"token", reqToken}}

		//get collection
		err := db.Collection("token_access").FindOne(context.TODO(), token_filter).Decode(&result)

		if err != nil || result.Valid == false {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}

		//get user
		var user_result user.User

		uid_filter := bson.D{{"_id", result.User_id}}

		err = db.Collection("users").FindOne(context.TODO(), uid_filter).Decode(&user_result)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error Occurred, Unauthorized"))
			return
		}

		//remove password
		user_result.Password = ""

		//set user
		ctx := context.WithValue(r.Context(), "user", user_result)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
