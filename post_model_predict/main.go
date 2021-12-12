package main

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "211.71.76.189:6379",
		DB:   0, // use default DB
	})
	ctx := context.Background()
	val, err := rdb.Get(ctx, "test").Result()
	if err == redis.Nil {
		log.Fatalf("key not exits")
	} else if err != nil {
		log.Printf("err:%v", err)
	} else {
		log.Println(val)
	}

}
