package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	username	string = "root"
	password	string = ""
	database	string = "yii2basic"
)

var (
	dsn = fmt.Sprintf("%v:%v@/%v", username, password, database)
)

//menghubungkan ke database
func Mysql() (*sql.DB, error){
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}