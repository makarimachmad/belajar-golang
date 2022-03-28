package main

import (
  "os"
  "fmt"  
  "strings"  
)

func main() {	
	// Set custom env variable
	os.Setenv("CUSTOM", "500")
	
	// fetcha all env variables
	for _, element := range os.Environ() {
        variable := strings.Split(element, "=")
        fmt.Println(variable[0],"=>",variable[1])		
    }
	
	// fetch specific env variables
	fmt.Println("CUSTOM=>", os.Getenv("CUSTOM"))
	fmt.Println("GOROOT=>", os.Getenv("GOROOT"))
	fmt.Println("DBUSER=>", os.Getenv("DB_USER"))
	fmt.Println("DBHOST=>", os.Getenv("DB_HOST"))
	fmt.Println("DBPORT=>", os.Getenv("DB_PORT"))
}