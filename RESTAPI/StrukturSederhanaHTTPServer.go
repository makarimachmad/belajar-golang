package main

import(
	"fmt"
	"log"
	"net/http"
)

func Beranda(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Selamat datang di MakarimWeb")
	fmt.Println("Endpoint Hit: Beranda")
}

func HandleRequest(){
	http.HandleFunc("/", Beranda)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	HandleRequest()
}