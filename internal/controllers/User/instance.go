package user

import "go.mongodb.org/mongo-driver/mongo"

type Application struct {
	userCollection  *mongo.Collection
	storeCollection *mongo.Collection
}

func NewApplication(userCollection, storeCollection *mongo.Collection) *Application {

	return &Application{
		userCollection:  userCollection,
		storeCollection: storeCollection,
	}

}
