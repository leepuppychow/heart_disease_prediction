package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
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

	fmt.Println(age, gender)
	http.Redirect(w, r, "/", http.StatusFound)
}
