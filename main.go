package main

import (
	"log"
	"net/http"

	"github.com/yuttasakcom/GoAPI/models"
	"github.com/yuttasakcom/GoAPI/routes"
)

const (
	port     = ":8080"
	mongoURL = "mongodb://127.0.0.1:27017"
)

func main() {
	h := routes.Router()
	err := models.Init(mongoURL)
	if err != nil {
		log.Fatalf("can not init model; %v", err)
	}
	http.ListenAndServe(":8080", h)
	// err := http.ListenAndServe(":8080", h)
	// if err != nil {
	// 	log.Println(err)
	// }
}
