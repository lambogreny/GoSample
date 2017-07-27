package main

import (
	"log"
	"net/http"

	"github.com/yuttasakcom/GoAPI/models"
	"github.com/yuttasakcom/GoAPI/routes"
)

const (
	port     = ":8080"
	mongoURL = "mongodb://0.0.0.0:27017"
)

func main() {
	h := routes.Router()
	err := models.Init(mongoURL)
	if err != nil {
		log.Fatalf("can not init model; %v", err)
	}
	err = http.ListenAndServe(port, h)

	if err != nil {
		log.Println(err)
	}
}
