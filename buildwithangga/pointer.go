package main

import "fmt"

//pointer struct
type murid struct{
	ID int
	Nama string
	GPA float32
}

//quiz

type gamer struct {
	Nama string
	Permainan []string
}

func main() {
	angkaA := 5
	angkaB := &angkaA

	fmt.Println(angkaA)
	fmt.Println(angkaB)
	fmt.Println(*angkaB)

	angkaA = 10
	fmt.Println(angkaA)

	var angkaC int = 5
	var angkaD *int = &angkaC
	fmt.Println(angkaD)

	//kasus pointer
	angkakasus := 5
	fmt.Println("angka awal: ", angkakasus)
	fmt.Println("alamat awal: ", &angkakasus)

	angkabaru := 100
	perubahangka(&angkakasus, angkabaru)
	fmt.Println("angka akhir: ", angkakasus)

	//pointer struct
	namamhs := murid{1,"Achmad Makarim Widyanto, ",3.47}
	perubahstruct(&namamhs)
	fmt.Println(namamhs)

	//method pointer
	namamhs.perubahanmethod()
	fmt.Println(namamhs)

	//quiz
	bantu := gamer{Nama:"makarim"}
	bantu.daftargamer("fifa 2020")
	bantu.daftargamer("pes 2021")
	bantu.daftargamer("moto gp")
	fmt.Println(bantu)

	for _, data := range bantu.Permainan{
		fmt.Println(data)
	}

}

//kasus pointer
func perubahangka(lama *int, baru int){
	*lama = baru //nilai dari alamat lama sama dengan baru
}

//pointer struct
func perubahstruct(siswa *murid){
	siswa.Nama = siswa.Nama + "S.Kom."
}

//method parameter
func (siswa *murid) perubahanmethod(){
	siswa.Nama = "Mas " + siswa.Nama
}

//quiz
func (pemain *gamer) daftargamer(nama string){
	pemain.Permainan = append(pemain.Permainan, nama)
}