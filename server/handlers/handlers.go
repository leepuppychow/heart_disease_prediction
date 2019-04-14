package handlers

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	c "github.com/leepuppychow/heart_disease_prediction/server/csv_helpers"
	"github.com/leepuppychow/heart_disease_prediction/server/messages"
	"github.com/leepuppychow/heart_disease_prediction/server/models"
)

var RowsAdded int

type Page struct {
	Title string
	// Score      string
	// Prediction string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	p := Page{
		Title: "Heart Disease Prediction Index",
	}
	t, err := template.ParseFiles("./static/index.html")
	if err != nil {
		log.Println("Error parsing HTML file")
	}
	t.Execute(w, p)
}

func NewPatientHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var p models.HeartDiseasePatient
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			log.Println("Error decoding request body into struct", err)
		}
		row := p.DataRow()
		prediction := ""
		if p.HasHeartDisease == "" {
			prediction = MakePrediction(row)
		} else {
			SaveNewDataPoint(row)
		}
		w.Write([]byte(prediction))
	}
}

func MakePrediction(row []string) string {
	resp, err := messages.Predict(row)
	if err == nil {
		log.Println("Successful POST /predict (prediction service)")
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading Response body", err)
	}
	bodyString := string(bodyBytes)
	log.Println(bodyString)
	return bodyString
}

func SaveNewDataPoint(row []string) {
	file := "./data/heart.csv"
	err := c.AppendToCSV(file, row)
	if err == nil {
		RowsAdded++
	}
	// TODO: change this eventually (either percentage of CSV file or set number of rows)
	if RowsAdded > 1 {
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
