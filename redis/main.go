package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func ConfigRediss() *redis.Client{
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 2,
	})
	fmt.Println("terhubung redis")
	return client
}

func main() {
	rdb := ConfigRediss()

	key := "key-1"
	data := "Halo redis"
	ttl := time.Duration(3) * time.Second

	//insert data to DB redis
	op1 := rdb.Set(context.Background(), key, data, ttl)
	if err := op1.Err(); err != nil{
		fmt.Printf("unable to set data error: %v", err)
		return
	}
	fmt.Println("set data berhasil")

	//get data from DB redis
	op2 := rdb.Get(context.Background(), key)
	if err := op2.Err(); err != nil{
		fmt.Printf("unable to get data. error %v", err)
		return
	}
	res, err := op2.Result()
	if err != nil{
		fmt.Printf("unable to get data. error %v", err)
		return
	}
	fmt.Println("get operation success. hasil: ", res)

}