package main

import (
	"log"
	"net/http"

	h "github.com/leepuppychow/heart_disease_prediction/server/handlers"
	m "github.com/leepuppychow/heart_disease_prediction/server/messages"
)

func main() {
	startServer(":8000")
}

func startServer(port string) {
	m.UpdateCSV("./data/heart.csv") // Send CSV to all services when server first starts

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", h.IndexHandler)
	http.HandleFunc("/csv-load-form", h.CSVLoadForm)
	http.HandleFunc("/patients", h.NewPatientHandler)
	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
