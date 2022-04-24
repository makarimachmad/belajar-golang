package config

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func GetConnection() (*sql.DB, error){
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/toko")
	if err != nil{
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil{
		panic(err.Error())
	}
	return db, err
}