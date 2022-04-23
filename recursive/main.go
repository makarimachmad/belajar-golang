package main

import "fmt"
var nilai int
func RecursiveInt(a int) int {
	if a < 2 {
		return a
	}
	nilai = nilai + 1
	recursive := RecursiveInt(a-1)
	// fmt.Println("recursive: ", recursive)
	recursiveln2 := RecursiveInt(a-2)
	// fmt.Println("recursive ln2: ", recursiveln2)
	return recursive + recursiveln2
}
//recursive 1 -> 1 ; karena kena di return a duluan
//recursive 2 -> 1 ; 1 + 0
//recursive 3 -> 1 + 1
//recursive 4 -> 3(1+1) + 2(1+0) = 3
//recursive 5 -> 4(3(2(1+0)+0)+3(2(1+0)+1), 3(2(1+0)+1), 2(1+0), 1+0 = 5
//recursive 6 -> 5(4(3(2(1+0)+))+2(1+0))+3(2(1+0)+0)) + 4(3(2(1+0)+0)+2(1+0)), 4(3(2(1+0)+0)+2(1+0)) + 3(2(1+0)+0), 3(2(1+0)+0) + 2(1+0)+0, 2(1+0)+0 + 0

func main() {
	fmt.Println(RecursiveInt(6))
}