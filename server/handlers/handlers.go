package handlers

import (
	"log"
	"net/http"

	c "github.com/leepuppychow/heart_disease_prediction/server/csv_helpers"
	"github.com/leepuppychow/heart_disease_prediction/server/messages"
)

var RowsAdded int

func IndexHandler() http.Handler {
	return http.FileServer(http.Dir("static"))
}

func NewPatientHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		hasHeartDisease := r.FormValue("hasHeartDisease")
		age := r.FormValue("age")
		sex := r.FormValue("sex")
		cp := r.FormValue("cp")
		trestbps := r.FormValue("trestbps")
		chol := r.FormValue("chol")
		fbs := r.FormValue("fbs")
		row := []string{age, sex, cp, trestbps, chol, fbs, hasHeartDisease}

		if hasHeartDisease == "" {
			MakePrediction(row)
		} else {
			SaveNewDataPoint(row)
		}
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func SaveNewDataPoint(row []string) {
	file := "./data/heart.csv"
	err := c.AppendToCSV(file, row)
	if err == nil {
		RowsAdded++
	}
	if RowsAdded > 1 {
		messages.Train("http://prediction:8080", file)
	}
}

func MakePrediction(row []string) {
	messages.Predict("http://prediction:8080")
}

func CSVLoadForm(w http.ResponseWriter, r *http.Request) {
	// Saves CSV file to the data directory
	file, header, err := r.FormFile("csvFile")
	if err != nil {
		log.Println("Error getting CSV file from form", err)
	}
	defer file.Close()
	err = c.SaveCSV(file, header.Filename)
	if err == nil {
		log.Println("File loaded successfully", header.Filename)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
