package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getBiodata(w http.ResponseWriter, r *http.Request) {
	var biodata Biodata
	var arr_biodata []Biodata
	var response Response

	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT id, nama, nim, email FROM biodata")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&biodata.Id, &biodata.Nama, &biodata.Nim, &biodata.Email); err != nil {
			log.Fatal(err.Error())

		} else {
			arr_biodata = append(arr_biodata, biodata)
		}
	}

	if arr_biodata == nil{
		response.Status = 400
		response.Message = "Error, Gak ada data"
	}else{
		response.Status = 200
		response.Message = "Success Show Data"
		response.Data = arr_biodata
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)

}

func insertBiodata(w http.ResponseWriter, r *http.Request) {
	//var biodata Biodata
	var arr_biodata []Biodata
	var response Response
	
	db := connect()
	defer db.Close()

	// pake form-data
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	nama := r.FormValue("nama")
	nim := r.FormValue("nim")
	email := r.FormValue("email")
	
	// pake x-www-form-urlencoded
	// err := r.ParseForm()
	// if err != nil {
	// 	panic(err)
	// }
	// nama := r.Form.Get("nama")
  	// nim := r.Form.Get("nim")
	// email := r.Form.Get("email")

	_, err = db.Exec("INSERT INTO biodata (nama, nim, email) values (?, ?, ?)",
		nama,
		nim,
		email,
	)
  
	if err != nil {
		log.Print(err)
	}
  
	response.Status = 200
	response.Message = "Berhasil Insert"
	response.Data = arr_biodata
		
	log.Print("Insert data to database")
  
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
  
}

func updateBiodata(w http.ResponseWriter, r *http.Request) {
	//var biodata Biodata
	var arr_biodata []Biodata
	var response Response
	

	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	id := r.FormValue("id")
	nama := r.FormValue("nama")
	nim := r.FormValue("nim")
	email := r.FormValue("email")

	// pake x-www-form-urlencoded
	// err := r.ParseForm()
	// if err != nil {
	// 	panic(err)
	// }
	// nama := r.Form.Get("nama")
  	// nim := r.Form.Get("nim")
	// email := r.Form.Get("email")

	_, err = db.Exec("UPDATE biodata set nama = ?, nim = ?, email = ? where id = ?",
		nama,
		nim,
		email,
		id,
	)

	if err != nil {
		log.Print(err)
	}

	response.Status = 200
	response.Message = "Berhasil Update Data"
	response.Data = arr_biodata

	log.Print("Update data ke database")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)

}

// func deleteBiodata(w http.ResponseWriter, r *http.Request) {

// 	var response Response

// 	db := connect()
// 	defer db.Close()

// 	err := r.ParseMultipartForm(4096)
// 	if err != nil {
// 		panic(err)
// 	}

// 	id := r.FormValue("id")

// 	_, err = db.Exec("DELETE from biodata where id = ?",
// 		id,
// 	)

// 	if err != nil {
// 		log.Print(err)
// 	}

// 	response.Status = 200
// 	response.Message = "Success Delete Data"
// 	log.Print("Delete data to database")

// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	json.NewEncoder(w).Encode(response)
// }

func deleteBiodata(w http.ResponseWriter, r *http.Request) {

	var response Response
	
	db := connect()
	defer db.Close()
	params := mux.Vars(r)


	stmt, err := db.Prepare("DELETE FROM biodata WHERE id = ?")
	if err != nil {
	  log.Print(err)
	}
	_, err = stmt.Exec(params["id"])
	if err != nil {
	  log.Print(err)
	}

	response.Status = 200
	response.Message = "Berhasil Hapus"

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	log.Print(w, "biodata dengan id = %s telah terhapus", params["id"])
	json.NewEncoder(w).Encode(response)
  }

func detailBiodata(w http.ResponseWriter, r *http.Request){
	var biodata Biodata
	var arr_biodata []Biodata
	var response Response

	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	id := r.FormValue("id")
	//id := r.URL.Query().Get("id")
	// id := mux.Vars(r)

	result, err := db.Query("SELECT * from biodata where id = ?",
		id,
	)

	if err != nil {
		log.Print(err)
	}
	defer result.Close()
	
	for result.Next() {
		if err := result.Scan(&biodata.Id, &biodata.Nama, &biodata.Nim, &biodata.Email); err != nil {
			log.Fatal(err.Error())

		} else {
			arr_biodata = append(arr_biodata, biodata)
		}
	}

	response.Status = 200
	response.Message = "Success Show Data"
	response.Data = arr_biodata
	log.Print("Detail data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func getSingleBiodata(w http.ResponseWriter, r *http.Request) {

	var response Response
	var biodata Biodata
	var arr_biodata []Biodata

	db := connect()
	defer db.Close()

	params := mux.Vars(r)
	hasil, err := db.Query("SELECT * FROM biodata WHERE id = ?", params["id"])
	if err != nil {
		log.Print(err)
	}
	defer hasil.Close()

	for hasil.Next() {
	  err := hasil.Scan(&biodata.Id, &biodata.Nama, &biodata.Nim, &biodata.Email)
	  if err != nil {
		log.Fatal(err.Error())
	  }else{
		  arr_biodata = append(arr_biodata, biodata)
	  }
	}

	if arr_biodata == nil{
		response.Status = 400
		response.Message = "Error, Gak ada data"
	}else{
		response.Status = 200
		response.Message = "Success Show Data"
		response.Data = arr_biodata
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
  }