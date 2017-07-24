package controllers

import (
	"net/http"

	"github.com/yuttasakcom/GoAPI/models"
	"github.com/yuttasakcom/GoAPI/views"
)

// Admin controller
func Admin(p string) http.HandlerFunc {
	switch p {
	case "login":
		return http.HandlerFunc(login)
	case "list":
		return http.HandlerFunc(list)
	case "create":
		return http.HandlerFunc(create)
	case "edit":
		return http.HandlerFunc(edit)
	default:
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.NotFound(w, r)
			return
		})
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	views.AdminLogin(w, r)
}

func list(w http.ResponseWriter, r *http.Request) {
	views.AdminList(w, r)
}

func create(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		n := models.News{
			Title:  r.FormValue("title"),
			Detail: r.FormValue("detail"),
		}

		models.CreateNews(n)

		http.Redirect(w, r, "/admin/create", http.StatusSeeOther)
		return
	}
	views.AdminCreateNews(w, nil)
}

func edit(w http.ResponseWriter, r *http.Request) {
	views.AdminEditNews(w, r)
}
