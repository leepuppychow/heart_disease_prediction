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
			Age:                      r.FormValue("age"),
			Sex:                      r.FormValue("sex"),
			ChestPainType:            r.FormValue("cp"),
			RestingBloodPress:        r.FormValue("trestbps"),
			SerumCholesterol:         r.FormValue("chol"),
			FastingBP:                r.FormValue("fbs"),
			RestECG:                  r.FormValue("restecg"),
			MaxHR:                    r.FormValue("thalach"),
			ExerciseInducedAngina:    r.FormValue("exang"),
			STDepressionWithExercise: r.FormValue("oldpeak"),
			SlopeSTSegment:           r.FormValue("slope"),
			NumberOfVesselsFlouro:    r.FormValue("ca"),
			Thal:                     r.FormValue("thal"),
			HasHeartDisease:          r.FormValue("hasHeartDisease"),
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

	// TODO: change this eventually (either percentage of CSV file or set number of rows)
	if RowsAdded > 0 {
		messages.UpdateCSV(file)
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
