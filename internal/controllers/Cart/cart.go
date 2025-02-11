package Cart

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	prodCollection  *mongo.Collection
	userCollection  *mongo.Collection
	orderCollection *mongo.Collection
}

func NewApplication(prodCollection, userCollection, orderCollection *mongo.Collection) *Application {
	return &Application{
		prodCollection:  prodCollection,
		userCollection:  userCollection,
		orderCollection: orderCollection,
	}
}
