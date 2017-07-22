package routes

import (
	"net/http"

	"os"

	"github.com/yuttasakcom/GoAPI/controllers"
	"github.com/yuttasakcom/GoAPI/middleware"
)

// Router provider
func Router() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/-/", http.StripPrefix("/-", http.FileServer(noDir{http.Dir("static")})))
	mux.Handle("/", controllers.Home("index"))
	mux.HandleFunc("/about", aboutHandler)
	mux.Handle("/staff", middleware.Chain(
		middleware.Authen("Bearer ABCD1234"),
		middleware.AllowRoles("staff"),
	)(http.HandlerFunc(staffHandler)))
	return mux
}

type noDir struct {
	http.Dir
}

func (d noDir) Open(name string) (http.File, error) {
	f, err := d.Dir.Open(name)
	if err != nil {
		return nil, err
	}
	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}
	if stat.IsDir() {
		return nil, os.ErrNotExist
	}
	return f, nil
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Page"))
}

func staffHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Staff!"))
}
