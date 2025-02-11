package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBSet() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://development:test@localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("Failed to connect to mongodb")
		return nil
	}
	fmt.Println("Successfully connected to mongodb")
	return client
}

var Client *mongo.Client = DBSet()

func UserData(client *mongo.Client, CollectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("Ecommerce").Collection(CollectionName)
	return collection
}

func ProductData(client *mongo.Client, CollectionName string) *mongo.Collection {
	var productcollection *mongo.Collection = client.Database("Ecommerce").Collection(CollectionName)
	return productcollection
}

func StoreData(client *mongo.Client, CollectionName string) *mongo.Collection {
	var storecollection *mongo.Collection = client.Database("Ecommerce").Collection(CollectionName)
	return storecollection
}

func OrderData(client *mongo.Client, CollectionName string) *mongo.Collection {
	var storecollection *mongo.Collection = client.Database("Ecommerce").Collection(CollectionName)
	return storecollection
}

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: "",
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Không thể kết nối Redis:", err)
	}

	log.Println("Đã kết nối Redis thành công!")

}
