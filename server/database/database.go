package database

import (
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

func Connect(attempts int) redis.Conn {
	c, err := redis.Dial("tcp", "redis:6379")
	for i := 1; i < attempts; i++ {
		if err == nil {
			return c
		} else {
			time.Sleep(200 * time.Millisecond)
			c, err = redis.Dial("tcp", "redis:6379")
		}
	}
	log.Fatal("Unable to connect to Redis", err)
	return nil
}

// Try saving each row same as CSV format in Redis
// Do the data parsing in Go

func AddRow(list, row string) {
	c := Connect(5)
	defer c.Close()

	reply, err := c.Do("LPUSH", list, row)
	if err != nil {
		log.Println("Error adding row", list, err)
	}
	log.Println("Successfully added row", reply)
}

func GetAllRows(list string) []string {
	c := Connect(5)
	defer c.Close()
	reply, err := redis.Strings(c.Do("LRANGE", list, "0", "-1"))
	if err != nil {
		log.Println("Error getting all data", err)
	}
	log.Println("Successfully got all data")

	return reply
}

func DataCount(list string) int64 {
	c := Connect(5)
	defer c.Close()

	reply, err := redis.Int64(c.Do("LLEN", list))
	if err != nil {
		log.Println("Error getting count of", list, err)
	}
	return reply
}

func DeleteList(list string) string {
	c := Connect(5)
	defer c.Close()
	reply, err := redis.String(c.Do("DEL", list))
	if err != nil {
		log.Println("Error deleting all data", err)
	}
	log.Println("Successfully deleted all data")

	return reply
}
