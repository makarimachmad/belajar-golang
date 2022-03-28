package main

import (
    "fmt"
)

func print(ch chan bool) {
    fmt.Println("Printing from goroutine")
    ch <- true
}

func main() {
    ch := make(chan bool)
    go print(ch)
    <-ch
    fmt.Println("Printing from main")
}