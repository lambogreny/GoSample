package routes

import (
	"net/http"

	"os"

	"github.com/yuttasakcom/GoAPI/controllers"
)

// Router provider
func Router() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/-/", http.StripPrefix("/-", http.FileServer(noDir{http.Dir("static")})))
	mux.Handle("/", controllers.Home("index"))
	mux.Handle("/news/", controllers.News("newsID"))

	adminMux := http.NewServeMux()
	adminMux.HandleFunc("/login", controllers.Admin("login"))
	adminMux.HandleFunc("/list", controllers.Admin("list"))
	adminMux.HandleFunc("/create", controllers.Admin("create"))
	adminMux.HandleFunc("/edit", controllers.Admin("edit"))

	mux.Handle("/admin/", http.StripPrefix("/admin", onlyAdmin(adminMux)))

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

func onlyAdmin(h http.Handler) http.Handler {
	return h
}
