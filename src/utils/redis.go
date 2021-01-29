package utils

import (
	"github.com/go-redis/redis"
	"fmt"
)

var redisdb *redis.Client

func init() {
	//redis initial , setting can see redis.go opt.init()
	redisdb = redis.NewClient(&redis.Options{ //return Client
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	cmd := redisdb.Ping()
	if cmd.Err() != nil {
		panic(cmd.Err())
	}
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

func GetDel(key string) (jsonStr string) {
	cmd := redisdb.Get(key)
	if cmd.Val() != "" {
		jsonStr = cmd.Val()
		redisdb.Del(key)
	}
	return
}
