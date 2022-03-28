package main

import (
	"runtime"
	"fmt"
	"time"
)


func main(){
	printMessageUsage()

	var overall [][]int
	for i := 0; i<4; i++{
		a := make([]int, 0, 9999999)
		overall = append(overall, a)

		printMessageUsage()
		time.Sleep(time.Second)
	}
	
	overall = nil
	printMessageUsage()

	runtime.GC()
	printMessageUsage()
}


// PrintmessageUsage outputs the current, total and os memory being used. As well as the number
//of garage collection cycles completed
func printMessageUsage(){
	var m runtime.MemStats

	runtime.ReadMemStats(&m)
	//for ini an each see https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGc = %v MiB", bToMb(uint64(m.NumGC)))
}

func bToMb(b uint64) uint64{
	return b/1024/1024
}