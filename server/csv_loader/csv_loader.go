package csv_loader

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"

	db "github.com/leepuppychow/heart_disease_prediction/server/database"
	"github.com/leepuppychow/heart_disease_prediction/server/messages"
)

func CsvToRedis() {
	db.DeleteList() // clear out Redis DB when server first starts
	csvFile, err := os.Open("./data/heart.csv")
	if err != nil {
		log.Println("Error loading CSV file", err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println("Error reading CSV line", err)
		}
		db.AddRow(strings.Join(line, ","))
	}

	messages.SendTo("prediction", "8080", "train")
}
