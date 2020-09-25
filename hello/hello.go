package main

import (
	"fmt"
	"log"
	"example.com/greetings"
)

func main() {

	// mendefinisikan logger
	// log menandai untuk mencegah keluaran print
	// waktu, file, dan nomor baris
	
	// Get a greeting message and print it.
	message, err := greetings.Hello("Makarim")

	// untuk mengembalikan nilai jika error
	// dan akan keluar

	if err != nil{
		log.Fatal(err)
	}

	// kalo gk eror muncul pesan ini
	fmt.Println(message)
}
