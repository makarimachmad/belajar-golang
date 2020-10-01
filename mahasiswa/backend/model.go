package main

type Biodata struct {
	Id		int `form:"id" json:"id"`
	Nama 	string `form:"nama" json:"nama"`
	Nim  	string `form:"nim" json:"nim"`
	Email 	string `form:"email" json:"email"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Biodata
}