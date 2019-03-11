package handlers

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"

	"github.com/leepuppychow/heart_disease_prediction/server/messages"
)

func IndexHandler() http.Handler {
	return http.FileServer(http.Dir("static"))
}

func NewPatientHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		csvFile, err := os.OpenFile("./data/heart.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer csvFile.Close()
		if err != nil {
			log.Println(err)
		}

		writer := csv.NewWriter(csvFile)
		defer writer.Flush()

		age := r.FormValue("age")
		sex := r.FormValue("sex")
		cp := r.FormValue("cp")
		trestbps := r.FormValue("trestbps")
		chol := r.FormValue("chol")
		fbs := r.FormValue("fbs")
		hasHeartDisease := r.FormValue("hasHeartDisease")

		if hasHeartDisease == "" {
			messages.SendTo("prediction", "8080", "predict")
		} else {
			row := []string{age, sex, cp, trestbps, chol, fbs, hasHeartDisease}
			err := writer.Write(row)
			if err != nil {
				log.Println(err)
			}
			log.Println("Row Added:", row)
			messages.SendTo("prediction", "8080", "train")
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
