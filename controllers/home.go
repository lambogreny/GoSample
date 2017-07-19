package controllers

import (
	"net/http"
)

// Home controller
func Home(p string) http.Handler {
	switch p {
	case "index":
		// This call Services
		return http.HandlerFunc(index)
	default:
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.NotFound(w, r)
			return
		})
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(`
		<!doctype html>
		<link rel=stylesheet href=/-/css/style.css>
		<title>GoAPI</title>
		<p class=red>Home Page</p>
	`))
}
