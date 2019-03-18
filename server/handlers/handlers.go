package handlers

import (
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/leepuppychow/heart_disease_prediction/server/messages"
)

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
	csvFile, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer csvFile.Close()
	log.Println("File opened successfully")

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	err := writer.Write(row)
	if err != nil {
		log.Println(err)
	}
	log.Println("Row Added:", row)

	// If row added count exceeds some threshold then:
	messages.Train("http://prediction:8080", file)
}

func MakePrediction(row []string) {
	messages.Predict("http://prediction:8080")
}

func CSVLoadForm(w http.ResponseWriter, r *http.Request) {
	// Saves CSV file to the data directory
	file, header, err := r.FormFile("csvFile")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	out, err := os.Create(filepath.Join("./data", header.Filename))
	if err != nil {
		log.Println("Unable to create the file for writing")
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		log.Println("Unable to copy contents to the file")
	}

	log.Println("File loaded successfully", header.Filename)
	http.Redirect(w, r, "/", http.StatusFound)
}
