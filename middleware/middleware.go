package middleware

import (
	"context"
	"github.com/rohimihsan/go-login-sys/config/db"
	"github.com/rohimihsan/go-login-sys/models"
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

func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//get token
		reqToken := r.Header.Get("Authorization")

		if reqToken == "" {
			w.Write([]byte("Unauthorized"))
			return
		}

		splitToken := strings.Split(reqToken, "Bearer")

		if len(splitToken) != 2 {
			w.Write([]byte("Invalid Token Format"))
			return
		}

		reqToken = strings.TrimSpace(splitToken[1])

		//get db
		db, _ := db.Con()

		var result models.Token_access

		//check for token
		email_filter := bson.D{{"token", reqToken}}

		//get collection
		err := db.Collection("token_access").FindOne(context.TODO(), email_filter).Decode(&result)

		if err != nil || result.Valid == false {
			w.Write([]byte("Unauthorized"))
			return
		}

		//reqToken = splitToken[1]

		next.ServeHTTP(w, r)
	})
}
