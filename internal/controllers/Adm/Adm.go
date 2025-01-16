package Adm

import (
	"github.com/Quanghh2233/Ecommerce/internal/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var ProductCollection *mongo.Collection = database.ProductData(database.Client, "Products")
