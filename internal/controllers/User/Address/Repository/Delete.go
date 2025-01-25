package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *addressRepository) DeleteAddress(ctx context.Context, userID, addressID primitive.ObjectID) (int64, error) {
	filter := bson.M{"_id": userID}
	update := bson.M{"$pull": bson.M{"address": bson.M{"_id": addressID}}}

	result, err := r.userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil
}
