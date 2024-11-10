package db

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	MongoDB *mongo.Database
}

func DBConnection() (*DB, error) {
	godotenv.Load(".env")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGODB")).SetServerAPIOptions(serverAPI)

	//konek ke database
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}

	db := client.Database("BookStore")
	return &DB{
		db,
	}, nil
}
