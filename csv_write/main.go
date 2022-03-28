package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	records := [][]string{
		{"nama_depan", "nama_belakang", "pekerjaan"},
		{"achmad", "makarim", "jr.software engineer"},
		{"makarim", "achmad", "data analyst"},
		{"widyanto", "makarim", "data engineer"},
		{"achmad", "widyanto", "data scientist"},
	}

	bantu := []string {"bantuan", "isi", "ketiga"}
	records = append(records, bantu)

	f, err := os.Create("user.csv")
	defer f.Close()
	if err != nil{
		log.Println("gagal membuat file", err)
	}

	w:=csv.NewWriter(f)
	err=w.WriteAll(records)
	if err != nil{
		log.Fatal(err)
	}
}