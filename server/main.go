package main

import (
	"log"
	"net/http"

	h "github.com/leepuppychow/heart_disease_prediction/server/handler"
)

func main() {
	http.HandleFunc("/", h.IndexHandler)
	http.HandleFunc("/patients", h.NewPatientHandler)
	port := ":8000"
	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
