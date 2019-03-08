package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/leepuppychow/heart_disease_prediction/prediction/database"
)

func TrainHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("HIT THE TRAIN HANDLER")
	fmt.Println(database.GetAllRows("dataList"))
}

func PredictHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("HIT THE PREDICT HANDLER")
}
