package messages

import (
	"errors"
	"fmt"
	"log"
	"net/http"
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
