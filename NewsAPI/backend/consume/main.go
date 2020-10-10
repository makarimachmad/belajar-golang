// uji coba ambil data dari newsapi org kategori bisnis
package main

import (
	"io/ioutil"
	//"log"
	"net/http"
	"belajar-golang/newsAPI/backend/consume/models"
	"encoding/json"
	"fmt"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


func main(){
	response, err := http.Get("https://newsapi.org/v2/top-headlines?apikey=6bc3cbc8dcf3473fb2527028734aedee&country=id&category=business")
	if err != nil {
		panic(err.Error())
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil{
		panic(err.Error())
	}

	var responseObjek models.Response
	json.Unmarshal(responseData, &responseObjek)
	
	fmt.Println(responseObjek.Totalresult)
	fmt.Println(len(responseObjek.Data))
	//fmt.Println(responseObjek.Data)

	for i := 0; i<len(responseObjek.Data); i++{
		fmt.Println("Source Id ",i," : ", responseObjek.Data[i].Source.Id)
		fmt.Println("Source Name ",i," : ", responseObjek.Data[i].Source.Name)
		fmt.Println("Author ",i," : ", responseObjek.Data[i].Author)
		fmt.Println("Title ",i," : ", responseObjek.Data[i].Title)
		fmt.Println("Description ",i," : ", responseObjek.Data[i].Description)
		fmt.Println("Url ",i," : ", responseObjek.Data[i].Url)
		fmt.Println("UrlToImage ",i," : ", responseObjek.Data[i].UrlToImage)
		fmt.Println("PublishedAt ",i," : ", responseObjek.Data[i].PublishedAt)
		fmt.Println("Content ",i," : ", responseObjek.Data[i].Content)
		fmt.Println("----------------------------------------")
	}

	Insert(responseObjek)
}

func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := ""
    dbName := "yii2basic"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())
    }
	return db
	
}
func Insert(responseObjek models.Response) {
    db := dbConn()
	for i := 0; i<len(responseObjek.Data); i++{
		sourceid := responseObjek.Data[i].Source.Id
		sourcename := responseObjek.Data[i].Source.Name
		author := responseObjek.Data[i].Author
		title := responseObjek.Data[i].Title
		description := responseObjek.Data[i].Description
		url := responseObjek.Data[i].Url
		UrlToImage := responseObjek.Data[i].UrlToImage
		publishedAt := responseObjek.Data[i].PublishedAt
		content := responseObjek.Data[i].Content

		insForm, err := db.Prepare("INSERT INTO berita(sourceid, sourcename, author, title, description, url, UrlToImage, publishedAt, content) VALUES(?,?,?,?,?,?,?,?,?)")
		
    	if err != nil {
        	panic(err.Error())
    	}
    	insForm.Exec(sourceid, sourcename, author, title, description, url, UrlToImage, publishedAt, content)
	}
    defer db.Close()
}
