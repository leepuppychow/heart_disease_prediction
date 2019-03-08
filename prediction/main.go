package main

import (
	"log"
	"net/http"

	h "github.com/leepuppychow/heart_disease_prediction/prediction/handlers"
)

func main() {
	startPredictionService(":8080")
}

func startPredictionService(port string) {
	http.HandleFunc("/train", h.TrainHandler)
	http.HandleFunc("/predict", h.PredictHandler)
	log.Println("Prediction service running on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
