package main

import "fmt"

type BangunDatar interface{
	HitungLuas() int
}

type persegi struct {
	Sisi int
}

type persegipanjang struct {
	Panjang int
	Lebar int
}

func (nilai persegi) HitungLuas() int{
	return nilai.Sisi * nilai.Sisi
}

func (nilai persegipanjang) HitungLuas() int{
	return nilai.Panjang * nilai.Lebar
}

func HitungBangunDatar(nilai BangunDatar) int{
	return nilai.HitungLuas()
}

func main() {
	//untuk memfleksibelkan parameter yg akan dipakai tanpa harus membuat banyak method

	persegipanjang := persegipanjang{Panjang:5, Lebar:8}
	luas := HitungBangunDatar(persegipanjang)
	fmt.Println(luas)

	persegi := persegi{Sisi:7}
	luas = HitungBangunDatar(persegi)
	fmt.Println(luas)
}