package handlers

import (
	"log"
	"net/http"
	"strings"

	db "github.com/leepuppychow/heart_disease_prediction/server/database"
	"github.com/leepuppychow/heart_disease_prediction/server/messages"
)

func IndexHandler() http.Handler {
	return http.FileServer(http.Dir("static"))
}

func NewPatientHandler(w http.ResponseWriter, r *http.Request) {
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
		row := strings.Join([]string{age, sex, cp, trestbps, chol, fbs, hasHeartDisease}, ",")
		db.AddRow(row)
		log.Println("Row Added:", row)
		messages.SendTo("prediction", "8080", "train")
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
