package main

import (
    "fmt"
    // "sync"
    "time"
)

// func worker(id int, wg *sync.WaitGroup) {

//     defer wg.Done()

//     fmt.Printf("Worker %d starting\n", id)

//     time.Sleep(time.Second)
//     fmt.Printf("Worker %d done\n", id)
// }

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func say() {
	fmt.Println("Halo ini berasal dari goroutine baru")
}

func main() {

    // var wg sync.WaitGroup

    // for i := 1; i <= 5; i++ {
    //     wg.Add(1)
    //     go worker(i, &wg)
    // }

    // wg.Wait()
    f("direct")
    go f("goroutine")

    go func(msg string){
        fmt.Println(msg)
    }("going")

    time.Sleep(time.Second)
    fmt.Println("done")

    go say()
	time.Sleep(1 * time.Second)
	fmt.Println("Ini Fungsi Utama")

    go say()
	//time.Sleep(1 * time.Second)
	fmt.Println("Ini Fungsi Utama")
}