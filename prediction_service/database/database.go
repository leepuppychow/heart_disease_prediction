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

func GetAllRows() []string {
	c := Connect(5)
	defer c.Close()
	reply, err := redis.Strings(c.Do("LRANGE", "dataList", "0", "-1"))
	if err != nil {
		log.Println("Error getting all data", err)
	}
	log.Println("Successfully got all data")

	return reply
}
