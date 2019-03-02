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
	http.Get(url)
	return true, nil
}
