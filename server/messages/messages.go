package messages

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	c "github.com/leepuppychow/heart_disease_prediction/server/csv_helpers"
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
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Successful request to Predict service", res)
	}
}

func Train(url, file string) {
	contents := c.GetCSVContents(file)
	log.Println(contents)
}
