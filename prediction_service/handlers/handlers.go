package handlers

import (
	"log"
	"net/http"
)

func TrainHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("HIT THE TRAIN HANDLER")
}

func PredictHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("HIT THE PREDICT HANDLER")
}
