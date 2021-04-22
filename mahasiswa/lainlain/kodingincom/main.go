package main
import (
    "net/http"
)
 
// Fungi Log yang berguna sebagai middleware untuk mengecek nilai param yg dikirim
func Auth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        uname, pwd, ok := r.BasicAuth()
        if !ok {
            w.Write([]byte("Username atau Password tidak boleh kosong"))
            return
        }
 
        if uname == "makarim" && pwd == "12345" {
            next.ServeHTTP(w, r)
            return
        }
 
        w.Write([]byte("Username atau Password tidak sesuai"))
        return
    })
}
 
// Fungsi GetMahasiswa untuk menampilkan text string param yg dikirim
func GetMahasiswa(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        w.Write([]byte("<h1>Anda Berhasil Mengakses Fungsi GetMahasiswa() </h1>"))
    }
}
 
func main() {
 
    // konfigurasi server
    server := &http.Server{
        Addr: ":8000",
    }
    // routing
    http.Handle("/", Auth(http.HandlerFunc(GetMahasiswa)))
 
    // jalankan server
    server.ListenAndServe()
}