package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	content, err := os.Open("input.in")
	if err != nil{
		fmt.Println("enggak ada filenya")
	}

	dict := map[int][]int{}
	var bantu, dari, sampai, jarak []int

	nilai := bufio.NewScanner(content)
	for nilai.Scan(){
		nilaiint, _ := strconv.Atoi(nilai.Text())
		bantu = append(bantu, nilaiint)
	}
	
	for i, j := range bantu{

		if (i + 2) % 3 == 0 && i != 0 {
			dari = append(dari, j)
		} else if ( i + 1) % 3 == 0 && i != 0 {
			sampai = append(sampai, j)
		} else if i % 3 == 0 && i != 0 {
			jarak = append(jarak, j)
		}
	}

	for i, j := range dari{
		for j <= sampai[i]{
			if j % jarak[i] == 0{
				dict[i] = append(dict[i], j)
			}
		}
	}

	for j,_ := range dict{
		fmt.Println("Case", j+1, ": ", len(dict[j]))
	}
}