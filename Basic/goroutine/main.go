package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 100; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func siapa(s string) {
	for i := 0; i < 100; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	t := time.Now()
	go siapa("world")
	go say("hello")
	t2 := time.Now()
	ellapsed := t2.Sub(t)
	fmt.Println(ellapsed)
}
