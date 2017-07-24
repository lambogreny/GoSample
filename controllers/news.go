package controllers

import (
	"net/http"

	"github.com/yuttasakcom/GoAPI/views"
)

// News controllers
func News(p string) http.Handler {
	switch p {
	case "newsID":
		return http.HandlerFunc(newsID)
	default:
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.NotFound(w, r)
			return
		})
	}
}

func newsID(w http.ResponseWriter, r *http.Request) {
	views.NewsID(w, r)
}
