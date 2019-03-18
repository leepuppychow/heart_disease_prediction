package messages

import (
	"log"
	"net/http"
	"strings"

	c "github.com/leepuppychow/heart_disease_prediction/server/csv_helpers"
)

func Predict(row []string) error {
	url := "http://prediction:8080/predict"
	rowString := strings.Join(row, ",")
	_, err := http.Post(url, "text/csv", strings.NewReader(rowString))
	if err != nil {
		log.Println(err)
	}
	return err
}

func Train(file string) error {
	url := "http://prediction:8080/train"
	contents := c.OpenCSV(file)
	client := &http.Client{}
	_, err := client.Post(url, "text/csv", contents)
	if err != nil {
		log.Println("Error sending CSV to prediction", err)
	}
	return err
}
