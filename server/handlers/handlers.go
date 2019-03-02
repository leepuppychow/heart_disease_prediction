package handlers

import (
	"net/http"

	db "github.com/leepuppychow/heart_disease_prediction/server/database"
)

func IndexHandler() http.Handler {
	return http.FileServer(http.Dir("static"))
}

func NewPatientHandler(w http.ResponseWriter, r *http.Request) {
	age := r.FormValue("age")
	gender := r.FormValue("gender")
	cp := r.FormValue("cp")
	row := age + "," + gender + "," + cp
	db.AddRow(row)
	http.Redirect(w, r, "/", http.StatusFound)
}
