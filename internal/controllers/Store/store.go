package Store

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	prodCollection  *mongo.Collection
	userCollection  *mongo.Collection
	storeCollection *mongo.Collection
}

func NewApplication(prodCollection, userCollection, storeCollection *mongo.Collection) *Application {
	return &Application{
		prodCollection:  prodCollection,
		userCollection:  userCollection,
		storeCollection: storeCollection,
	}
}
