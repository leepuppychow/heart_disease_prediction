package main

import (
	"fmt"

	db "github.com/leepuppychow/heart_disease_prediction/database"
)

func main() {
	fmt.Println("HELLO THERE, WHERE IS MY DATA???")
	redisConn := db.Connect()

	for i := 0; i < 20; i++ {
		fmt.Println(db.Exec(redisConn, "HGET", "hash", "age"))
	}
}
