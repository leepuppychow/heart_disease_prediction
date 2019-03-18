package messages

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

func SendTo(domain, port, msg string) (bool, error) {
	if msg == "" || domain == "" || port == "" {
		log.Printf("Missing information for sending message")
		return false, errors.New("Missing information (either domain, port, or message)")
	}

	url := fmt.Sprintf("http://%s:%s/%s", domain, port, msg)
	_, err := http.Get(url)

	if err != nil {
		log.Println("Error sending message to prediction service", err)
		return false, err
	}
	return true, nil
}

func Predict(url string) {
	csvFile, err := os.Open(file)
	r := csv.NewReader(csvFile)

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal("PREDICT ERROR", err)
	}

	log.Println(records)
}

func Train(url, file string) {
	csvFile, err := os.Open(file)
	r := csv.NewReader(csvFile)

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal("TRAIN ERROR", err)
	}

	log.Println(records)
}
