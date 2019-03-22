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
		// "http://statistics:8111/stats",
	}
	for _, url := range urls {
		go func(url string) {
			contents := c.OpenCSV(filepath)
			defer contents.Close()

			client := &http.Client{}
			_, err := client.Post(url, "text/csv", contents)
			if err != nil {
				log.Println("Error sending CSV", err)
			}
		}(url)
	}
	err := backup.SaveToS3(filepath)
	if err == nil {
		log.Println("Success saving file to S3")
	}
}

func Predict(row []string) error {
	url := "http://prediction:8080/predict"
	rowString := strings.Join(row, ",")
	_, err := http.Post(url, "text/csv", strings.NewReader(rowString))
	if err != nil {
		log.Println(err)
	}
	return err
}
