package initializers

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var CLIENT *mongo.Client
var DBNAME string

func ConnectToDb() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") // Update URI as needed
	CLIENT, _ = mongo.Connect(context.Background(), clientOptions)
	DBNAME = os.Getenv("DBNAME")

	err := CLIENT.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	return CLIENT, nil
}
