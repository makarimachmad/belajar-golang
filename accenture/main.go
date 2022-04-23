package main

import "fmt"

func main() {
	fmt.Println(Kalkulasi("147"))
}

func Kalkulasi(nilai string) int {

	baru := 0

	if len(nilai) != 1 {
		for _, j := range nilai {
			fmt.Println(j)
			baru = baru + int(j)
		}
	}
	return baru
}