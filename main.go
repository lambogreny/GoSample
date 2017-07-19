package main

import (
	"log"
	"net/http"

	"github.com/yuttasakcom/GoAPI/routes"
)

func main() {
	h := routes.Router()
	err := http.ListenAndServe(":8080", h)
	if err != nil {
		log.Println(err)
	}
}
