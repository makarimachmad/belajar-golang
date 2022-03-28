package main

import (
	"fmt"
	"strconv"
)

func main() {
	// isi := PickingNumber([]int{7, 12, 13, 19, 17, 7, 3, 18, 9, 18, 13, 12, 3, 13, 7, 9, 18, 9, 18, 9, 13, 18, 13 , 13, 18, 18, 17, 17, 13, 3, 12, 13, 19, 17, 19, 12, 18, 13, 7, 3, 3, 12, 7, 13, 7, 3, 17, 9, 13, 13, 13, 12, 18, 18, 9, 7, 19, 17, 13, 18, 19, 9, 18, 18, 18, 19, 17, 7, 12, 3, 13, 19, 12, 3, 9, 17, 13, 19, 12, 18, 13, 18, 18, 18, 17, 13, 3, 18, 19, 7, 12, 9 ,18, 3, 13, 13, 9, 7})
	isi := PickingNumber([]int{10,10,10,10,10})
	fmt.Println(isi)
}

func Kembalian(uang, bayar int) (string, []string) {
	rupiah := []int{100000, 50000, 25000, 10000, 5000, 1000}
	banyak := 0
	pecahan := 0
	Kembalian := uang - bayar
	hasil := []string{}

	if uang % 1000 != 0 || uang < bayar{
		return "invalid", nil
	}

	for _, i := range rupiah {
		if i > Kembalian {
			continue
		}
		banyak = Kembalian / i
		pecahan = i
		Kembalian = Kembalian - i
		hasil = append(hasil, strconv.Itoa(pecahan) + " ada " + strconv.Itoa(banyak) + ", ")
	}
	return "", hasil
}

func PickingNumber(angka []int) interface{} {
	dict := make(map[int]int)
	for _, j := range angka {
		dict[j] = dict[j] + 1
	}
	maks := 0
	if len(dict) > 1 {
		for i:=0; i < len(dict); i++ {
			for k, _ := range dict{
				if maks < dict[k] + dict[k+1]{
					maks = dict[k] + dict[k+1]
				}
			}
		}
	}else if len(dict) == 1 {
		for _, j := range dict{
			maks = j
		}
	}
	
	return maks
}