package cart

import (
	"context"
	"time"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func BuyItemFromCart(ctx context.Context, userCollection *mongo.Collection, userID string) (*models.Order, error) {
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, ErrUserIdIsNotValid
	}

	var user models.User
	err = userCollection.FindOne(ctx, bson.M{"_id": userObjID}).Decode(&user)
	if err != nil {
		return nil, ErrUserIdIsNotValid
	}

	if len(user.UserCart) == 0 {
		return nil, ErrCartEmpty
	}

	order := models.Order{
		Order_ID:       primitive.NewObjectID(),
		Order_Cart:     user.UserCart,
		Price:          calculateTotalPrice(user.UserCart),
		Ordered_At:     time.Now(),
		Payment_method: models.Payment{COD: true},
	}

	update := bson.M{
		"$push": bson.M{"orders": order},
		"$set":  bson.M{"usercart": []models.ProdutUser{}},
	}

	_, err = userCollection.UpdateOne(ctx, bson.M{"_id": userObjID}, update)
	if err != nil {
		return nil, ErrCantBuyCartItem
	}
	return &order, nil
}
