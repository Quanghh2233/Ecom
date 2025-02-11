package cart

import (
	"context"
	"time"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InstantBuyer(ctx context.Context, prodCollection, userCollection *mongo.Collection, productID primitive.ObjectID, UserID string) error {
	userObjID, err := primitive.ObjectIDFromHex(UserID)
	if err != nil {
		return ErrUserIdIsNotValid
	}

	var product models.ProdutUser
	err = prodCollection.FindOne(ctx, bson.M{"product_id": productID}).Decode(&product)
	if err != nil {
		return ErrCantFindProduct
	}

	order := models.Order{
		Order_ID:       primitive.NewObjectID(),
		Order_Cart:     []models.ProdutUser{product},
		Price:          product.Price,
		Ordered_At:     time.Now(),
		Payment_method: models.Payment{COD: true},
	}
	update := bson.M{"$push": bson.M{"orders": order}}

	_, err = userCollection.UpdateOne(ctx, bson.M{"user_id": userObjID}, update)
	if err != nil {
		return ErrCantBuyProduct
	}
	return nil
}
