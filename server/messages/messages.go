package messages

import (
	"log"
	"net/http"
	"strings"

	"github.com/leepuppychow/heart_disease_prediction/server/backup"
	c "github.com/leepuppychow/heart_disease_prediction/server/csv_helpers"
)

func UpdateCSV(filepath string) {
	urls := []string{
		"http://prediction:8080/train",
		"http://visualization:8888/histograms",
		"http://statistics:8111/stats",
	}
	go backup.SaveToS3(filepath)
	for _, url := range urls {
		go func(url string) {
			contents := c.OpenCSV(filepath)
			defer contents.Close()

			client := &http.Client{}
			_, err := client.Post(url, "text/csv", contents)

			// TODO: Get response from statistics endpoint and save that data to display in index.html
			if err != nil {
				log.Println("Error sending CSV", err)
			}
		}(url)
	}
}

func Predict(row []string) (*http.Response, error) {
	url := "http://prediction:8080/predict"
	rowString := strings.Join(row, ",")
	rowString = rowString[:len(rowString)-1] // remove last field from array (empty target value)
	resp, err := http.Post(url, "text/csv", strings.NewReader(rowString))
	if err != nil {
		log.Println(err)
	}
	return resp, err
}
