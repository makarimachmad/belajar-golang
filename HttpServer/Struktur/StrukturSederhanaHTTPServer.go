package main

import (
	"fmt"
	"log"
	"net/http"
)

func Berandaas(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Selamat datang di MakarimWeb")
	fmt.Println("Endpoint Hit: Beranda")
}

func HandleRequests() {
	http.HandleFunc("/", Berandaas)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	HandleRequests()
}