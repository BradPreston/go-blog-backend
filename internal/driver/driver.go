package driver

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Collection *mongo.Collection
}

var conn = &MongoDB{}

func ConnectMongoDB(connString, database, collection string) (*MongoDB, error) {
	c, err := NewCollection(connString, database, collection)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	conn.Collection = c

	return conn, nil
}

func NewCollection(connString, database, collection string) (*mongo.Collection, error) {
	// env, err := godotenv.Read(".env")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	client, err := mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer client.Disconnect(ctx)

	c := client.Database(database).Collection(collection)

	return c, nil
}
