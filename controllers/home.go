package controllers

import (
	"net/http"

	"github.com/yuttasakcom/GoAPI/models"
	"github.com/yuttasakcom/GoAPI/views"
)

// Home controller
func Home(p string) http.Handler {
	switch p {
	case "index":
		return http.HandlerFunc(index)
	default:
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.NotFound(w, r)
			return
		})
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	list, err := models.ListNews()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	views.Index(w, &views.IndexData{
		List: list,
	})
}
