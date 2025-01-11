package cart

import (
	"context"
	"time"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func BuySelectedItems(ctx context.Context, userCollection *mongo.Collection, userID string, selectedItems []string) (*models.Order, error) {
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

	var orderItems []models.ProdutUser
	for _, item := range selectedItems {
		for _, cartItem := range user.UserCart {
			if cartItem.Product_ID.Hex() == item {
				orderItems = append(orderItems, cartItem)
			}
		}
	}

	if len(orderItems) == 0 {
		return nil, ErrorNoItemFound
	}

	order := models.Order{
		Order_ID:       primitive.NewObjectID(),
		Order_Cart:     orderItems,
		Price:          calculateTotalPrice(orderItems),
		Ordered_At:     time.Now(),
		Payment_method: models.Payment{COD: true},
	}

	update := bson.M{
		"$push": bson.M{"orders": order},
		"$pull": bson.M{"usercart": bson.M{"productid": bson.M{"$in": selectedItems}}},
	}

	_, err = userCollection.UpdateOne(ctx, bson.M{"_id": userObjID}, update)
	if err != nil {
		return nil, ErrCantBuyCartItem
	}
	return &order, nil

}
