package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	
	// docs := []interface{}{
	// 	bson.D{
	// 	{"name", "Oscar"},
	// 	{"age", 20},
	// 	{"gender", "F"},
	// 	{"senior", true},
	// 	},
	// 	bson.D{
	// 		{"name", "Tano"},
	// 		{"age", 22},
	// 		{"gender", "M"},
	// 		{"senior", false},
	// 		},
	// }

	// result, err := coll.InsertMany(ctx, docs)


	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// fmt.Printf("inserted document with Id %v\n", result.InsertedIDs)

	//pake time parse

	jd01 := parseTime("2022-07-02 15:04:05")
	jd02 := parseTime("2022-07-03 15:04:05")
	// jd03 := parseTime("2022-07-04 15:04:05")

	
	
	docs := []interface{}{
		bson.D{
		{"name", "Melati"},
		{"age", 29},
		{"gender", "F"},
		{"joinDate", primitive.NewDateTimeFromTime(jd01)},
		{"senior", true},
		},
		bson.D{
			{"name", "Anggar"},
			{"age", 22},
			{"gender", "M"},
			{"joinDate", primitive.NewDateTimeFromTime(jd02)},
			{"senior", false},
			},
	}

	result, err := coll.InsertMany(ctx, docs)


	if err != nil {
		log.Println(err.Error())
	}

	fmt.Printf("inserted document with Id %v\n", result.InsertedIDs)

	fmt.Println(parseTime("2021-02-1"))

}

func parseTime(date string) time.Time {
	layoutFormat := "2006-01-02"
	parse, _ := time.Parse(layoutFormat, date)
	return parse
}



/*

* Buat koneksi ke MongoDb (url) --> mongodb://127.0.0.1:27017
mongodb://localhost:27017

* Siapkan user Auth: username & password

*/