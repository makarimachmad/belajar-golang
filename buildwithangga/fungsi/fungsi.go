package main

import "fmt"
import "errors"

func main() {
	hasil := add(4, 9)
	fmt.Println(hasil)

	hasilkurang := kurang(4, 9)
	fmt.Println(hasilkurang)

	luas, keliling := calculate(10, 2)
	fmt.Println(luas, keliling)

	luasv2, kelilingv2 := calculatev2(10, 2)
	fmt.Println(luasv2, kelilingv2)

	//quiz
	nilai := []int{10, 5, 8, 9, 7}
	total := sum(nilai)
	fmt.Println(total)

	hasil, err := soal2(10, 2, "=")
	if err != nil{
		fmt.Println(err.Error())
	}
	fmt.Println(hasil)
}

func add(angka1 int, angka2 int) int{
	return angka1 + angka2
}

func kurang(angka1, angka2 int) int{
	return angka1 - angka2
}

func calculate(panjang int, lebar int) (int, int){
	luas := panjang * lebar
	keliling := 2 * (panjang+lebar)

	return luas, keliling
}

func calculatev2(panjang int, lebar int) (luas int, keliling int){
	luas = panjang * lebar
	keliling = 2 * (panjang+lebar)

	return
}

// quiz
func sum(nilai []int) int{
	var total int
	for _, bwa := range nilai{
		total = total + bwa
	}
	return total
}

func soal2(nilai1, nilai2 int, operasi string)(int, error){
	var hasil int
	var hasileror error
	switch operasi{
	case "+":
		hasil = nilai1 + nilai2
	case "-":
		hasil = nilai1 + nilai2
	case "*":
		hasil = nilai1 + nilai2
	case "/":
		hasil = nilai1 + nilai2
	default:
		hasileror = errors.New("tidak diketahui")
	}
	return hasil, hasileror
}