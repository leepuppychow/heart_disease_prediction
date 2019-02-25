package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	db "github.com/leepuppychow/heart_disease_prediction/server/database"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/index.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, "")
}

func NewPatientHandler(w http.ResponseWriter, r *http.Request) {
	age := r.FormValue("age")
	gender := r.FormValue("gender")
	cp := r.FormValue("cp")
	row := age + "," + gender + "," + cp

	db.AddRow(row)
	fmt.Println(db.GetAllRows())
	fmt.Println(db.DataCount())

	http.Redirect(w, r, "/", http.StatusFound)
}
