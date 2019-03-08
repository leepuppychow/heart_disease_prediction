package main

import (
	"log"
	"net/http"

	"github.com/leepuppychow/heart_disease_prediction/visualization/histogram"
)

func main() {
	startVisualizationService(":8888")
}

func startVisualizationService(port string) {
	http.HandleFunc("/histogram", histogram.Histogram)
	log.Println("Visualization service running on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
