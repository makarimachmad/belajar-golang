package main

import (
	"fmt"
	"net/http")

func myLogin(w http.ResponseWriter, r *http.Request){

	if r.Method == "GET"{
		fmt.Fprintln(w, "menggunakan method get")
	}
	if r.Method == "POST"{
		fmt.Fprintln(w, "menggunakan method post")
	}
	fmt.Fprintln(w, "ssiapp")
	fmt.Fprintf(w, "request is: %+v", r)
}


func myWelcome(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "welcome")
	fmt.Fprintf(w, "request is: %+v", r)
}


func main() {
	http.Handle("/login", http.HandlerFunc(myLogin))
	http.Handle("/welcome", http.HandlerFunc(myWelcome))
	http.ListenAndServe("localhost:8080", nil)
}