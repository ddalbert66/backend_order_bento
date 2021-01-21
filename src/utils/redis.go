package utils

import (
	"github.com/go-redis/redis"
	"fmt"
)

var redisdb *redis.Client

func init() {
	redisdb = redis.NewClient(&redis.Options{ //return Client
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
}

func GetRedisDb() *redis.Client {
	return redisdb
}

func CheckExist(key string) bool {
	cmd := redisdb.Get(key)
	if cmd.Err() == redis.Nil {
		fmt.Println(cmd.Err())
		return false
	}
	if cmd.Val() != "" {
		fmt.Println(cmd.Val())
		return true
	} else {
		fmt.Println(cmd.Val())
		return false
	}
}
