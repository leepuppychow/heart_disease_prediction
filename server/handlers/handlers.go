package handlers

import (
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	c "github.com/leepuppychow/heart_disease_prediction/server/csv_helpers"
	"github.com/leepuppychow/heart_disease_prediction/server/messages"
)

func IndexHandler() http.Handler {
	return http.FileServer(http.Dir("static"))
}

func NewPatientHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		csvFile := c.OpenCSV("./data/heart.csv")
		defer csvFile.Close()

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
