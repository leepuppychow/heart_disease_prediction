package main

import (
	"log"
	"net/http"
)

func main() {
	startPredictionService(":8080")
}

func startPredictionService(port string) {
	http.HandleFunc("/train", TrainHandler)
	http.HandleFunc("/predict", PredictHandler)
	log.Println("Prediction service running on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func TrainHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("HIT THE TRAIN HANDLER")
}

func PredictHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("HIT THE PREDICT HANDLER")
}
