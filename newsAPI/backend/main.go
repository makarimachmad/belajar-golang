// uji coba ambil data dari newsapi org kategori bisnis
package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main(){
	resp, err := http.Get("https://newsapi.org/v2/top-headlines?apikey=6bc3cbc8dcf3473fb2527028734aedee&country=id&category=business")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		log.Fatal(err)
	}
	log.Println(string(body))
}