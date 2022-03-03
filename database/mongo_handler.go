package database

import (	
	"time"
	"fmt"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)
const uri = "mongodb://root:MongoDB2021!@localhost:27017/?maxPoolSize=20&w=majority"

var Collection *mongo.Collection
var Ctx = context.TODO()

func init() {

	clientOptions := options.Client().ApplyURI(uri)
	client, _ := mongo.NewClient(clientOptions)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := client.Connect(ctx)

	if err != nil {
		panic(err)
	}

	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	Collection = client.Database("todo-list").Collection("tasks")
	fmt.Println("Successfully connected and pinged.")

}