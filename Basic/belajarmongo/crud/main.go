package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// You will be using this Trainer type later in the program
type Trainer struct {
    Name string
    Age  int
    City string
}

func main() {
	//buka koneksi
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	//berhasil tersambung
	fmt.Println("Connected to MongoDB!")

	//membuat database dan koleksi(field)
	collection := client.Database("voli").Collection("trainers")
	err = client.Disconnect(context.TODO())
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("koneksi mongoDB terputus")

	//set value
	ash := Trainer{"ash", 10, "malang"}
	misty := Trainer{"yoi", 14, "surabaya"}
	brock := Trainer{"ret", 20, "jakarta"}

	//insert 1 nilai
	ashResult, err := collection.InsertOne(context.TODO(), ash)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("dimasukkan ke dalam single dokumen", ashResult.InsertedID)
	
	trainers := []interface{}{misty, brock}
	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
}