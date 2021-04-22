package models

type Mahasiswa struct {
	Id		int 	`json:"id"`
	Nama	string 	`json:"nama"`
	Nim		string	`json:"nim"`
	Email	string	`json:"email"`	
}