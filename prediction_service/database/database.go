package database

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

func Connect() redis.Conn {
	c, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		log.Println("Error connecting to Redis", err)
	} else {
		log.Println("Connected to Redis successfully")
	}
	return c
}

func Exec(conn redis.Conn, command string, args ...interface{}) interface{} {
	var reply interface{}
	var err error

	if command == "HGETALL" {
		reply, err = redis.StringMap(conn.Do(command, args...))
	} else {
		reply, err = redis.String(conn.Do(command, args...))
	}
	if err != nil {
		log.Println(err)
	}

	return reply
}
