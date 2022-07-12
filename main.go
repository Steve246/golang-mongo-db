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

	// filterId, _ := primitive.ObjectIDFromHex("62ccfe1d8412e9e781a34716")

	// filterData := bson.D{{"_id", "62ccfe1d8412e9e781a34716"}}

	// updateData := bson.D{{
	// 	"$set", bson.D{{
	// 		Key: "senior", Value: true,
	// 	}},
	// }}

	// result, err := coll.UpdateOne(ctx, filterId, updateData)
	// result, _ := coll.UpdateOne(ctx, bson.M{"_id": filterId}, updateData)

	// fmt.Printf("Documents updated: %v\n", result.MatchedCount)



	//delete

	// deleteData := bson.M{
	// 	"name": "Amberr",
	// }

	// result, err := coll.DeleteOne(ctx, deleteData)

	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)


	//Read -- Find

	// cursor, err := coll.Find(ctx, bson.D{})

	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// var students []bson.D //nama collections, jangan sampe salah
	// err = cursor.All(ctx, &students)

	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// for _, student := range students {
	// 	fmt.Println(student)
	// }

	//projection 

	// opts := options.Find().SetProjection(bson.D{
	// 	{"_id", 0},
	// 	{"name", 1},
	// 	{"age", 1},
	// })

	// cursor, err := coll.Find(ctx, bson.D{}, opts)

	// if err != nil {
	// 	panic(err)
	// }

	// var students []bson.D

	// err = cursor.All(ctx, &students)

	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// for _, student := range students {
	// 	fmt.Println(student)
	// }

	//Logical

	// var students []bson.D

	// filterGenderAndAge := bson.D{
	// 	{"$and", bson.A{
	// 		bson.D{
	// 			{"gender", "F"},
	// 			{"age", bson.D{{"$gte", 25}}},
	// 		},
	// 	}},
	// }

	// projection := bson.D{
	// 	{"_id", 1}, 
	// 	{"fullName", 1},
	// 	{"gender", 1},
	// 	{"age", 1},
	// }

	// cursor, err := coll.Find(ctx, filterGenderAndAge, options.Find().SetProjection(projection))

	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// err = cursor.All(ctx, &students)

	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// for _, student := range students {
	// 	fmt.Println("FILTER BY GENDER & AGE", student)
	// }

	//merging result querry ke struct

	// var students []bson.D

	// filterGenderAndAge := bson.D{
	// 	{"$and", bson.A{
	// 		bson.D{
	// 			{"gender", "F"},
	// 			{"age", bson.D{{"$gte", 25}}},
	// 		},
	// 	}},
	// }

	// projection := bson.D{
	// 	{"_id", 1}, 
	// 	{"fullName", 1},
	// 	{"gender", 1},
	// 	{"age", 1},
	// }

	// filterGenderAndAgeResult := make([]*Student, 0)

	// cursor, err := coll.Find(ctx, filterGenderAndAge, options.Find().SetProjection(projection))

	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// for cursor.Next(ctx) {
	// 	var student Student
	// 	err := cursor.Decode(&student)
	// 	if err != nil {
	// 		log.Println(err.Error())
	// 	}
	// 	filterGenderAndAgeResult = append(filterGenderAndAgeResult, &student)
	// }

	// for _, student := range filterGenderAndAge {
	// 	fmt.Println("FILTER BY GENDER & AGE (WITH STRUCT)", student)
	// }


	//AGGREGATION

	// coll = connect.Database("enigma").Collection("products")

	// count, err := coll.CountDocuments(ctx, bson.D{})

	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println("Product Total: ", count)

	//with filter

	// count, err = coll.CountDocuments(ctx, bson.D{{"category", "food"}})

	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println("Product Total with category[food]: ", count)

	//Match, group, sort, dll

	matchStage := bson.D{
		{
			"$match", bson.D{
				{"category", "Food"},
			}},
		}

	groupStage := bson.D{
		{"$group", bson.D{
			{"_id", "$category"}, 
			{"Total", bson.D{{"$sum", 1}}},
			}},
		}

	cursor, err := coll.Aggregate(ctx, mongo.Pipeline{matchStage, groupStage})

	if err != nil {
		log.Println(err.Error())
	}

	var productCount []bson.M
	err = cursor.All(ctx, &productCount)
	if err != nil {
		log.Println(err.Error())
	}

	for _, product := range productCount {
		fmt.Printf("Group[%v], Total[%v]\n ", product["_id"], product["count"])
	}

	


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