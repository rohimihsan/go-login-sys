package middleware

import "net/http"

func MiddlewareAllowOnlyGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.Write([]byte("Only GET is allowed"))
			return
		}

		next.ServeHTTP(w, r)
	})
}
//
//func MiddlewareAuth(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//
//		next.ServeHTTP(w, r)
//	})
//}
