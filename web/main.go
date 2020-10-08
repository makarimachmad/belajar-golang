//sumber belajar https://www.golangprograms.com/example-of-golang-crud-using-mysql-from-scratch.html
package main

import (
    "database/sql"
    "log"
    "net/http"
    "text/template"

    _ "github.com/go-sql-driver/mysql"
)

type Employee struct {
    Id    int
    Nama  string
	Nim string
	Email string
}

func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := ""
    dbName := "mahasiswa"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM biodata ORDER BY id ASC")
    if err != nil {
        panic(err.Error())
    }
    emp := Employee{}
    res := []Employee{}
    for selDB.Next() {
        var id int
        var nama, nim, email string
        err = selDB.Scan(&id, &nama, &nim, &email)
        if err != nil {
            panic(err.Error())
        }
        emp.Id = id
        emp.Nama = nama
		emp.Nim = nim
		emp.Email = email
        res = append(res, emp)
    }
    tmpl.ExecuteTemplate(w, "Index", res)
    defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")
    selDB, err := db.Query("SELECT * FROM biodata WHERE id=?", nId)
    if err != nil {
        panic(err.Error())
    }
    emp := Employee{}
    for selDB.Next() {
        var id int
        var nama, nim, email string
        err = selDB.Scan(&id, &nama, &nim, &email)
        if err != nil {
            panic(err.Error())
        }
        emp.Id = id
        emp.Nama = nama
		emp.Nim = nim
		emp.Email = email
    }
    tmpl.ExecuteTemplate(w, "Show", emp)
    defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
    tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    nId := r.URL.Query().Get("id")
    selDB, err := db.Query("SELECT * FROM biodata WHERE id=?", nId)
    if err != nil {
        panic(err.Error())
    }
    emp := Employee{}
    for selDB.Next() {
        var id int
        var nama, nim, email string
        err = selDB.Scan(&id, &nama, &nim, &email)
        if err != nil {
            panic(err.Error())
        }
        emp.Id = id
        emp.Nama = nama
		emp.Nim = nim
		emp.Email = email
    }
    tmpl.ExecuteTemplate(w, "Edit", emp)
    defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        nama := r.FormValue("nama")
		nim := r.FormValue("nim")
		email := r.FormValue("email")
		insForm, err := db.Prepare("INSERT INTO biodata(nama, nim, email) VALUES(?,?,?)")
		
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(nama, nim, email)
        log.Println("INSERT: Nama: " + nama + " | Nim: " + nim + "| Email: " + email)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        nama := r.FormValue("nama")
		nim := r.FormValue("nim")
		email := r.FormValue("email")
        id := r.FormValue("uid")
        insForm, err := db.Prepare("UPDATE biodata SET nama=?, nim=?, email=? WHERE id=?")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(nama, nim, email, id)
        log.Println("UPDATE: Nama: " + nama + " | Nim: " + nim + "| Email: " + email)
    }
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
    db := dbConn()
    emp := r.URL.Query().Get("id")
    delForm, err := db.Prepare("DELETE FROM biodata WHERE id=?")
    if err != nil {
        panic(err.Error())
    }
    delForm.Exec(emp)
    log.Println("DELETE")
    defer db.Close()
    http.Redirect(w, r, "/", 301)
}

func main() {
    log.Println("Server started on: http://localhost:1234")
    http.HandleFunc("/", Index)
    http.HandleFunc("/show", Show)
    http.HandleFunc("/new", New)
    http.HandleFunc("/edit", Edit)
    http.HandleFunc("/insert", Insert)
    http.HandleFunc("/update", Update)
    http.HandleFunc("/delete", Delete)
    http.ListenAndServe(":1234", nil)
}