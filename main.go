package main

import (
	"log"
	"github.com/keremavci/go-helm-cloudbuild/handlers"
	"net/http"
)

func main(){
	log.Println("Starting...")
	log.Fatal(http.ListenAndServe(":8080",handlers.Handlers()))
}