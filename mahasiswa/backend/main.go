package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)
func HandleRequest(){
	RouterSaya := mux.NewRouter().StrictSlash(true)
	RouterSaya.HandleFunc("/biodata", getBiodata).Methods("GET")
	RouterSaya.HandleFunc("/tambah", insertBiodata).Methods("POST")
	RouterSaya.HandleFunc("/lihat", detailBiodata).Methods("GET")
	RouterSaya.HandleFunc("/detail/{id}", getSingleBiodata).Methods("GET")
	RouterSaya.HandleFunc("/update", updateBiodata).Methods("PUT")
	RouterSaya.HandleFunc("/update", updateBiodata).Methods("PATCH")
	RouterSaya.HandleFunc("/hapus/{id}", deleteBiodata).Methods("DELETE")

	http.Handle("/", RouterSaya)
	fmt.Println("Connected to port 8090")
	log.Fatal(http.ListenAndServe(":8090", RouterSaya))

}
func main() {
	
	HandleRequest()
}