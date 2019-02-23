package main

import (
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func main() {
	http.HandleFunc("/", IndexHandler)
	port := ":8000"
	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
