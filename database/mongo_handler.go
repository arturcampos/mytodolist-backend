package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Collection *mongo.Collection
var Ctx = context.TODO()

func viperEnvVariable(key string) string {
	configFileName := os.Getenv("CONFIG_FILE")
	if len(configFileName) == 0 {
		configFileName = "config.yaml"
	}

	fmt.Println(configFileName)

	viper.SetConfigFile(configFileName)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
	value, ok := viper.Get(key).(string)

	if !ok {
		panic("Invalid type assertion")
	}

	return value
}

func init() {
	uri := viperEnvVariable("MONGO_URL")
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
