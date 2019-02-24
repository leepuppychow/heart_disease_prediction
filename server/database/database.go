package database

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

var RedisConn redis.Conn

func init() {
	RedisConn = Connect()
}

func Connect() redis.Conn {
	c, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		log.Println("Error connecting to Redis", err)
	} else {
		log.Println("Connected to Redis successfully")
	}
	return c
}

func Exec(command string, args ...interface{}) interface{} {
	var reply interface{}
	var err error

	if command == "HGETALL" {
		reply, err = redis.StringMap(RedisConn.Do(command, args...))
	} else {
		reply, err = redis.String(RedisConn.Do(command, args...))
	}
	if err != nil {
		log.Println(err)
	}

	return reply
}
