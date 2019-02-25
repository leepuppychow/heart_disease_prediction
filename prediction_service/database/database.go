package database

import (
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

func Connect() redis.Conn {
	c, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		log.Println("Error connecting to Redis, trying again...", err)
		time.Sleep(200 * time.Millisecond)
		return Connect()
	} else {
		log.Println("Connected to Redis successfully")
	}
	return c
}

// Try saving each row same as CSV format in Redis
// Do the data parsing in Go

func AddRow(row string) {
	c := Connect()
	defer c.Close()

	reply, err := c.Do("LPUSH", "dataList", row)
	if err != nil {
		log.Println("Error adding row to dataList", err)
	}
	log.Println("Successfully added row", reply)
}

func GetAllRows() []string {
	c := Connect()
	defer c.Close()
	reply, err := redis.Strings(c.Do("LRANGE", "dataList", "0", "-1"))
	if err != nil {
		log.Println("Error getting all data", err)
	}
	log.Println("Successfully got all data")

	return reply
}

func DataCount() int64 {
	c := Connect()
	defer c.Close()

	reply, err := redis.Int64(c.Do("LLEN", "dataList"))
	if err != nil {
		log.Println("Error getting dataList length", err)
	}
	return reply
}
