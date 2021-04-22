package main

import (
	"fmt"
	//"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"z
)
func HandleRequest(){
	router := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "application/json; charset=UTF-8"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})


	router.HandleFunc("/biodata", getBiodata).Methods("GET")
	router.HandleFunc("/tambah", insertBiodata).Methods("POST")
	router.HandleFunc("/lihat", detailBiodata).Methods("GET")
	router.HandleFunc("/detail/{id}", getSingleBiodata).Methods("GET")
	router.HandleFunc("/update", updateBiodata).Methods("PUT")
	router.HandleFunc("/update", updateBiodata).Methods("PATCH")
	router.HandleFunc("/hapus/{id}", deleteBiodata).Methods("DELETE")

	//http.HandleFunc("/", handler)
	fmt.Println("Connected to port 8090")
	//log.Fatal(http.ListenAndServe(":8090", nil))
	http.ListenAndServe(":8090", handlers.CORS(headers, methods, origins)(router))

}
func main() {
	
	HandleRequest()
}