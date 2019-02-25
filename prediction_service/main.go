package main

import (
	"fmt"

	db "github.com/leepuppychow/heart_disease_prediction/prediction_service/database"
)

func main() {
	fmt.Println("HELLO THERE, WHERE IS MY DATA???")

	for i := 0; i < 5; i++ {
		fmt.Println(db.Exec("HGET", "hash", "name"))
	}
}
