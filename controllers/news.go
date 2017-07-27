package controllers

import (
	"net/http"

	"github.com/yuttasakcom/GoAPI/models"
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
	id := r.URL.Path[1:]
	n, err := models.GetNews(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	views.NewsID(w, n)
}
