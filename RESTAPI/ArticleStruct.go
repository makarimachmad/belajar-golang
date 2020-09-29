package main

import(
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// variabel-variabel artikel
type Artikel struct {
	Judul string `json:"Judul"`
	Penulis string `json:"Penulis"`
	Link string `json:"Link"`
}

// artikel
var Artikel_arr []Artikel

func Beranda(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Selamat datang di MakarimWeb")
	fmt.Println("Endpoint Hit: Beranda")
}

func HandleRequest(){
	http.HandleFunc("/", Beranda)

	// jalur artikel dan memetakan pengembalian nilai
	// yang ada di semua fungsi artikel
	http.HandleFunc("/artikel", returnSemuaArtikel)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func returnSemuaArtikel(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: returnSemuaArtikel")
	json.NewEncoder(w).Encode(Artikel_arr)
}

func main(){
	Artikel_arr = []Artikel{
		Artikel{
			Judul: "Dasar Machine Learning",
			Penulis: "Achmad",
			Link: "https://Dicoding.com",
		},
		Artikel{
			Judul: "Data Analysis Dengan Google DataStudio",
			Penulis: "Makarim",
			Link: "https://SkillAcademy.com",
		},
		Artikel{
			Judul: "Aplikasi Kognitif",
			Penulis: "Widyanto",
			Link: "https://Udemy.com",
		},
	}
	HandleRequest()
}
