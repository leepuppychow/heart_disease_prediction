package main

import (
	"log"
	"net/http"

	"github.com/leepuppychow/heart_disease_prediction/server/csv_loader"
	h "github.com/leepuppychow/heart_disease_prediction/server/handlers"
)

func main() {
	csv_loader.CsvToRedis()
	startServer(":8000")
}

func startServer(port string) {
	http.Handle("/", h.IndexHandler())
	http.HandleFunc("/patients", h.NewPatientHandler)
	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
