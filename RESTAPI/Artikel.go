//sumber belajar https://tutorialedge.net/golang/creating-restful-api-with-golang/
package main

import(
	"encoding/json"
	"fmt"
	"log"
	"io/ioutil"
	"net/http"

	//database
	// "database/sql"
	// _ "github.com/go-sql-driver/mysql"

	//mux router
	"github.com/gorilla/mux"
)

// variabel-variabel artikel
type Artikel struct {
	Id 		string `json:"Id"`
	Judul 	string `json:"Judul"`
	Penulis string `json:"Penulis"`
	Link 	string `json:"Link"`
}

// artikel
var Artikel_arr []Artikel

func Beranda(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Selamat datang di MakarimWeb")
	fmt.Println("Endpoint Hit: Beranda")
}

func returnSemuaArtikel(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: returnSemuaArtikel")
	json.NewEncoder(w).Encode(Artikel_arr)
}

func returnSingleArtikel(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]

	for _, artikel := range Artikel_arr{
		if artikel.Id == key{
			json.NewEncoder(w).Encode(artikel)
		}
	}
}

func tambahArtikel(w http.ResponseWriter, r *http.Request){
	reqBody, _ := ioutil.ReadAll(r.Body)
	var artikel Artikel
	json.Unmarshal(reqBody, &artikel)
	Artikel_arr = append(Artikel_arr, artikel)
}

func hapusArtikel(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]

	for index, artikel := range Artikel_arr{
		if artikel.Id == id{
			Artikel_arr = append(Artikel_arr[:index], Artikel_arr[index+1:]...)
		}
	}
}

func HandleRequest(){
	http.HandleFunc("/", Beranda)

	// jalur artikel dan memetakan pengembalian nilai
	// yang ada di semua fungsi artikel
	// 1
	// http.HandleFunc("/artikel", returnSemuaArtikel)
	// log.Fatal(http.ListenAndServe(":8000", nil))
	// 2
	// handlerequest menerapkan MuxRouter
	RouterSaya := mux.NewRouter().StrictSlash(true)
	RouterSaya.HandleFunc("/", Beranda)
	RouterSaya.HandleFunc("/artikel", returnSemuaArtikel)
	RouterSaya.HandleFunc("/detailartikel/{id}", returnSingleArtikel)
	RouterSaya.HandleFunc("/tambahArtikel", tambahArtikel)
	RouterSaya.HandleFunc("/hapusArtikel/{id}", hapusArtikel)

	log.Fatal(http.ListenAndServe(":8000", RouterSaya))
}

func main(){
	Artikel_arr = []Artikel{
		Artikel{
			Id: "1",
			Judul: "Dasar Machine Learning",
			Penulis: "Achmad",
			Link: "https://Dicoding.com",
		},
		Artikel{
			Id: "2",
			Judul: "Data Analysis Dengan Google DataStudio",
			Penulis: "Makarim",
			Link: "https://SkillAcademy.com",
		},
		Artikel{
			Id: "3",
			Judul: "Aplikasi Kognitif",
			Penulis: "Widyanto",
			Link: "https://Udemy.com",
		},
	}
	HandleRequest()
}
