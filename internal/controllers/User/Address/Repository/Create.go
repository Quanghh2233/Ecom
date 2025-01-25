package repository

import (
	"context"
	"errors"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *addressRepository) AddAddress(ctx context.Context, userID primitive.ObjectID, address models.Address) error {
	filter := bson.M{"_id": userID}
	update := bson.M{"$push": bson.M{"address": address}}

	_, err := r.userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *addressRepository) GetAddressCount(ctx context.Context, userID primitive.ObjectID) (int32, error) {
	matchFilter := bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: userID}}}}
	projectStage := bson.D{{Key: "$project", Value: bson.D{{Key: "addressCount", Value: bson.D{{Key: "$size", Value: "$address"}}}}}}

	cursor, err := r.userCollection.Aggregate(ctx, mongo.Pipeline{matchFilter, projectStage})
	if err != nil {
		return 0, err
	}
	defer cursor.Close(ctx)

	var result []bson.M
	if err = cursor.All(ctx, &result); err != nil {
		return 0, err
	}

	if len(result) == 0 {
		return 0, errors.New("no result found")
	}

	addressCount := result[0]["addressCount"].(int32)
	return addressCount, nil
}
