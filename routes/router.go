package routes

import (
	"net/http"

	"github.com/yuttasakcom/GoAPI/middleware"
)

// Router provider
func Router() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/about", aboutHandler)

	mux.Handle("/staff", middleware.Chain(
		middleware.Authen("Bearer ABCD1234"),
		middleware.AllowRoles("staff"),
	)(http.HandlerFunc(staffHandler)))

	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index Page"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Page"))
}

func staffHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Staff!"))
}
