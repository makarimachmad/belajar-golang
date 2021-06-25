package main

import (
	"fmt"
	"time"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type Product struct {
	ID          primitive.ObjectID     	`json:"_id" bson:"_id"`
	Name        string   			`json:"name" bson:"id"`
	Price       int   			`json:"price" bson:"price"`
	Currency    string   			`json:"currency" bson:"currency"`
	Quantity    int   			`json:"quantity" bson:"quantity"`
	Discount    int   			`json:"discount" bson:"discount"`
	Vendor      string  			`json:"vendor" bson:"vendor"`
	Accessories []string 			`json:"accessores,omitempty" bson:"accessories,omitempty"`
	SkuID       string   			`json:"sku_id" bson:"sku_id"`
}

var samsung = Product{
	ID:	primitive.NewObjectID(),
	Name:	"galaxyS21",
	Price:	100,
	Currency: "Rp",
	Quantity: 50,
	Discount: 10,
	Vendor: "samsungIndonesia",
	Accessories: []string{"charger","manual book","box"},
	SkuID: "12345",
}

var microsoft = Product{
	ID:	primitive.NewObjectID(),
	Name:	"SurfaceGo2",
	Price:	150,
	Currency: "Rp",
	Quantity: 50,
	Discount: 10,
	Vendor: "microsoftIndonesia",
	Accessories: []string{"charger","sling bag","box"},
	SkuID: "112233",
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil{
		fmt.Println(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	
	err = client.Connect(ctx)
	if err != nil{
		fmt.Println(err)
	}
	
	db := client.Database("tronics")
	collection := db.Collection("products")
	res, err := collection.InsertOne(context.Background(), samsung)

	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(res.InsertedID.(primitive.ObjectID).Timestamp())
}