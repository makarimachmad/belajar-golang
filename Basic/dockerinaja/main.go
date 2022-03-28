package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

type Users struct{
	Id		int	`json:"id"`
	Username	string	`json:"username"`
	IsActive	int	`json:"is_active"`
}

func main() {
	e := echo.New()
	var users	Users
	var arr_users	[]Users


	//db config ditaruh dalam env docker
	db_user	:= os.Getenv("DB_User")
	db_pass	:= os.Getenv("DB_PASS")
	db_host := os.Getenv("DB_HOST")
	db_name := os.Getenv("DB_NAME")
	db_port := os.Getenv("DB_PORT")
	apps_port := os.Getenv("PORT")

	db, err := sql.Open("mysql", db_user+":"+db_pass+"@tcp("+db_host+":"+db_port+")/"+db_name)
	
	defer db.Close()
	
	if err != nil{
		fmt.Println(err.Error())
	}

	//routing
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "berhasil")
	})

	e.GET("/user/getall", func(c echo.Context) error {
		arr_users = []Users{}
		res, err := db.Query("SELECT * FROM users order by id DESC")
		if err != nil{
			return c.String(http.StatusInternalServerError, err.Error())
		}

		for res.Next(){
			err = res.Scan(&users.Id, &users.Username, &users.IsActive)

			if err != nil{
				return c.String(http.StatusInternalServerError, err.Error())
			}
			arr_users = append(arr_users, users)
		}
		return c.JSON(http.StatusOK, arr_users)
	})
	e.Logger.Fatal(e.Start(":"+apps_port))
}