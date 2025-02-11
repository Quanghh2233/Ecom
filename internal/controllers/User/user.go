package user

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	prodCollection  *mongo.Collection
	userCollection  *mongo.Collection
	storeCollection *mongo.Collection
	redisClient     *redis.Client
}

func NewApplication(prodCollection, userCollection, storeCollection *mongo.Collection, redisClient *redis.Client) *Application {
	app := &Application{
		prodCollection:  prodCollection,
		userCollection:  userCollection,
		storeCollection: storeCollection,
		redisClient:     redisClient,
	}
	log.Printf("Debug: userCollection assigned with collection name: %s", userCollection.Name())
	return app
}

func (app *Application) CheckUserStore(userID string) (bool, error) {
	if userID == "" {
		log.Println("[CheckUserStore] User ID is empty")
		return false, errors.New("userID is empty")
	}

	if app.storeCollection == nil {
		log.Println("[CheckUserStore] Store collection is nil")
		return false, errors.New("store collection is nil")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := app.storeCollection.CountDocuments(ctx, bson.M{"owner": userID})
	if err != nil {
		log.Printf("[CheckUserStore] Error counting stores: %v", err)
		return false, err
	}

	log.Printf("[CheckUserStore] Found %d stores for user %s", count, userID)
	return count > 0, nil
}
