package main

import (

	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	buatTabel()
}

func koneksiMysql() (db *sql.DB){
	dbDriver := "mysql"
    dbUser := "root"
    dbPass := ""
    db, err := sql.Open(dbDriver, dbUser + ":" + dbPass + "@/")
    if err != nil {
        panic(err.Error())
    }
	return db
}

func buatDB(){

	db := koneksiMysql()

	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS ujian")
	
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Successfully created database..")
	}
}

func koneksiDB() (db *sql.DB) {

    dbDriver := "mysql"
    dbUser := "root"
    dbPass := ""
    dbName := "ujian"
    db, err := sql.Open(dbDriver, dbUser+":"+  dbPass + "@/" + dbName)
    if err != nil {
        panic(err.Error())
    }
	return db
}

func buatTabel(){

	db := koneksiDB()

	_, err := db.Exec("USE ujian")
	
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB selected successfully..")
	}

	stmt, err := db.Prepare(`CREATE TABLE employee (
		id int NOT NULL AUTO_INCREMENT, 
		first_name varchar(50), 
		last_name varchar(30), 
		PRIMARY KEY (id) )`)

	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec()
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Table created successfully..")
	}
	defer db.Close()
}