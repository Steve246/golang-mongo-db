package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongodbUri = "mongodb://127.0.0.1:27017"

func main() {

	credential := options.Credential{
		Username: "stevejo",
		Password: "password",
	}
	// Username --> stevejo
	// Password --> password

	clientsOptions := options.Client()
	clientsOptions.ApplyURI(mongodbUri).SetAuth(credential)

	connect, err := mongo.Connect(context.Background(), clientsOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)

	defer cancel()

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected....")
	}

	defer func() {
		if err := connect.Disconnect(context.Background()); err != nil{
			panic(err)
		}
	}()

	//membuat sebuah db - collection

	db := connect.Database("enigma")
	coll := db.Collection("student")

	

	// Create 

	//insert One

	// newId, err := coll.InsertOne(ctx, bson.D{
	// 	{"name", "Jack"},
	// 	{"age", 22},
	// 	{"gender", "M"},
	// 	{"senior", false},
	// })

	//insert One to Many

	
	docs := []interface{}{
		bson.D{
		{"name", "Oscar"},
		{"age", 20},
		{"gender", "F"},
		{"senior", true},
		},
		bson.D{
			{"name", "Tano"},
			{"age", 22},
			{"gender", "M"},
			{"senior", false},
			},
	}

	result, err := coll.InsertMany(ctx, docs)


	if err != nil {
		log.Println(err.Error())
	}

	fmt.Printf("inserted document with Id %v\n", result.InsertedIDs)

}



/*

* Buat koneksi ke MongoDb (url) --> mongodb://127.0.0.1:27017
mongodb://localhost:27017

* Siapkan user Auth: username & password

*/