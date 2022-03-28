package main

import "fmt"

func main(){

	num := 4

	switch num{
	case 1 :
		fmt.Println("kurang")
	case 4 :
		fmt.Println(num)
	default:
		fmt.Println("default")
	}

	for i:=1; i<=100; i++{
		fmt.Println(i)
	}

	i:=1
	for i <=100{
		fmt.Println(i)
		i++
	}

	nama:="achmad makarim widyanto"
	for index, letter:=range nama{
		fmt.Println("index: ",index, "letter: ", string(letter))
	}
	for _, letter:=range nama{
		fmt.Println("letter: ", string(letter))
	}

	//quiz
	for index, letter:=range nama{
		if index%2 == 0{
			fmt.Println("index: ",index, "letter: ", string(letter))
		}
		letterstring := string(letter)
		switch letterstring{
		case "a","i","u","e","o":
			fmt.Println(letterstring)
		default:
			fmt.Println("bukan")
		}
	}
}