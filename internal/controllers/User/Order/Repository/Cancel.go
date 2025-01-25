package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *orderRepository) CancelOrder(ctx context.Context, userID, orderID primitive.ObjectID) (int64, error) {
	filter := bson.M{
		"_id":    userID,
		"orders": bson.M{"$elemMatch": bson.M{"_id": orderID}},
	}

	update := bson.M{
		"$pull": bson.M{"orders": bson.M{"_id": orderID}},
	}

	result, err := r.userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil
}

func (r *orderRepository) CancelAllOrders(ctx context.Context, userID string) error {
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("invalid user ID format")
	}

	filter := bson.M{"_id": userObjID}
	update := bson.M{"$set": bson.M{"orders": []bson.M{}}}

	_, err = r.userCollection.UpdateOne(ctx, filter, update)
	return err
}
