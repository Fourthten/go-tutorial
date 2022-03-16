package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func newRedisClient(redisHost string, redisPassowrd string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPassowrd,
		DB:       0,
	})
	return client
}

func main() {
	var redisHost = "localhost:6379"
	var redisPassowrd = ""

	rdb := newRedisClient(redisHost, redisPassowrd)
	fmt.Println("redis client:", rdb)

	key := "key-1"
	data := "Hello Redis"
	ttl := time.Duration(3) * time.Second

	// store data using SET command
	op1 := rdb.Set(context.Background(), key, data, ttl)
	if err := op1.Err(); err != nil {
		fmt.Printf("unable to set data: %v", err)
		return
	}
	log.Println("Set operation success")

	// time.Sleep(time.Duration(4) * time.Second)

	// get data
	op2 := rdb.Get(context.Background(), key)
	if err := op2.Err(); err != nil {
		fmt.Printf("unable to GET data: %v", err)
		return
	}
	res, err := op2.Result()
	if err != nil {
		fmt.Printf("unable to get data: %v", err)
		return
	}
	log.Println("Get operation success. result: ", res)
}

// https://redis.io/topics/quickstart
