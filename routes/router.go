package routes

import (
	"net/http"

	"github.com/yuttasakcom/GoAPI/controllers"
	"github.com/yuttasakcom/GoAPI/middleware"
)

// Router provider
func Router() http.Handler {
	// Serve mux
	mux := http.NewServeMux()

	// Serve file assets
	mux.Handle("/-/", http.StripPrefix("/-", http.FileServer(http.Dir("assets"))))

	// Home index
	mux.Handle("/", controllers.Home("index"))

	mux.HandleFunc("/about", aboutHandler)

	mux.Handle("/staff", middleware.Chain(
		middleware.Authen("Bearer ABCD1234"),
		middleware.AllowRoles("staff"),
	)(http.HandlerFunc(staffHandler)))

	return mux
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Page"))
}

func staffHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Staff!"))
}
