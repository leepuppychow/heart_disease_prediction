package csv_helpers

import (
	"log"
	"os"
)

func OpenCSV(file string) *os.File {
	csvFile, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	log.Println("File opened successfully:", file)
	return csvFile
}

