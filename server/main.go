package main

import (
	"log"
	"net/http"

	h "github.com/leepuppychow/heart_disease_prediction/server/handlers"
)

func main() {
	startServer(":8000")
}

func startServer(port string) {
	http.Handle("/", h.IndexHandler())
	http.HandleFunc("/patients", h.NewPatientHandler)
	http.HandleFunc("/csv-load-form", h.CSVLoadForm)
	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
