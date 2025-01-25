package repository

import (
	"context"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *cartRepository) GetUserCart(ctx context.Context, userID primitive.ObjectID) (*models.User, error) {
	var user models.User
	err := r.userCollection.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *cartRepository) AggregateCartTotal(ctx context.Context, userID primitive.ObjectID) (float64, error) {
	filterMatch := bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: userID}}}}
	unwind := bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$usercart"}}}}
	group := bson.D{{Key: "$group", Value: bson.D{{Key: "_id", Value: "$_id"}, {Key: "total", Value: bson.D{{Key: "$sum", Value: "$usercart.price"}}}}}}

	cursor, err := r.userCollection.Aggregate(ctx, mongo.Pipeline{filterMatch, unwind, group})
	if err != nil {
		return 0, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return 0, err
	}

	if len(results) == 0 {
		return 0, nil
	}

	total := results[0]["total"].(float64)
	return total, nil
}
