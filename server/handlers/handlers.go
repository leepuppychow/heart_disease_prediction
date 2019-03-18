package handlers

import (
	"log"
	"net/http"

	c "github.com/leepuppychow/heart_disease_prediction/server/csv_helpers"
	"github.com/leepuppychow/heart_disease_prediction/server/messages"
	"github.com/leepuppychow/heart_disease_prediction/server/models"
)

var RowsAdded int

func IndexHandler() http.Handler {
	return http.FileServer(http.Dir("static"))
}

func NewPatientHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		p := models.HeartDiseasePatient{
			Age:                   r.FormValue("age"),
			Sex:                   r.FormValue("sex"),
			ChestPainType:         r.FormValue("cp"),
			RestingBloodPress:     r.FormValue("trestbps"),
			SerumCholesterol:      r.FormValue("chol"),
			FastingBP:             r.FormValue("fbs"),
			ExerciseInducedAngina: r.FormValue("angina"),
			HasHeartDisease:       r.FormValue("hasHeartDisease"),
		}
		row := p.DataRow()
		if p.HasHeartDisease == "" {
			MakePrediction(row)
		} else {
			SaveNewDataPoint(row)
		}
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func MakePrediction(row []string) {
	err := messages.Predict(row)
	if err == nil {
		log.Println("Successful POST /predict (prediction service)")
	}
}

func SaveNewDataPoint(row []string) {
	file := "./data/heart.csv"
	err := c.AppendToCSV(file, row)
	if err == nil {
		RowsAdded++
	}
	if RowsAdded > 0 {
		err = messages.Train(file)
		if err == nil {
			log.Println("Successful POST /train (prediction service)")
		}
	}
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
