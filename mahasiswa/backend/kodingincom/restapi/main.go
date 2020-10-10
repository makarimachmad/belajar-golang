//menjalankan aksi yang ada pada controller
//sumber belajar 1. https://kodingin.com/golang-koneksi-database-mysql/

package main

import (
	"log"
	"fmt"
	"context"
	"strconv"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"belajar-golang/mahasiswa/backend/kodingincom/restapi/controller"
	"belajar-golang/mahasiswa/backend/kodingincom/restapi/models"
	"belajar-golang/mahasiswa/backend/kodingincom/restapi/utils"
)

func main() {

	handleRequest()

    //koneksi db
	// db, e := config.Mysql()
	// if e != nil{
	// 	log.Fatal(e)
	// }

	// eb := db.Ping()
	// if eb != nil{
	// 	panic(eb.Error())
	// }

	// fmt.Println("Berhasil gan")

	err := http.ListenAndServe(":8090", nil)
	if err != nil{
		log.Fatal(err)
	}
}

//tukang panggil fungsi
func handleRequest(){
	// http.HandleFunc("/", getMahasiswa)
	// http.HandleFunc("/tambah", postMahasiswa)
	router := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"Content-Type", "application/json; charset=UTF-8"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})


    router.HandleFunc("/", getMahasiswa).Methods("GET")
	router.HandleFunc("/tambah", tambahMahasiswa).Methods("POST")
	router.HandleFunc("/update", updateMahasiswa).Methods("PUT")
	router.HandleFunc("/update", updateMahasiswa).Methods("PATCH")
	router.HandleFunc("/hapus", hapusMahasiswa).Methods("DELETE")
	

	fmt.Println("Terhubung ke 8090")
	http.ListenAndServe(":8090", handlers.CORS(headers, methods, origins)(router))
}

//fungsi middleware cek isian dan method GET untuk ambil data
func getMahasiswa(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        ctx, cancel := context.WithCancel(context.Background())
		
        defer cancel()
		
        mahasiswas, err := mahasiswa.GetAll(ctx)
 
        if err != nil {
            fmt.Println(err)
        }
 
        utils.ResponseJSON(w, mahasiswas, http.StatusOK)
        return
    }
 
    http.Error(w, "Tidak diijinkan", http.StatusNotFound)
    return
}

//fungsi middleware cek isian dan method POST untuk insert data
func tambahMahasiswa(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
		
		//data yg dikirim harus berupa json
        if r.Header.Get("Content-Type") != "application/json" {
            http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
            return
        }
 
        ctx, cancel := context.WithCancel(context.Background())
        defer cancel()
 
        var mhs models.Mahasiswa
 
        if err := json.NewDecoder(r.Body).Decode(&mhs); err != nil {
            utils.ResponseJSON(w, err, http.StatusBadRequest)
            return
        }
 
        if err := mahasiswa.Tambah(ctx, mhs); err != nil {
            utils.ResponseJSON(w, err, http.StatusInternalServerError)
            return
        }
 
        res := map[string]string{
            "status": "berhasil gan",
        }
 
        utils.ResponseJSON(w, res, http.StatusCreated)
        return
    }
 
    http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
    return
}

//memperbarui data dengan fungsi middleware
func updateMahasiswa(w http.ResponseWriter, r *http.Request) {
    if r.Method == "PUT" {
 
        if r.Header.Get("Content-Type") != "application/json" {
            http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
            return
        }
 
        ctx, cancel := context.WithCancel(context.Background())
        defer cancel()
 
        var mhs models.Mahasiswa
 
        if err := json.NewDecoder(r.Body).Decode(&mhs); err != nil {
            utils.ResponseJSON(w, err, http.StatusBadRequest)
            return
        }
 
        fmt.Println(mhs)
 
        if err := mahasiswa.Update(ctx, mhs); err != nil {
            utils.ResponseJSON(w, err, http.StatusInternalServerError)
            return
        }
 
        res := map[string]string{
            "status": "Succesfully",
        }
 
        utils.ResponseJSON(w, res, http.StatusCreated)
        return
    }
 
    http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
    return
}

//hapus data
func hapusMahasiswa(w http.ResponseWriter, r *http.Request) {
 
    if r.Method == "DELETE" {
 
        ctx, cancel := context.WithCancel(context.Background())
        defer cancel()
 
        var mhs models.Mahasiswa
 
        id := r.URL.Query().Get("id")
 
        if id == "" {
            utils.ResponseJSON(w, "id tidak boleh kosong", http.StatusBadRequest)
            return
        }
        mhs.Id, _ = strconv.Atoi(id)
 
        if err := mahasiswa.Hapus(ctx, mhs); err != nil {
 
            kesalahan := map[string]string{
                "error": fmt.Sprintf("%v", err),
            }
 
            utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
            return
        }
 
        res := map[string]string{
            "status": "Succesfully",
        }
 
        utils.ResponseJSON(w, res, http.StatusOK)
        return
    }
 
    http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
    return
}