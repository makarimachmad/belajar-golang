package main

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

func ConfigRedis() *redis.Client{
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 2,
	})
	fmt.Println("terhubung redis")
	return client
}