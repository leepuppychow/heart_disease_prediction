package main

import (
	"fmt"
	"log"
	"net/http"

	db "github.com/leepuppychow/heart_disease_prediction/server/database"
	h "github.com/leepuppychow/heart_disease_prediction/server/handler"
)

func main() {
	fmt.Println(db.Exec("HMSET", "hash", "name", "lee", "age", "30"))
	fmt.Println(db.Exec("HGETALL", "hash"))

	http.HandleFunc("/", h.IndexHandler)
	http.HandleFunc("/patients", h.NewPatientHandler)
	port := ":8000"
	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
