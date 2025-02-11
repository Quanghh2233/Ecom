package global

import (
	"github.com/Quanghh2233/Ecommerce/internal/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var ProductCollection *mongo.Collection = database.ProductData(database.Client, "Products")
var UserCollection *mongo.Collection = database.UserData(database.Client, "Users")
var StoreCollection *mongo.Collection = database.ProductData(database.Client, "Store")
var App *Application

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
