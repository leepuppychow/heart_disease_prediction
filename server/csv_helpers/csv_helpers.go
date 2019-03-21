package csv_helpers

import (
	"encoding/csv"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

func OpenCSV(file string) *os.File {
	csvFile, err := os.Open(file)
	if err != nil {
		log.Println("Error opening CSV file", err)
	}
	return csvFile
}

func AppendToCSV(file string, row []string) error {
	csvFile, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Error opening file", err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	err = writer.Write(row)
	if err != nil {
		log.Println("Error appending to CSV", err)
	}
	return err
}

func SaveCSV(file multipart.File, fileName string) error {
	out, err := os.Create(filepath.Join("./data", fileName))
	if err != nil {
		log.Println("Unable to create the file for writing")
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		log.Println("Unable to copy contents to the file")
	}
	return err
}
