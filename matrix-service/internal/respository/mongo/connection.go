package mongorepo

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbname           = "hacks"
	matrixCollection = "matrixes"
)

func CreateMongoClient(ctx context.Context) *mongo.Client {
	dbURI := os.Getenv("MONGO_URI")
	for i := 0; i < 5; i++ {
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbURI))
		if err == nil {
			err = client.Ping(context.TODO(), nil)
			if err == nil {
				return client
			}
		}
		log.Printf("Failed to connect to MongoDB (attempt %d): %v", i+1, err)
		time.Sleep(5 * time.Second)
	}
	log.Fatal("db is not answering ")
	return nil
}
