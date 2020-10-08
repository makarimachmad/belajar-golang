package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

var baseURL = "https://newsapi.org/v2/top-headlines?apikey=6bc3cbc8dcf3473fb2527028734aedee&country=id&category=business"

type berita struct {
	id 			int
	sourcename 	string
	sourceid 	string
	author		string
	title 		string
	url 		string
	urlToImage	string
	description	string
	content 	string
	publishedAt	string
}

func fetchBerita() ([]berita, error){
	var err error
	var client = &http.Client{}
	var data []berita

	request, err := http.NewRequest("GET", baseURL, nil)
	if err != nil{
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil{
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func main(){
	var berita, err = fetchBerita()
	if err != nil {
		fmt.Println("error bro", err.Error())
		return
	}

	for _, each := range berita{
		fmt.Printf("sourceid: %s\t sourcename: %s\t author: %s\t title: %s\t url: %s\t urlToImage: %s\t deskripsi: %s\t publishedAt: %s\t konten: %s\t", each.sourceid, each.sourcename, each.author, each.title, each.url, each.urlToImage, each.description, each.publishedAt, each.content)
	}
}