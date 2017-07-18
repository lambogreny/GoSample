package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", indexHandler)

	http.Handle("/staff", chain(
		authen("ABC123"),
		allowRols("staff"),
	)(http.HandlerFunc(staffHandler)))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
	}
}

type middleware func(http.Handler) http.Handler

func chain(hs ...middleware) middleware {
	return func(h http.Handler) http.Handler {
		for i := len(hs); i > 0; i-- {
			h = hs[i-1](h)
		}
		return h
	}
}

func authen(token string) middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println(token)
			h.ServeHTTP(w, r)
		})
	}
}

func allowRols(roles ...string) middleware {
	allow := make(map[string]struct{})
	for _, role := range roles {
		allow[role] = struct{}{}
	}
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if _, ok := allow[r.Header.Get("Role")]; !ok {
				log.Println("Not Staff!")
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Index Page"))
}

func staffHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Staff"))
}
