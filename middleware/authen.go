package middleware

import (
	"log"
	"net/http"
)

// Authen middleware
func Authen(token string) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("Authen Pass")
			h.ServeHTTP(w, r)
		})
	}
}
