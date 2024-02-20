package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
)

func Connect(ctx context.Context) (*mongo.Client, error) {
	// get mongouri from env
	mongoURL := os.Getenv("MONGO_URI")
	fmt.Println("mongoURL: ", mongoURL)
	//mongoURL := "mongodb://root:example@127.0.0.1:27017"
	// Set client options
	clientOptions := options.Client().ApplyURI(mongoURL)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("flowdash").Collection(collectionName)

}

//func Ping() {
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//	client, err := connect(ctx)
//	if err != nil {
//		cancel()
//		panic(err)
//	}
//	err = client.Disconnect(ctx)
//	if err != nil {
//		cancel()
//		panic(err)
//	}
//	defer cancel()
//}
