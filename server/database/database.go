package database

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

func Connect() redis.Conn {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Println("Error connecting to Redis", err)
	}
	return c
}

func Exec(conn redis.Conn, command string, args ...interface{}) string {
	reply, err := redis.String(conn.Do(command, args...))
	if err != nil {
		log.Println(err)
	}
	return reply
}
