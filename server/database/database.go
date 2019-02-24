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

func Exec(command string, args ...interface{}) interface{} {
	var reply interface{}
	var err error
	c := Connect()
	defer c.Close()

	if command == "HGETALL" {
		reply, err = redis.StringMap(c.Do(command, args...))
	} else {
		reply, err = redis.String(c.Do(command, args...))
	}
	if err != nil {
		log.Println(err)
	}

	return reply
}
