//menjalankan aksi yang ada pada controller
//sumber belajar 1. https://kodingin.com/golang-koneksi-database-mysql/

package main

import (
	"log"
	"fmt"
	"context"
	"strconv"
	"net/http"
	//"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"belajar-golang/NewsAPI/backend/web/controller"
	"belajar-golang/NewsAPI/backend/web/models"
	"belajar-golang/NewsAPI/backend/web/utils"
)

func main() {

	handleRequest()

	err := http.ListenAndServe(":8090", nil)
	if err != nil{
		log.Fatal(err)
	}
}

//tukang panggil fungsi
func handleRequest(){
	
	router := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"Content-Type", "application/json; charset=UTF-8"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

    router.HandleFunc("/", getArticle).Methods("GET")
    router.HandleFunc("/lihat", getSelectedArticle).Methods("GET")
	
	fmt.Println("Terhubung ke 8090")
	http.ListenAndServe(":8090", handlers.CORS(headers, methods, origins)(router))
}

//fungsi middleware cek isian dan method GET untuk ambil data
func getArticle(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        ctx, cancel := context.WithCancel(context.Background())
		
        defer cancel()
		
        articles, err := article.GetAll(ctx)
 
        if err != nil {
            fmt.Println(err)
        }
 
        utils.ResponseJSON(w, articles, http.StatusOK)
        return
    }
 
    http.Error(w, "Tidak diijinkan", http.StatusNotFound)
    return
}

func getSelectedArticle(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
 
        ctx, cancel := context.WithCancel(context.Background())
        defer cancel()
 
        var art models.Article
 
        id := r.URL.Query().Get("id")
 
        if id == "" {
            utils.ResponseJSON(w, "id tidak boleh kosong", http.StatusBadRequest)
            return
        }
        art.Id, _ = strconv.Atoi(id)
 
        articles, err := article.GetSelected(ctx, art)
 
        if err != nil {
            fmt.Println(err)
        }
 
        utils.ResponseJSON(w, articles, http.StatusOK)
        return
 
    }
 
    http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
    return
}