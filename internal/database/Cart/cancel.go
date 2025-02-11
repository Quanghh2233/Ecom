package cart

import (
	"context"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CancelAllOrder(ctx context.Context, userCollection *mongo.Collection, userID string) error {
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return ErrUserIdIsNotValid
	}

	update := bson.M{"$set": bson.M{"orders": []models.Order{}}}

	result, err := userCollection.UpdateOne(ctx, bson.M{"_id": userObjID}, update)
	if err != nil {
		return ErrCantCancelOrders
	}

	if result.MatchedCount == 0 {
		return ErrUserIdIsNotValid
	}
	return nil
}
