package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongodbUri = "mongodb://127.0.0.1:27017"

type Student struct {
	Id primitive.ObjectID `bson:"_id"`
	Name string 	`bson:"fullname"`
	Age int  `bson:"age"`
	Gender string  `bson:"gender"`
	// JoinDate time.Time `bson:"joinDate"`

	JoinDate primitive.DateTime `bson:"joinDate"`

	Senior bool  `bson:"senior"`
}
func main() {

	credential := options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
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

	// jd01 := parseTime("2022-07-02 15:04:05")
	// jd02 := parseTime("2022-07-03 15:04:05")
	// jd03 := parseTime("2022-07-04 15:04:05")

	
	
	// students := []interface{}{
	// 	bson.D{
	// 	{"name", "Melati"},
	// 	{"age", 29},
	// 	{"gender", "F"},
	// 	{"joinDate", primitive.NewDateTimeFromTime(jd01)},
	// 	{"senior", true},
	// 	},
	// 	bson.D{
	// 		{"name", "Anggar"},
	// 		{"age", 22},
	// 		{"gender", "M"},
	// 		{"joinDate", primitive.NewDateTimeFromTime(jd02)},
	// 		{"senior", false},
	// 		},
	// }

	// result, err := coll.InsertMany(ctx, students)


	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// fmt.Printf("inserted document with Id %v\n", result.InsertedIDs)

	// fmt.Println(parseTime("2021-02-1"))


	//pake struct 

	// newStudent := Student{
	// Id: primitive.NewObjectID(), 
	// Name: "Steve" ,
	// Age: 23,
	// Gender: "M",
	// // JoinDate: parseTime("2022-07-13 15:04:03"),
	// JoinDate: primitive.NewDateTimeFromTime(parseTime("2022-07-13 00:00:00")),
	// Senior: false,
	// }

	// newId, err := coll.InsertOne(ctx, newStudent)


	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// fmt.Printf("inserted document with Id %v\n", newId.InsertedID)

	//update

	filterId, _ := primitive.ObjectIDFromHex("62ccfe1d8412e9e781a34716")

	// filterData := bson.D{{"_id", "62ccfe1d8412e9e781a34716"}}

	updateData := bson.D{{
		"$set", bson.D{{
			Key: "senior", Value: true,
		}},
	}}

	// result, err := coll.UpdateOne(ctx, filterId, updateData)
	result, _ := coll.UpdateOne(ctx, bson.M{"_id": filterId}, updateData)

	fmt.Printf("Documents updated: %v\n", result.MatchedCount)



	//delete

	// deleteData := bson.M{
	// 	"name": "Amberr",
	// }

	// result, err := coll.DeleteOne(ctx, deleteData)

	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)


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